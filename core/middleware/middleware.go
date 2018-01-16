package middleware

import (
	"net/http"

	"github.com/tokopedia/panics"
)

func Handle(h http.HandlerFunc) http.HandlerFunc {
	return panics.CaptureHandler(h)
}
