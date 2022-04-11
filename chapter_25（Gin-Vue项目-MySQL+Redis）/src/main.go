package main

import "routes"

func main() {
	panic(routes.InitRoutes().Run(":8080"))
}
