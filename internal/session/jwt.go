package session

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var secretKey []byte

const cookieName = "session"

var algo = jwt.SigningMethodHS256

var invalidTokenErr = fmt.Errorf("invalid token")

func init() {
	secretKeyStr, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		log.Printf("no secretKey in .env file")
		return
	}
	secretKey = []byte(secretKeyStr)
}

type JwtSession struct {
	UsersDb *gorm.DB
}

func (jwtSession *JwtSession) Check(r *http.Request) (*Session, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, invalidTokenErr
	}
	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{algo.Name}))
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, invalidTokenErr
	}
	subjId, err := token.Claims.GetSubject()
	if err != nil {
		return nil, invalidTokenErr
	}
	id, err := strconv.Atoi(subjId)
	if err != nil {
		return nil, invalidTokenErr
	}
	session := Session{Id: uint(id)}
	return &session, nil

}

func (jwtSession *JwtSession) Create(w http.ResponseWriter, user UserInterface) error {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(user.GetId()),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, _ := claims.SignedString(secretKey)

	// TODO: add csrf token
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    tokenString,
		Expires:  time.Now().Add(90 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusCreated)
	respBody, _ := json.Marshal(user)
	w.Write(respBody)

	return nil
}
