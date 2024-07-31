package main

import (
	"fmt"

	"github.com/legend123213/go_togo/Task04/router"
)


func main(){
	root := router.Api()
	fmt.Println("")
	err :=root.Run(":8000")
	if err!=nil {
		fmt.Println("errror")
	}else{
		fmt.Println("work fine")
	}

}