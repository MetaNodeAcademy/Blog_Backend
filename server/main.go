package main

import (
	"Blog_Backend/config"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	config.InitTables()
	RouterInit(gin.Default())
}
