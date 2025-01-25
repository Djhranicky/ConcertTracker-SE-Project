package main

import "github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
