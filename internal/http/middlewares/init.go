package middlewares

import "mnc/internal/repositories"

type Middleware struct {
	repo repositories.Querier
}

func NewMiddleware(repo repositories.Querier) *Middleware {
	return &Middleware{repo: repo}
}
