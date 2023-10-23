package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gocrud/kit/web/ginx"
)

func Console(r ginx.FilterResult) {
	fmt.Println(r.Body(), r.Status())
}

func main() {
	r := gin.Default()
	ginx.SetFilter(Console)
	r.Use(ginx.RecoveryMiddleware())
	r.GET("/success", ginx.Handle(success))
	r.GET("/failure", ginx.Handle(failure))
	r.GET("/panic", ginx.Handle(panic1))
	r.GET("/null", ginx.Handle(null))

	r.Run()
}

// status code 200
func success(c *gin.Context) ginx.Result {
	return ginx.Success("success")
}

// status code 400
func failure(c *gin.Context) ginx.Result {
	err := errors.New("400错误")
	return ginx.Bad(err)
}

// status code 500
func panic1(c *gin.Context) ginx.Result {
	var err = errors.New("panic")
	return ginx.Internal(err)
}

// status code 204
func null(c *gin.Context) ginx.Result {
	return nil
}
