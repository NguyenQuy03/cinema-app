package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultUserEmail = "quy@gmail.com"
)

func TestGenerateAccessToken(t *testing.T) {
	provider := &JWTProvider{}

	token, expTime, err := provider.GenerateAccessToken(defaultUserEmail, false)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, ExpireAccessTokenInSeconds, expTime)
}

func TestGenerateRefreshToken(t *testing.T) {
	provider := &JWTProvider{}

	token, expTime, err := provider.GenerateRefreshToken(defaultUserEmail, false)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, ExpireRefreshTokenInSeconds, expTime)
}

func TestValidateToken(t *testing.T) {
	provider := &JWTProvider{}
	token, _, _ := provider.GenerateAccessToken(defaultUserEmail, false)

	validatedToken, err := provider.ValidateToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, validatedToken)
}
