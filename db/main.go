package main

import (
	"fmt"
	"log"

	"github.com/miuer/ncepu-work/db/miu-bookstore/conf"
	"github.com/miuer/ncepu-work/db/miu-bookstore/controller/gin"
)

func main() {
	fmt.Println("Hollow World!")

	log.SetFlags(log.Ldate | log.Lshortfile)

	conf.InitConfig()

	gin.Init()

}
