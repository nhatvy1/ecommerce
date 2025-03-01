package main

import (
	"fmt"
	"strconv"
	"user-service/global"
	"user-service/internal/initialize"
)

func main() {
	r := initialize.Run()

	r.Run(fmt.Sprintf(":%s", strconv.Itoa(global.Config.Server.Port)))
}
