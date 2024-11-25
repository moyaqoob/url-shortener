package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/",func (c *gin.Context)  {
		c.JSON(200,gin.H{
			"message":"hey go url-shortener",
		})
	})
	err:= r.Run(":8080")
	if err!=nil{
		panic(fmt.Sprintf("Failed to start web server. Error %v",err))
	}
}