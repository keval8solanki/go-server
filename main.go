package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Ping struct {
  Message string `json:"message"`
}

func ping(ctx *gin.Context) {

	var pingMessage = Ping{
		Message: "pong",
	}

	ctx.IndentedJSON(http.StatusOK, pingMessage)
}

func main() {
	router := gin.Default()

	router.GET("/ping", ping)

	router.Run()
}
