package service

import (
	"authService/configs"
	"authService/internal/jwt/model"
	role "authService/internal/role/service"
	user "authService/internal/user/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

type Service struct {
	userService     user.Service
	roleService     role.Service
	privateKey      string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	redis           *redis.Client
}

func NewService(cnf configs.JWTConfig, userService user.Service, roleService role.Service, redis *redis.Client) Service {
	return Service{
		userService:     userService,
		roleService:     roleService,
		privateKey:      cnf.Secret,
		accessTokenTTL:  time.Duration(cnf.AccessTokenTTL) * time.Minute,
		refreshTokenTTL: time.Duration(cnf.RefreshTokenTTL) * time.Minute,
		redis:           redis,
	}
}

func (s Service) GenerateAccessToken(userID uuid.UUID, role string, permissions []string) (string, error) {
	tokenID := s.generateTokenUUID(userID, "refresh")
	claims := jwt.MapClaims{
		"type":        "access",
		"jti":         tokenID,
		"sub":         userID.String(),
		"role":        role,
		"permissions": permissions,
		"ip":          22,
		"device_info": "iphone_13",
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(s.accessTokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.privateKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s Service) GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	tokenID := s.generateTokenUUID(userID, "refresh").String()
	claims := jwt.MapClaims{
		"type": "refresh",
		"jti":  tokenID,
		"sub":  userID.String(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(s.refreshTokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.privateKey))
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		log.Printf("Failed to marshal claims to JSON: %v", err)
	}

	if err = s.redis.Set(ctx, tokenID, claimsJSON, s.refreshTokenTTL).Err(); err != nil {
		log.Printf("Can`t add to redis key=%s, value=%s", tokenID, string(claimsJSON))
	}

	return signedToken, nil
}

func (s Service) ValidateToken(tokenString string) (*model.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method")
		}
		return []byte(s.privateKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}

	if claims, ok := jwtToken.Claims.(*model.Claims); ok && jwtToken.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (s Service) RefreshTokens(ctx context.Context, refreshToken string) (tokens model.Tokens, err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.privateKey), nil
	})

	if err != nil || !token.Valid {
		err = fmt.Errorf("invalid refresh token")
		return
	}

	if claims["type"] != "refresh" {
		err = fmt.Errorf("provided token is not refresh")
		return
	}

	userID, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		err = fmt.Errorf("invalid user ID from token")
		return
	}

	foundUser, err := s.userService.GetById(userID)
	if err != nil {
		return
	}

	permissions, err := s.roleService.GetRolePermCodesById(foundUser.RoleID)
	if err != nil {
		return
	}

	newAccessToken, err := s.GenerateAccessToken(userID, foundUser.Role.Name, permissions)
	if err != nil {
		return
	}

	newRefreshToken, err := s.GenerateRefreshToken(ctx, userID)
	if err != nil {
		return
	}

	tokens = model.Tokens{AccessToken: newAccessToken, RefreshToken: newRefreshToken}
	return
}

func (s Service) generateTokenUUID(userID uuid.UUID, tokenType string) uuid.UUID {
	timestamp := time.Now().UnixNano()
	rawID := userID.String() + tokenType + strconv.FormatInt(timestamp, 10)

	newUUID := uuid.NewSHA1(uuid.New(), []byte(rawID))
	return newUUID
}
