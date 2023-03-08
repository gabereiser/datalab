package network

import (
	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	fiber *fiber.App
}

func NewWebServer() *WebServer {
	return &WebServer{
		fiber: fiber.New(),
	}
}

func (s *WebServer) Listen(addr string) {

}
