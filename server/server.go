package server

import (
	"github.com/imthaghost/configtree/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents the server.
type Server struct {
	e  *echo.Echo
	db *mongo.Database
}

// NewServer creates a Server and instantiates the DB if not provided
func New(database *mongo.Database) *Server {
	if database == nil {
		database = db.Connect()
	}

	return &Server{
		e:  echo.New(),
		db: database,
	}
}

// Start Server functionality
func (s *Server) Start(port string) {
	// register routes
	s.Routes()
	// start the server
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close stops the Server
func (s *Server) Close() {
	// stop the server
	s.e.Close()
}
