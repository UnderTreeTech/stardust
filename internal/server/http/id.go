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
	"apps/stardust/pb/v1/stardust"
	"net/http"

	"github.com/UnderTreeTech/waterdrop/pkg/utils/xreply"

	"github.com/gin-gonic/gin"
)

func getId(c *gin.Context) {
	req := &stardust.IdReq{}
	if err := c.BindQuery(req); err != nil {
		return
	}

	reply, err := svc.GetUniqueId(c.Request.Context(), req)
	c.JSON(http.StatusOK, xreply.Reply(c.Request.Context(), reply, err))
}

func getIds(c *gin.Context) {
	req := &stardust.IdReq{}
	if err := c.BindQuery(req); err != nil {
		return
	}

	reply, err := svc.GetUniqueIds(c.Request.Context(), req)
	c.JSON(http.StatusOK, xreply.Reply(c.Request.Context(), reply, err))
}

func parseId(c *gin.Context) {
	req := &stardust.ParseReq{}
	if err := c.BindQuery(req); err != nil {
		return
	}

	reply, err := svc.ParseId(c.Request.Context(), req)
	c.JSON(http.StatusOK, xreply.Reply(c.Request.Context(), reply, err))
}
