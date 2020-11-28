package main

import (
	"api-starter/routers"
	"api-starter/utils"
)

func main() {
	router := routers.GetRouter()
	port := utils.EnvVar("SERVER_PORT")
	router.Run(port)
}
