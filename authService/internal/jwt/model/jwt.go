package model

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}
type RedisJWTInfo struct {
	Claims     Claims
	DeviceInfo string `json:"device_info"`
	Ip         string `json:"ip"`
}

func (c *Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	if c.ExpiresAt.IsZero() {
		return nil, fmt.Errorf("expiration time is not set")
	}
	return c.ExpiresAt, nil
}

func (c *Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	if c.IssuedAt.IsZero() {
		return nil, fmt.Errorf("issued at time is not set")
	}
	return nil, nil
}

func (c *Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.NotBefore, nil
}

func (c *Claims) GetIssuer() (string, error) {
	return c.Issuer, nil
}

func (c *Claims) GetSubject() (string, error) {
	return c.Subject, nil
}

func (c *Claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.Audience, nil
}
