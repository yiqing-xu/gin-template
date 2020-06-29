/*
@Time : 2020/6/29 9:05
@Author : xuyiqing
@File : auth.py
*/

package jwt

import (
	"gin-template/conf"
	"gin-template/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义jwt载荷
type UserClaims struct {
	jwt.StandardClaims
	ID uint `json:"id"`
	Username string `json:"username"`
}

// 根据payload查询user返回
func (c *UserClaims) GetUserByID() *models.Account {
	var user models.Account
	models.DB.Model(&models.Account{}).First(&user, c.ID)
	if user.ID > 0 {
		return &user
	} else {
		return nil
	}
}

// 生成jwt token字符串
func GenToken(id uint, username string) (string, error) {
	expiredTime := time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims := UserClaims{
		jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
		id,
		username,
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(conf.JwtSecretKey.SecretKey))
	return token, err
}

// 验证token合法性
func ValidateJwtToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSecretKey.SecretKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 断言设定ctx的当前用户
func AssertUser(ctx *gin.Context) *models.Account {
	currentUser, isExists := ctx.Get("CurrentUser")
	if !isExists {
		return nil
	}
	user, ok := currentUser.(*models.Account)
	if ok {
		return user
	} else {
		return nil
	}
}
