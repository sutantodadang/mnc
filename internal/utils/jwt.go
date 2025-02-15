package utils

import (
	"mnc/internal/repositories"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func GenerateJwt(data repositories.SelectOneUserByPhoneNumberRow, exp int64) (token string, err error) {
	claimsJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp,
		"sub": data.UserID.String(),
	})

	token, err = claimsJwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	return
}
