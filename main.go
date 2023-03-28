package main

import "book-simple-rest-api/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
