package authorization

import (
	"binginx.com/brush/cmd/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HeaderParams struct {
	Authorization string `header:"accessToken" binding:"required,min=20"`
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerParams := HeaderParams{}
		if err := ctx.ShouldBindHeader(&headerParams); err != nil {
			ctx.Abort()
			response.MkResponse(ctx, http.StatusBadRequest, "missing jwt token", nil)
			return
		}
	}
}
