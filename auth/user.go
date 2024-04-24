package auth

import (
	"context"
	"fmt"
)

type User struct {
	UUID string

	// could contain any additional info about the auth (permissions, name, profiles etc.)
}

func UserFromContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value("user").(User)
	if !ok {
		return User{}, fmt.Errorf("user missing from context")
	}

	return user, nil
}
