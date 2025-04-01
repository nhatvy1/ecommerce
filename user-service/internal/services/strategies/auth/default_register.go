package auth_strategies

import (
	"context"
	"fmt"
	"net/http"
	"user-service/internal/model"
)

type DefaultRegisterStrategy struct{}

func (s *DefaultRegisterStrategy) Register(ctx context.Context, body *model.UserRegister) (int, int, error) {
	return -1, http.StatusFound, fmt.Errorf("register failed")
}
