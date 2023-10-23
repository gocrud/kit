package ginx

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var DefaultCorsOptions = CorsMiddlewareOptions{
	AllowOrigins:  []string{"*"},
	Headers:       []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token"},
	ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	Methods:       []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
	MaxAge:        0,
}

type CorsMiddlewareOptions struct {
	AllowOrigins  []string
	Headers       []string
	ExposeHeaders []string
	Methods       []string
	MaxAge        int
}

// CorsMiddleware 函数返回一个中间件，用于处理跨域请求。
func CorsMiddleware(opts CorsMiddlewareOptions) gin.HandlerFunc {
	origins := strings.Join(opts.AllowOrigins, ",")
	headers := strings.Join(opts.Headers, ",")
	methods := strings.Join(opts.Methods, ",")
	exposeHeaders := strings.Join(opts.ExposeHeaders, ",")
	maxAge := strconv.Itoa(opts.MaxAge)
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if len(origin) != 0 {
			ctx.Header("Access-Control-Allow-Origin", origins)
			ctx.Header("Access-Control-Allow-Headers", headers)
			ctx.Header("Access-Control-Allow-Methods", methods)
			ctx.Header("Access-Control-Expose-Headers", exposeHeaders)
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Max-Age", maxAge)
		}
		//允许类型校验
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}

// RecoveryMiddleware Panic处理函数
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error(
					err.(string),
					slog.String("method", ctx.Request.Method),
					slog.String("path", ctx.Request.URL.Path),
				)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg":  "系统错误",
					"code": "system",
				})
			}
		}()
		ctx.Next()
	}
}
