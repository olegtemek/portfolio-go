package rest

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/olegtemek/portfolio-go/docs"
	"github.com/olegtemek/portfolio-go/internal/config"
	"github.com/olegtemek/portfolio-go/internal/middleware/logger"
	"github.com/olegtemek/portfolio-go/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	Log      *slog.Logger
	Cfg      *config.Config
	Handler  *chi.Mux
	Services *service.Service
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		Log:      log,
		Cfg:      cfg,
		Handler:  chi.NewRouter(),
		Services: services,
	}
}

func (h *Handler) Init() *http.Server {

	h.Handler.Use(middleware.RequestID)
	h.Handler.Use(middleware.URLFormat)
	h.Handler.Use(middleware.Recoverer)
	h.Handler.Use(logger.NewMiddleware(h.Log))
	h.Handler.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
	}))

	fs := http.FileServer(http.Dir("web/assets"))
	h.Handler.Handle("/web-assets/*", http.StripPrefix("/web-assets/", fs))
	h.Handler.Mount("/swagger", httpSwagger.WrapHandler)

	h.InitAllRoutes()

	srv := &http.Server{
		Addr:        h.Cfg.Address,
		ReadTimeout: h.Cfg.Timeout,
		Handler:     h.Handler,
	}

	return srv
}

func (h *Handler) InitAllRoutes() {

	h.Handler.Group(func(r chi.Router) {
		r.Get("/", h.Index)
	})
	h.Handler.Group(func(auth chi.Router) {

		auth.Use(middleware.BasicAuth("auth", map[string]string{
			h.Cfg.User: h.Cfg.Password,
		}))

		auth.Route("/info", func(r chi.Router) {
			r.Post("/", h.CreateOrUpdateInfo)
		})

		auth.Route("/experience", func(r chi.Router) {
			r.Post("/", h.CreateExperience)
			r.Patch("/{id}", h.UpdateExperience)
			r.Delete("/{id}", h.DeleteExperience)
		})

		auth.Route("/project", func(r chi.Router) {
			r.Post("/", h.CreateProject)
			r.Patch("/{id}", h.UpdateProject)
			r.Delete("/{id}", h.DeleteProject)
		})

		auth.Route("/stack", func(r chi.Router) {
			r.Post("/", h.CreateStack)
			r.Delete("/{id}", h.DeleteStack)
		})
	})
}
