package network

import (
	"embed"
	"net/http"
	"time"

	"github.com/gabereiser/datalab/config"
	"github.com/gabereiser/datalab/controllers"
	"github.com/gabereiser/datalab/controllers/auth"
	"github.com/gabereiser/datalab/controllers/workbook"
	"github.com/gabereiser/datalab/log"
	"github.com/gabereiser/datalab/security"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	jwtware "github.com/gofiber/jwt/v3"
	html "github.com/gofiber/template/html"
)

type WebServer struct {
	fiber *fiber.App
}

func NewWebServer(views embed.FS) *WebServer {
	return &WebServer{
		fiber: fiber.New(fiber.Config{
			ServerHeader: "hypercloud",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			AppName:      "datalab",
			BodyLimit:    100 * 1024 * 1024,
			Views:        html.NewFileSystem(http.FS(views), ".html"),
		}),
	}
}

func (s *WebServer) Listen(addr string) {
	err := s.fiber.Listen(addr)
	if err != nil {
		log.Err("%v", err)
	}
}

func (s *WebServer) Stop() {
	err := s.fiber.Shutdown()
	if err != nil {
		log.Err("%v", err)
	}
}

func (s *WebServer) SetupRoutes(view embed.FS) {
	/*s.fiber.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(view),
	}))*/
	s.fiber.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(view),
		PathPrefix: "public",
		Browse:     true,
	}))
	//s.fiber.Static("/assets", "public/assets")

	// public routes
	s.fiber.Get("/", controllers.IndexScreen)
	s.fiber.Get("/privacy", controllers.PrivacyScreen)
	s.fiber.Get("/terms", controllers.TermsScreen)
	s.fiber.Get("/register", auth.RegisterScreen)
	s.fiber.Post("/register", auth.RegisterHandler)
	s.fiber.Get("/login", controllers.LoginScreen)

	authGroup := s.fiber.Group("auth")
	authGroup.Post("/login", auth.LoginHandler)

	// JWT Middleware
	s.fiber.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config.SecretKey),
		ContextKey: "_auth",
	}))
	s.fiber.Use(security.NewAuthMiddleware())

	// authenticated routes
	authGroup.Get("/refresh", auth.RefreshHandler)

	apiGroup := s.fiber.Group("api/v1")
	apiGroup.Get("/workbook/list", workbook.WorkbookListHandler)
	apiGroup.Get("/workbook/:id", workbook.WorkbookGetHandler)
}
