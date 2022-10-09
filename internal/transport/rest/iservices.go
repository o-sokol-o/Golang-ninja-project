package rest

import (
	"context"
	"time"

	"github.com/tarkovskynik/Golang-ninja-project/internal/domain"
)

// Интерфейсы должны объявляться на том уровне абстракции (в том файле),
// где они используются, а не на том где реализуется

//go:generate mockgen -source=iservices.go -destination=mocks/mock.go

// Определение интерфейсов к сущностям бизнес логики

type Users interface {
	SignUp(ctx context.Context, inp domain.SignUpInput) error
	SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error)
	ParseToken(token string) (int, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
	GetRefreshTokenTTL() time.Duration
}

type FilesService interface {
	Upload(ctx context.Context, input domain.File) (string, error)
	GetFiles(ctx context.Context, id int) ([]domain.File, error)
	StoreFileInfo(ctx context.Context, input domain.File) error
}
