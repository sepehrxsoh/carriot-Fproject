package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewListCustomersClient(conn)
	g := gin.Default()
	g.GET("/vehicles/:payload", func(ctx *gin.Context) {
		payload, err2 := strconv.ParseUint(ctx.Param("payload"), 10, 64)
		if err2 != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		}
		req := &proto.Request{Payload: int64(payload)}
		if response, err := client.makeList(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("faild to run server: %v", err)
	}
}
