package main

import (
	"go_blog/model"
	"go_blog/routes"
)

func main(){
	model.InitDb()
	routes.InitRouter()
}

