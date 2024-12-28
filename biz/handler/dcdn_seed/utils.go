package dcdn_seed

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"seed.manager/biz/model/dcdn_seed"
)

type ResponseExtentd struct {
	*dcdn_seed.BaseResponse
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *app.RequestContext, data interface{}) {
	resp := &ResponseExtentd{BaseResponse: &dcdn_seed.BaseResponse{}, Data: data}
	c.JSON(consts.StatusOK, resp)
}

func ResponseError(c *app.RequestContext, ctx context.Context, code dcdn_seed.ErrorCode, message string) {
	logger.CtxErrorf(ctx, "%s failed code=%d message=%s", c.Request.URI().String(), code, message)
	resp := dcdn_seed.BaseResponse{Code: int32(code), Message: message}
	c.JSON(consts.StatusOK, &resp)
}
