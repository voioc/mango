package handler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
 * @Author: Cedar
 * @Date: 2021-01-26 16:10:01
 * @LastEditors: Cedar
 * @LastEditTime: 2021-01-26 17:46:29
 * @FilePath: /Mango/app/handler/Login.go
 */

// Login 登陆
func Login(c *gin.Context) {

	claims := &jwt.StandardClaims{
		Id: "1001",
	}
	now := time.Now()
	claims.IssuedAt = now.Unix() - 300
	claims.ExpiresAt = now.Add(time.Hour * 24 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generateJwtToken, _ := token.SignedString([]byte("cxeAh1!#OJsDLN4R"))

	c.JSON(200, gin.H{"code": 0, "message": "success", "data": gin.H{"token": generateJwtToken}})
	return

}
