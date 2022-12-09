package main

import (
	"fmt"
	"testing"

	"github.com/gn02963770/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do
	default:
		t.Error(fmt.Sprintf("type is notchi.Mux, type is %T", v))
	}
}
