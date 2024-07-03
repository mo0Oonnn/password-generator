package routes_test

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/mo0Oonnn/password-generator/internal/http-server/routes"
	"github.com/mo0Oonnn/password-generator/internal/lib/logger/slogdiscard"
)

func TestRoutes(t *testing.T) {
	mux := routes.Routes(slogdiscard.NewDiscardLogger())

	switch T := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("unexpected type %T", T)
	}
}
