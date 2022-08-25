package domain

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"zero/internal/auth/domain/common"
)

type User struct {
	ID        int
	UID       string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(uid string, email string, name string) User {
	return User{
		UID:   uid,
		Email: email,
		Name:  name,
	}
}

type UserClaim struct {
	jwt.RegisteredClaims
	User
}

func GenerateUserToken(user User, signingKey []byte, expiryDuration time.Duration, issuer string) (string, common.Error) {
	claim := &UserClaim{
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expiryDuration)},
			Issuer:    issuer,
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		user,
	}

	// Generate Signed JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", common.NewError(common.ErrorCodeInternalProcess, err, common.WithMsg("failed to generate token"))
	}

	return signedToken, nil
}

func ParseUserFromToken(signedToken string, signingKey []byte) (*User, common.Error) {
	token, err := jwt.ParseWithClaims(signedToken, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok && e.Errors == jwt.ValidationErrorExpired {
			msg := "token is expired"
			return nil, common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg))
		} else {
			return nil, common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg("failed to parse token"))
		}
	}

	if !token.Valid {
		msg := "invalid token"
		return nil, common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg))
	}

	claim, ok := token.Claims.(*UserClaim)
	if !ok {
		msg := "failed to parse claim"
		return nil, common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg))
	}

	return &claim.User, nil
}
