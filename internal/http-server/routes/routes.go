package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mo0Oonnn/password-generator/internal/http-server/handlers"
)

func Routes(logger *slog.Logger) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.URLFormat)

	mux.Get("/password/{size}/{useSpecialChars}", handlers.CreateGeneratePasswordHandler(logger))

	return mux
}
