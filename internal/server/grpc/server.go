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

package grpc

import (
	"apps/stardust/internal/service"
	"apps/stardust/pb/v1/stardust"
	"fmt"
	"net"

	"github.com/UnderTreeTech/waterdrop/pkg/conf"
	"github.com/UnderTreeTech/waterdrop/pkg/registry"
	"github.com/UnderTreeTech/waterdrop/pkg/server/rpc"
	"github.com/UnderTreeTech/waterdrop/pkg/utils/xnet"
	"google.golang.org/grpc"
)

type ServerInfo struct {
	Server      *rpc.Server
	ServiceInfo *registry.ServiceInfo
}

func New(s *service.Service) *ServerInfo {
	srvConfig := &rpc.ServerConfig{}
	if err := conf.Unmarshal("server.rpc", srvConfig); err != nil {
		panic(fmt.Sprintf("unmarshal grpc server config fail, err msg %s", err.Error()))
	}

	server := rpc.NewServer(srvConfig)
	registerServers(server.Server(), s)

	addr := server.Start()
	_, port, _ := net.SplitHostPort(addr.String())
	serviceInfo := &registry.ServiceInfo{
		Name:    "service.stardust.v1",
		Scheme:  "grpc",
		Addr:    fmt.Sprintf("%s://%s:%s", "grpc", xnet.InternalIP(), port),
		Version: "1.0.0",
	}

	return &ServerInfo{Server: server, ServiceInfo: serviceInfo}
}

func registerServers(g *grpc.Server, s *service.Service) {
	stardust.RegisterStarDustServer(g, s)
}
