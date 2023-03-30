package main

import (
	"book-simple-rest-api/routers"

	_ "github.com/lib/pq"
)

func main() {

	// endpoints
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
