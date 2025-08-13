package api

import (
	"Blog_Backend/model"
	"Blog_Backend/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func InitUsersRouter(router *gin.Engine) {
	v1 := router.Group("v1/users")
	v1.POST("/register", register)
	v1.POST("/login", login)
}

// 用户注册功能
func register(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBind(&user); err == nil {
		bcryptPwd, err2 := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  err2.Error(),
			})
			return
		}
		user.Password = string(bcryptPwd)
		repo.CreateUser(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Success",
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  "Internal Server Error",
	})
	return
}

// 用户登录功能
func login(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBind(&user); err == nil {
		dbPwd := repo.FindUser(user.Email)
		if "" != dbPwd {
			err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(user.Password))
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "Unauthorized",
				})
				return
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":       user.ID,
				"username": user.UserName,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})
			tokenString, err := token.SignedString([]byte("MyBlog"))
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "Success",
				"data": gin.H{
					"token": tokenString,
				},
			})
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Unauthorized",
			})
			return
		}
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  "Internal Server Error",
	})
	return
}
