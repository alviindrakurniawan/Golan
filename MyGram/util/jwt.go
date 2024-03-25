package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)



func GenerateJWTToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func GetJWTClaims (tokenString string) (map[string]any,error){
	token,err:= jwt.Parse(tokenString,func(t *jwt.Token)(interface{},error){
		_,ok:= t.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,errors.New("invalid method")
		}
		return []byte("secret"),nil
	
	})
	if err!=nil{
		return nil,err
	}
	if !token.Valid{
		return nil,errors.New("invalid token")
	}
	claims,ok:= token.Claims.(jwt.MapClaims)
	if !ok {
		return nil,errors.New("invalid claims")
	}
	return claims,nil
	
}

func GetSubFromClaims(claims any) (any,error){

	mapClaims,ok:= claims.(map[string]any)
	if !ok{
		return nil,errors.New("not a map")
	}
	
	sub,ok:= mapClaims["sub"]
	if !ok{
		return nil,errors.New("not found")
	}

	return sub,nil
}

