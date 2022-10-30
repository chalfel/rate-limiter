package network

import (
	"net/http"

	"github.com/chalfel/rate-limiter/pkg/attempt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RateLimitHandler struct {
	receiver *attempt.Receiver
}

type TryPayload struct {
	Ip     string `json:"ip"`
	Region string `json:"region"`
}

func NewRateLimitHandler(receiver *attempt.Receiver) *RateLimitHandler {
	return &RateLimitHandler{
		receiver: receiver,
	}
}

func (rl *RateLimitHandler) Try(c *gin.Context) {
	tp := &TryPayload{}

	if err := c.ShouldBindBodyWith(tp, binding.JSON); err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	err := rl.receiver.Try(tp.Ip, tp.Region)

	if err != nil {
		c.Status(http.StatusTooManyRequests)
		return
	}

	c.Status(http.StatusOK)
}
