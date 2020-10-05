// Package health is a general purpose health check http middleware
package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const body = `{ "alive": true }`

func Handler(path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == path {
			w.WriteHeader(200)
			w.Write([]byte(body))
		}
	}
}

func GinHandler(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := Handler(path)
		handler(c.Writer, c.Request)
	}
}
