package auth

import (
    "fmt"
    "net/http"
    //"encoding/json"
    "time"
    "github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
    Token *jwt.Token
}

type AuthUser interface {
    VerifyUser(r *http.Request) bool
    VerifyAdmin(r *http.Request) bool
}

//jwt token signing string
var signString []byte = []byte("oboeMadSauceSupremeGammaTrainSuprippp$%&*%^@@@vsmsoiosvh")

/*

Assign new token to the struct instead of returning a string?

*/

//func (j *JwtToken) VerifyUser(r *http.Request) bool {
//
//}

func (j *JwtToken) GenerateToken(Id int64, Role int) string {
	//jwt token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["userId"] = Id
	token.Claims["userRole"] = Role
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(signString)

	if err != nil {
		//handle error
		return ""
	}

	//return token
	return tokenString
}

func (j *JwtToken) ParseToken(r *http.Request) (*jwt.Token, error) {
	myToken :=r.Header.Get("User-Token")

	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {

		//verify signing type
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return signString, nil
	})

	//return error or decoded token
	if err == nil && token.Valid {
		return token, nil
	} else {
		return nil, err
	}
}
