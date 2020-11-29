package context

import (
	"context"

	"github.com/kazijawad/PhotoGallery/models"
)

type privateKey string

const (
	userKey privateKey = "user"
)

// WithUser returns a copy of the context in which the user is available.
func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// User returns a user from the context if it exists.
func User(ctx context.Context) *models.User {
	if temp := ctx.Value(userKey); temp != nil {
		if user, ok := temp.(*models.User); ok {
			return user
		}
	}
	return nil
}
