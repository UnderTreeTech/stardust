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

package service

import (
	"apps/stardust/internal/dao"
	"apps/stardust/internal/snowflake"
	"context"
	"fmt"

	"github.com/UnderTreeTech/waterdrop/pkg/conf"
	"github.com/golang/protobuf/ptypes/empty"
)

type Service struct {
	dao    dao.Dao
	pumper *snowflake.IdWorker
}

type SnowflakeConf struct {
	DataCenterId int64
	WorkerId     int64
	MaxIdNum     int64
}

func New(d dao.Dao) *Service {
	pumperConfig := &SnowflakeConf{}
	if err := conf.Unmarshal("Snowflake", pumperConfig); err != nil {
		panic(fmt.Sprintf("unmarshal snowflake config fail,err is %s", err.Error()))
	}

	pumper := snowflake.New(pumperConfig.DataCenterId, pumperConfig.WorkerId, pumperConfig.MaxIdNum)
	service := &Service{
		dao:    d,
		pumper: pumper,
	}

	d.SyncWorkerInfo(pumper)

	return service
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
