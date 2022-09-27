package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/tarkovskynik/Golang-ninja-project/internal/domain"
)

const (
	authorizationHeader = "Authorization"
	ctxUserID           = "userID"
)

func (h *Handler) getTokenFromRequest(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", domain.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", domain.ErrInvalidAuthHeader
	}

	if len(headerParts[1]) == 0 {
		return "", domain.ErrEmptyToken
	}

	return headerParts[1], nil
}

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := h.getTokenFromRequest(c)
		if err != nil {
			logError("authMiddleware", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "parse token error"})
			return
		}

		id, err := h.usersService.ParseToken(token)
		if err != nil {
			logError("authMiddleware", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "accessToken invalid or expired"})
			return
		}
		logrus.Infof("id %d", id)
		c.Set(ctxUserID, id)
		c.Next()
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(ctxUserID)
	if !ok {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
