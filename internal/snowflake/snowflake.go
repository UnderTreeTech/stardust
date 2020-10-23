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

package snowflake

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/UnderTreeTech/waterdrop/pkg/status"

	"github.com/UnderTreeTech/waterdrop/pkg/log"
)

const (
	epoch              = int64(1603469789668)                 //初始时间值为某一天的时间戳毫秒级
	workerIdBits       = uint(3)                              //3位worker
	dataCenterBits     = uint(2)                              //2位idc
	businessTypeBits   = uint(5)                              //5位业务线
	maxWorkerId        = -1 ^ (-1 << workerIdBits)            // 最大机器id是7,8台机器
	maxBusinessType    = -1 ^ (-1 << businessTypeBits)        //最大业务线31,32个业务
	maxDataCenter      = -1 ^ (-1 << dataCenterBits)          //最大机房数3，4个idc
	sequenceBits       = uint(12)                             //12位序列号
	workerIdShift      = sequenceBits                         //机器ID位于序列号左侧12位
	dataCenterShift    = workerIdShift + workerIdBits         //idc位于序列号左侧15位
	businessTypeShift  = dataCenterShift + dataCenterBits     // 业务id位于序列号左侧12+3+2=17之后
	timestampLeftShift = businessTypeShift + businessTypeBits //时间戳位于序列号左侧12+3+2+5=22位
	sequenceMask       = -1 ^ (-1 << sequenceBits)            //左移12位 值为(2^12)-1  序列掩码保证1毫秒内可以生成4096个id
	workerMask         = maxWorkerId << sequenceBits
	dataCenterMask     = maxDataCenter << dataCenterShift
	bizTypeMask        = maxBusinessType << businessTypeShift
	maxNextIdsNum      = 1000 //一次批量获取id的最大数目
)

var (
	idcExceedErr    = errors.New(fmt.Sprintf("idc id must be between 0 ~ %d", maxDataCenter))
	workerExceedErr = errors.New(fmt.Sprintf("worker id must be between 0 ~ %d", maxWorkerId))
	bizTypeExceed   = status.New(100000, "biz_type必须在0~31之间")
	timeBackwards   = status.New(100001, "服务器时钟回拨")
)

type IdWorker struct {
	sequence      int64      //12bit   序列数
	lastTimestamp int64      //42bit时间戳 秒级 可以用139年  64-12-4-6=42 第一位符号位 剩余41位 2^42-1/(3600*24*365*1000)约为139.4年
	workerId      int64      //4bit 部署在哪台机器上
	mutex         sync.Mutex //互斥锁 保证1毫秒内生成序列号不溢出
	idcId         int64      //idc
	maxIdNum      int64      //批量获取最大ID数
}

type sid struct {
	DataCenter int64
	WorkerId   int64
	BizType    int64
	Time       int64
	Sequence   int64
}

// NewIdWorker new a snowflake id generator object.
func New(idcId int64, workerId int64, maxIdNum int64) *IdWorker {
	if workerId > maxWorkerId || workerId < 0 {
		panic(workerExceedErr)
	}

	if idcId > maxDataCenter || idcId < 0 {
		panic(idcExceedErr)
	}

	if maxIdNum >= maxNextIdsNum || maxIdNum < 0 {
		maxIdNum = maxNextIdsNum
	}

	worker := &IdWorker{
		workerId: workerId,
		idcId:    idcId,
		maxIdNum: maxIdNum,
	}

	return worker
}

// timeGen generate a unix Millisecond.
func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// tilNextMillis spin wait till next second.
func tilNexSecond(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}

// NextIds pump num ids one time
func (id *IdWorker) NextIds(ctx context.Context, businessType, num int64) ([]int64, error) {
	if num > id.maxIdNum || num < 0 {
		num = id.maxIdNum
	}

	if businessType > maxBusinessType || businessType < 0 {
		return nil, bizTypeExceed
	}

	ids := make([]int64, num)
	id.mutex.Lock()
	for i := int64(0); i < num; i++ {
		timestamp := timeGen()
		if timestamp < id.lastTimestamp {
			log.Error(ctx, "", log.Int64("current_timestamp", timestamp), log.Int64("last_timestamp", id.lastTimestamp))
			id.mutex.Unlock()
			return nil, timeBackwards
		}
		if id.lastTimestamp == timestamp {
			id.sequence = (id.sequence + 1) & sequenceMask
			if id.sequence == 0 {
				timestamp = tilNexSecond(id.lastTimestamp)
			}
		} else {
			rand.Seed(int64(time.Now().Nanosecond()))
			//extend: 百库百表
			id.sequence = rand.Int63n(100)
		}
		id.lastTimestamp = timestamp

		ids[i] = ((timestamp - epoch) << timestampLeftShift) | (businessType << businessTypeShift) | (id.idcId << dataCenterShift) | (id.workerId << workerIdShift) | id.sequence
	}
	id.mutex.Unlock()

	return ids, nil
}

//NextId pump one id one time
func (id *IdWorker) NextId(ctx context.Context, businessType int64) (int64, error) {
	if businessType > maxBusinessType || businessType < 0 {
		return 0, bizTypeExceed
	}

	id.mutex.Lock()

	timestamp := timeGen()
	if timestamp < id.lastTimestamp {
		log.Error(ctx, "", log.Int64("current_timestamp", timestamp), log.Int64("last_timestamp", id.lastTimestamp))
		id.mutex.Unlock()
		return 0, timeBackwards
	}

	if id.lastTimestamp == timestamp {
		id.sequence = (id.sequence + 1) & sequenceMask
		if id.sequence == 0 {
			timestamp = tilNexSecond(id.lastTimestamp)
		}
	} else {
		rand.Seed(int64(time.Now().Nanosecond()))
		//extend: 百库百表
		id.sequence = rand.Int63n(100)
	}

	id.lastTimestamp = timestamp

	pumpId := ((timestamp - epoch) << timestampLeftShift) | (businessType << businessTypeShift) | (id.idcId << dataCenterShift) | (id.workerId << workerIdShift) | id.sequence

	id.mutex.Unlock()

	return pumpId, nil
}

//parse snowflake id
func (id *IdWorker) ParseId(ctx context.Context, parseId int64) (s *sid) {
	s = &sid{}
	s.Time = (parseId >> timestampLeftShift) + epoch
	s.BizType = parseId & bizTypeMask >> businessTypeShift
	s.DataCenter = parseId & dataCenterMask >> dataCenterShift
	s.WorkerId = parseId & workerMask >> workerIdShift
	s.Sequence = parseId & sequenceMask

	return
}

func (id *IdWorker) GetWorkerId() int64 {
	return id.workerId
}

func (id *IdWorker) GetDataCenter() int64 {
	return id.idcId
}

func (id *IdWorker) GetCurrentTimestamp() int64 {
	return timeGen()
}
