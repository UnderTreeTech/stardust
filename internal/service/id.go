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
	"apps/stardust/pb/v1/stardust"
	"context"

	"github.com/UnderTreeTech/waterdrop/pkg/log"
)

func (s *Service) GetUniqueIds(ctx context.Context, req *stardust.IdReq) (reply *stardust.IdsReply, err error) {
	reply = &stardust.IdsReply{}
	ids, err := s.pumper.NextIds(ctx, req.BizType, req.Len)
	if err != nil {
		return reply, err
	}

	log.Info(ctx, "get ids", log.Any("ids", ids))

	reply.Ids = ids
	return
}

func (s *Service) GetUniqueId(ctx context.Context, req *stardust.IdReq) (reply *stardust.IdReply, err error) {
	reply = &stardust.IdReply{}
	id, err := s.pumper.NextId(ctx, req.BizType)
	if err != nil {
		return reply, err
	}

	log.Info(ctx, "get id", log.Int64("id", id))

	reply.Id = id
	return
}

func (s *Service) ParseId(ctx context.Context, req *stardust.ParseReq) (reply *stardust.ParseReply, err error) {
	parsed := s.pumper.ParseId(ctx, req.Id)

	reply = &stardust.ParseReply{}
	reply.Idc = parsed.DataCenter
	reply.BizType = parsed.BizType
	reply.Time = parsed.Time
	reply.Worker = parsed.WorkerId
	reply.Sequence = parsed.Sequence

	log.Info(ctx, "parse info",
		log.Int64("parse_id", req.Id),
		log.Int64("idc", reply.Idc),
		log.Int64("biz_type", reply.BizType),
		log.Int64("timestamp", reply.Time),
		log.Int64("worker", reply.Worker),
		log.Int64("sequence", reply.Sequence),
	)

	return
}
