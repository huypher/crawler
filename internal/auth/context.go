package auth

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

const (
	jwtTokenKey = "jwtToken"
	uidKey      = "uid"
)

type UID struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
}

func UIDFromContext(ctx context.Context) UID {
	val := ctx.Value(uidKey)
	if uid, ok := val.(UID); ok {
		return uid
	}

	return UID{}
}

func withUID(ctx *gin.Context, jwtClaims jwt.MapClaims) *gin.Context {
	uid := UID{}

	if userID, ok1 := jwtClaims["user_id"]; ok1 {
		uid.ID, _ = uuid.Parse(fmt.Sprintf("%s", userID))
	}

	if username, ok1 := jwtClaims["username"]; ok1 {
		if v, ok2 := username.(string); ok2 {
			uid.Username = v
		}
	}

	if name, ok1 := jwtClaims["name"]; ok1 {
		if v, ok2 := name.(string); ok2 {
			uid.Name = v
		}
	}

	ctx.Set(uidKey, uid)

	return ctx
}
