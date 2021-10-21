package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func RegisterGet(c *gin.Context) {
	//返回html
	fmt.Println("hello")
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}
