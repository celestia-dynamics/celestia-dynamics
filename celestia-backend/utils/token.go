package utils

import(
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string, email string) (string,error) {
	claims := jwt.MapClaims{
		"user_id":userID,
		"email": email,
		"exp": time.Now().Add(24*time.Hour).Unix(),  //expires in one day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
} 

func ValidateToken(tokenString string) (jwt.MapClaims,error){
	token,err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		if _,ok:= token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,errors.New("unexpected signing method")
		}
		return jwtSecret,nil
	})

	if err!= nil || !token.Valid{
		return nil,errors.New("invalid token")
	}

	claims,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return nil,errors.New("could not parse claims")
	}

	return claims,nil
}