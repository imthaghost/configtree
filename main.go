package main

import (
	"github.com/imthaghost/configtree/server"
)

// application entrypoint
func main() {
	// create a new server, the db will connect in new
	// (we allow the ability to pass a db in case you want to use a test database or mock)
	server := server.New(nil)
	// start the server on port :8080
	server.Start(":8080")
}
