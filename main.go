package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "./")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":80")
}
