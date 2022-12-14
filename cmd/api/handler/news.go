package handler

import (
	"binginx.com/brush/cmd/api/middleware/authorization"
	"binginx.com/brush/cmd/api/response"
	"binginx.com/brush/config"
	"binginx.com/brush/internal/logs"
	"binginx.com/brush/model"
	"binginx.com/brush/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func News(ctx *gin.Context) {
	headerParams := authorization.HeaderParams{}
	ctx.ShouldBindHeader(&headerParams)
	logs.Logger.Infof("Token:[%v]")
	params := &config.Params{
		Params: map[string]string{
			"qtime":    strconv.Itoa(time.Now().Nanosecond()),
			"columnId": "ecosysNews",
			"pageNum":  "1",
			"pageSize": "200",
			"needAll":  "true",
		},
	}
	err, ecosysNewsIds := service.News(&model.UserInfo{
		Token: headerParams.Authorization,
	}, params)
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, ecosysNewsIds)
	return
}

func View(ctx *gin.Context) {
	headerParams := authorization.HeaderParams{}
	ctx.ShouldBindHeader(&headerParams)
	logs.Logger.Infof("Token:[%v]")
	contentId, ok := ctx.GetQuery("contentId")
	if !ok {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return

	}
	params := &config.Params{
		Params: map[string]string{
			"qtime":     strconv.Itoa(time.Now().Nanosecond()),
			"contentId": contentId,
		},
	}
	err := service.View(&model.UserInfo{
		Token: headerParams.Authorization,
	}, params)
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}

func DoBehavior(ctx *gin.Context) {
	headerParams := authorization.HeaderParams{}
	ctx.ShouldBindHeader(&headerParams)
	logs.Logger.Infof("Token:[%v]")
	resourceId, ok := ctx.GetQuery("resourceId")
	if !ok {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return

	}
	flag, ok := ctx.GetQuery("flag")
	if !ok {
		response.MkResponse(ctx, http.StatusBadRequest, response.ParamInvalid, nil)
		return

	}
	params := &config.Params{
		Params: map[string]string{
			"qtime":      strconv.Itoa(time.Now().Nanosecond()),
			"resourceId": resourceId,
			"flag":       flag,
		},
	}
	err := service.DoBehavior(&model.UserInfo{
		Token: headerParams.Authorization,
	}, params)
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, nil)
	return
}
