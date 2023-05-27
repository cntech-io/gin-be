package ginbe

import "github.com/gin-gonic/gin"

type RouteMethod string

const (
	GET    RouteMethod = "GET"
	POST   RouteMethod = "POST"
	PUT    RouteMethod = "PUT"
	PATCH  RouteMethod = "PATCH"
	DELETE RouteMethod = "DELETE"
)

type resource struct {
	path        string
	method      RouteMethod
	handler     gin.HandlerFunc
	middlewares []gin.HandlerFunc
}

func NewResource(method RouteMethod, path string) *resource {
	return &resource{
		path:   path,
		method: method,
	}
}

func (a *resource) Handler(h gin.HandlerFunc) *resource {
	a.handler = h
	return a
}

func (a *resource) AddMiddleware(m gin.HandlerFunc) *resource {
	a.middlewares = append(a.middlewares, m)
	return a
}
