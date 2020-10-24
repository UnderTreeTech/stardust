package service

import (
	"apps/stardust/pb/v1/stardust"
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var s = &Service{}

func TestServiceGetUniqueIds(t *testing.T) {
	Convey("GetUniqueIds", t, func() {
		var (
			ctx = context.Background()
			req *stardust.IdReq
		)
		Convey("When everything goes positive", func() {
			reply, err := s.GetUniqueIds(ctx, req)
			Convey("Then err should be nil.reply should not be nil.", func() {
				So(err, ShouldBeNil)
				So(reply, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceGetUniqueId(t *testing.T) {
	Convey("GetUniqueId", t, func() {
		var (
			ctx = context.Background()
			req = &stardust.IdReq{}
		)
		Convey("When everything goes positive", func() {
			reply, err := s.GetUniqueId(ctx, req)
			Convey("Then err should be nil.reply should not be nil.", func() {
				So(err, ShouldBeNil)
				So(reply, ShouldNotBeNil)
			})
		})
	})
}

func TestServiceParseId(t *testing.T) {
	Convey("ParseId", t, func() {
		var (
			ctx = context.Background()
			req = &stardust.ParseReq{}
		)
		Convey("When everything goes positive", func() {
			reply, err := s.ParseId(ctx, req)
			Convey("Then err should be nil.reply should not be nil.", func() {
				So(err, ShouldBeNil)
				So(reply, ShouldNotBeNil)
			})
		})
	})
}
