package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/mo0Oonnn/password-generator/internal/lib/api/response"
	"github.com/mo0Oonnn/password-generator/internal/lib/random"
)

func CreateGeneratePasswordHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const operation = "handlers.CreateMakePasswordHandler"

		log = log.With(
			slog.String("operation", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		size := chi.URLParam(r, "size")
		sizeInt, err := strconv.Atoi(size)
		if err != nil || sizeInt < 4 {
			log.Info("invalid size")
			render.JSON(w, r, response.Error("invalid size"))
			return
		}

		password := random.CreateRandomPassword(sizeInt)

		log.Info("password generated")
		response.PasswordResponseOK(w, r, password, sizeInt)
	}
}
