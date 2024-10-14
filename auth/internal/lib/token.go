package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	AccessTokenPayload struct {
		Username string
		ID       string
		Data     map[string]interface{}
	}

	RefreshTokenPayload struct {
		Username string
		ID       string
	}
)

var jwt_secret = os.Getenv("JWT_SECRET")

func (p *AccessTokenPayload) GenAccessToken(data ...map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	duration, err := intFromEnv("JWT_ACCESS_DURATION_MIN")
	if err != nil {
		return "", err
	}

	setCommonClaims(claims, p.Username, p.ID, duration)
	claims["data"] = data

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(jwt_secret))
	return tokenString, nil
}

func (p *RefreshTokenPayload) GenRefershToken() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	duration, err := intFromEnv("JWT_REFRESH_DURATION_MIN")
	if err != nil {
		return "", err
	}

	setCommonClaims(claims, p.Username, p.ID, duration)

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(jwt_secret))
	return tokenString, nil
}

func setCommonClaims(claims jwt.MapClaims, username string, id string, duration int) {
	claims["exp"] = time.Now().Add(time.Duration(duration) * time.Minute).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = username
	claims["id"] = id
}

func ParseAccessToken(tokenString string) (*AccessTokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["id"].(string)
	username, _ := claims["username"].(string)
	data, _ := claims["data"].(map[string]interface{})
	return &AccessTokenPayload{
		ID:       id,
		Username: username,
		Data:     data,
	}, nil
}

func ParseRefreshToken(tokenString string) (*RefreshTokenPayload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	id, _ := claims["id"].(string)
	username, _ := claims["username"].(string)
	return &RefreshTokenPayload{
		ID:       id,
		Username: username,
	}, nil
}

func ValidateToken(Token string) error {
	token, err := jwt.Parse(Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}
	return nil
}

func GenTokenPair(accessPayload AccessTokenPayload, refreshPayload RefreshTokenPayload) (string, string, error) {
	access, err := accessPayload.GenAccessToken()
	if err != nil {
		return "", "", err
	}
	refresh, err := refreshPayload.GenRefershToken()
	if err != nil {
		return "", "", err
	}
	return access, refresh, nil
}
