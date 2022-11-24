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

func UserInfo(ctx *gin.Context) {
	headerParams := authorization.HeaderParams{}
	ctx.ShouldBindHeader(&headerParams)
	logs.Logger.Infof("Token:[%v]")
	err, userInfo := service.UserInfo(&model.UserInfo{
		Token: headerParams.Authorization,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, userInfo)
	return
}

func Score(ctx *gin.Context) {
	headerParams := authorization.HeaderParams{}
	ctx.ShouldBindHeader(&headerParams)
	logs.Logger.Infof("Token:[%v]")
	err, score := service.Score(&model.UserInfo{
		Token: headerParams.Authorization,
	}, &config.Params{
		Params: map[string]string{
			"qtime": strconv.Itoa(time.Now().Nanosecond()),
		},
	})
	if err != nil {
		response.MkResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.MkResponse(ctx, http.StatusOK, response.Success, score)
	return
}
