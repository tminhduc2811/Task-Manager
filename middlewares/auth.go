package middlewares

import (
	. "../common"
	. "../models"
	"../repository"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func AuthenticationFunc(c *gin.Context) (interface{}, error) {
	var loginModel	*LoginValidate
	var userModel	*User
	if err := c.ShouldBindJSON(&loginModel); err != nil {
		return nil, err
	}
	config, err := ReadConfig("./config.yml")
	if err != nil {
		return nil, err
	}
	dbConnect, err := mgo.Dial(config.DatabaseAddr)
	defer dbConnect.Close()
	userRepo := repository.NewMgoUserRepository(dbConnect, config.DatabaseName, config.UserCollection)
	userModel, err = userRepo.FindOne(bson.M{"email": loginModel.Email})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(userModel.HashPassword, []byte(loginModel.Password))
	if err != nil {
		return nil, err
	}
	userModel.PassWord = ""
	userModel.HashPassword = nil
	return &userModel, nil
}

func AuthorizationFunc(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*User); ok && v.Role == "admin" {
		return true
	}
	return false
}

func UnauthorizedFunc(c *gin.Context, code int, message string)  {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
	})
}

func AuthenticationMiddleware() *jwt.GinJWTMiddleware {
	authenMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:			"Authentication",
		Key:			[]byte("Admin Secret Key"),
		Timeout:		30 * time.Minute,
		MaxRefresh:		30 * time.Minute,
		IdentityKey: "role",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					jwt.IdentityKey: v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator:	AuthenticationFunc,
		Authorizator:	AuthorizationFunc,
		Unauthorized:	UnauthorizedFunc,
		SendCookie:		true,
		SecureCookie:	false, //non HTTPS dev environment
		CookieHTTPOnly:	true,
		CookieDomain:	"localhost:8080",
		CookieName:		"token", //default jwt
		TokenLookup:	"cookie:token",
	})
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}
	return authenMiddleware
}