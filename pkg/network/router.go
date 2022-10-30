package network

import (
	"github.com/chalfel/rate-limiter/pkg/attempt"
	"github.com/chalfel/rate-limiter/pkg/rule"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine           *gin.Engine
	RateLimitHandler *RateLimitHandler
}

func NewRouter() (*Router, error) {
	router := gin.Default()

	return &Router{
		Engine:           router,
		RateLimitHandler: &RateLimitHandler{receiver: attempt.NewReceiver(attempt.NewInMemoryAttemptStore(), *rule.NewRuler(rule.NewInMemoryRuleStore()))},
	}, nil
}

func (r *Router) RegisterRoutes() {
	r.Engine.POST("/", r.RateLimitHandler.Try)

}
