package main

import (
	"MyBlog/routers"
)

func main() {
	router := routers.InitRouter()
	//静态资源
	router.Static("/static", "./static")
	router.Run(":8081")
}

//test push
