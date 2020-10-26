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

package http

import (
	"apps/stardust/internal/service"
	"fmt"
	"net"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/UnderTreeTech/waterdrop/pkg/conf"

	"github.com/UnderTreeTech/waterdrop/pkg/utils/xnet"

	"github.com/UnderTreeTech/waterdrop/pkg/registry"

	"github.com/UnderTreeTech/waterdrop/pkg/server/http"
)

var svc *service.Service

type ServerInfo struct {
	Server      *http.Server
	ServiceInfo *registry.ServiceInfo
}

func New(s *service.Service) *ServerInfo {
	srvConfig := &http.ServerConfig{}
	if err := conf.Unmarshal("server.http", srvConfig); err != nil {
		panic(fmt.Sprintf("unmarshal http server config fail, err msg %s", err.Error()))
	}
	server := http.NewServer(srvConfig)
	router(server)

	addr := server.Start()
	_, port, _ := net.SplitHostPort(addr.String())
	serviceInfo := &registry.ServiceInfo{
		Name:    "server.http.stardust",
		Scheme:  "http",
		Addr:    fmt.Sprintf("%s://%s:%s", "http", xnet.InternalIP(), port),
		Version: "1.0.0",
	}
	svc = s
	binding.Validator.Engine().(*validator.Validate).SetTagName("validate")
	return &ServerInfo{Server: server, ServiceInfo: serviceInfo}
}

func router(s *http.Server) {
	g := s.Group("/api")
	{
		g.GET("/stardust/id", getId)
		g.GET("/stardust/ids", getIds)
		g.GET("/stardust/parse", parseId)
	}
}
