package router

import (
	"net/http"
	"prisma-golang-vue/pkg/types/routes"
	HomeHandler "prisma-golang-vue/src/controllers/home"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
