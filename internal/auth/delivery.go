package auth

import (
	"net/http"

	"github.com/pghuy/dobi-oms/domain"

	"github.com/gin-gonic/gin"
)

const (
	loginPath    = "/login"
	userInfoPath = "/user"
)

type authDelivery struct {
	authUsecase domain.AuthUsecase
}

func NewAuthDelivery(authUsecase domain.AuthUsecase) *authDelivery {
	return &authDelivery{
		authUsecase: authUsecase,
	}
}

func (d *authDelivery) Handler(c *gin.RouterGroup) {
	group := c.Group("auth")
	group.Handle(http.MethodPost, loginPath, d.loginHandler())
}
