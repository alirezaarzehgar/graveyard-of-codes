package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping/:param", func(ctx *gin.Context) {
		p, _ := ctx.Params.Get("param")
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "MF!",
			"param": p,
		})
	})
	r.Run(":8000")
}
