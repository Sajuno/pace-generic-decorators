package auth

import (
	"context"
	"fmt"
)

type User struct {
	uuid string

	authorized bool
}

func (u User) Authorized() bool {
	return u.Authorized()
}

func (u User) UUID() string {
	return u.uuid
}

func UserFromContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value("user").(User)
	if !ok {
		return User{}, fmt.Errorf("user missing from context")
	}

	return user, nil
}
