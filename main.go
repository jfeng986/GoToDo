package main

import (
	"GoToDo/config"
	"GoToDo/internal/router"
)

func main() {
	r := router.NewRouter()
	r.Run(config.HttpPort)
}
