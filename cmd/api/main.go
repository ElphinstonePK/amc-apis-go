package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		meezan := v1.Group("/meezan")
		{
			meezan.POST("/receiveToken")
		}

	}
	router.Run("localhost:8080")
}
