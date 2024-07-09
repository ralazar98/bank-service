package http

import "net/http"

func TechRouting(m http.ServeMux) {
	g := m.Group("/api/v1")

	g.POST("/user/balance"
}


