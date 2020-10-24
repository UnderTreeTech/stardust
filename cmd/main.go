/*
 *
 * Copyright 2020 stardust authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"apps/stardust/internal/dao"
	"apps/stardust/internal/server/grpc"
	"apps/stardust/internal/server/http"
	"apps/stardust/internal/service"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UnderTreeTech/waterdrop/pkg/registry/etcd"
	"google.golang.org/grpc/resolver"

	"github.com/UnderTreeTech/waterdrop/pkg/log"

	"github.com/UnderTreeTech/waterdrop/pkg/conf"
	"github.com/UnderTreeTech/waterdrop/pkg/trace/jaeger"
)

type App struct {
	svc  *service.Service
	http *http.ServerInfo
	grpc *grpc.ServerInfo
}

func NewApp(svc *service.Service, h *http.ServerInfo, g *grpc.ServerInfo) (*App, func()) {
	etcdConf := &etcd.Config{}
	if err := conf.Unmarshal("etcd", etcdConf); err != nil {
		panic(fmt.Sprintf("unmarshal etcd config fail, err msg %s", err.Error()))
	}
	etcd := etcd.New(etcdConf)
	etcd.Register(context.Background(), g.ServiceInfo)
	etcd.Register(context.Background(), h.ServiceInfo)
	resolver.Register(etcd)

	app := &App{
		svc:  svc,
		http: h,
		grpc: g,
	}

	srvClose := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		if err := g.Server.Stop(ctx); err != nil {
			log.Errorf("stop grpc server fail", log.String("error", err.Error()))
		}
		if err := h.Server.Stop(ctx); err != nil {
			log.Errorf("stop http server fail", log.String("error", err.Error()))
		}
		cancel()
	}

	close := func() {
		srvClose()
		svc.Close()
		etcd.Close()
	}

	return app, close
}

func main() {
	flag.Parse()

	conf.Init()
	defer jaeger.Init()()

	logConfig := &log.Config{}
	if err := conf.Unmarshal("log", logConfig); err != nil {
		panic(fmt.Sprintf("parse log config fail, err msg %s", err.Error()))
	}
	defer log.New(logConfig).Sync()

	d := dao.New()
	service := service.New(d)
	http := http.New(service)
	rpc := grpc.New(service)
	_, close := NewApp(service, http, rpc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("get signal", log.String("signal", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			close()
			time.Sleep(time.Second)
			log.Infof("stardust exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
