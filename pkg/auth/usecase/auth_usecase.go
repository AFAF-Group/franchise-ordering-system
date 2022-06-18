package usecase

import (
	"errors"
	"time"

	"afaf-group.com/common/utils"
	"afaf-group.com/domain"
	"afaf-group.com/domain/models"
	"afaf-group.com/domain/repository"
	"afaf-group.com/domain/request"
	"afaf-group.com/domain/usecase"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	dberr "github.com/ytnobody/gomysqlerror/v80error"
)

type authUseCase struct {
	authMySQLRepository repository.AuthMySQLRepository
}

func NewAuthUseCase(authMySQLRepository repository.AuthMySQLRepository) usecase.AuthUseCase {
	return &authUseCase{
		authMySQLRepository: authMySQLRepository,
	}
}

func (u authUseCase) Login(ctx echo.Context, loginRequest *request.AuthRequest) (*models.Auth, error) {
	user, err := u.authMySQLRepository.FindOneByEmail(ctx, loginRequest)
	if err != nil {
		return nil, err
	}

	if err := utils.NewBCrypt().CheckPasswordHash(loginRequest.Password, user.Password); err != nil {
		return nil, err
	}
	tokenDetails, createTokenErr := u.CreateToken(ctx, user)
	if createTokenErr != nil {
		return nil, err
	}

	return tokenDetails, err
}

func (u authUseCase) CreateToken(ctx echo.Context, user *models.User) (*models.Auth, error) {
	if user == nil {
		return nil, errors.New("error: user not found")
	}
	expiresAt := time.Now().Add(12 * time.Hour).Unix()
	claims := models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Subject:   user.Email,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(domain.SECRET_JWT))
	if err != nil {
		return nil, err
	}

	tokenDetails := models.Auth{
		AccessToken: tokenString,
		UserID:      user.ID,
		ExpiresAt:   expiresAt,
	}
	return &tokenDetails, nil
}

func (u authUseCase) Register(ctx echo.Context, registerRequest *request.AuthRequest) error {
	passwordHash, err := utils.NewBCrypt().HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    registerRequest.Email,
		Password: passwordHash,
	}

	errRegister := u.authMySQLRepository.Register(ctx, &user)
	if errRegister != nil {
		if dberr.IsServerErrorDupEntry(errRegister) {
			return domain.ErrEmailAlreadyExists
		}
		return errRegister
	}
	return nil
}
