package service

import (
	"authService/internal/auth/model"
	jwt "authService/internal/jwt/model"
	jwtModel "authService/internal/jwt/model"
	jwtServ "authService/internal/jwt/service"
	role "authService/internal/role/service"
	userModel "authService/internal/user/model"
	user "authService/internal/user/service"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserService user.Service
	RoleService role.Service
	JWTService  jwtServ.Service
}

func NewService(userService user.Service, roleService role.Service, jwtService jwtServ.Service) Service {
	return Service{
		UserService: userService,
		RoleService: roleService,
		JWTService:  jwtService,
	}
}

func (s Service) Register(regForm model.Register) error {
	if err := regForm.Validate(); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(regForm.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userRole, err := s.RoleService.GetById(regForm.RoleID)

	if err != nil {
		return fmt.Errorf("not found role with id %d", regForm.RoleID)
	}

	newUser := userModel.User{
		Email:    regForm.Email,
		Password: string(hash),
		Username: regForm.Username,
		Role:     userRole,
	}

	err = s.UserService.Create(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Login(context context.Context, loginForm model.Login) (tokens jwtModel.Tokens, err error) {
	if err = loginForm.Validate(); err != nil {
		return
	}

	foundUser, err := s.UserService.GetByEmail(loginForm.Email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginForm.Password))
	if err != nil {
		err = fmt.Errorf("invalid password")
		return
	}

	codes, err := s.RoleService.GetRolePermCodesById(foundUser.RoleID)
	if err != nil {
		return
	}

	accessToken, err := s.JWTService.GenerateAccessToken(foundUser.ID, codes)
	if err != nil {
		return
	}

	refreshToken, err := s.JWTService.GenerateRefreshToken(context, foundUser.ID)
	if err != nil {
		return
	}

	tokens = jwtModel.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}
	return
}

func (s Service) Auth(tokenStr string) (*jwt.Claims, error) {
	claims, err := s.JWTService.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (s Service) RefreshTokens(ctx context.Context, refreshToken string) (jwtModel.Tokens, error) {
	tokens, err := s.JWTService.RefreshTokens(ctx, refreshToken)
	if err != nil {
		return jwtModel.Tokens{}, err
	}

	return tokens, nil
}
