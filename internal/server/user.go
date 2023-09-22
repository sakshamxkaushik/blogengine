package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakshamxkaushik/blogengine_armur/internal/store"
)

func Signup(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error binding the user"})
		return
	}
	store.Users = append(store.Users, user)
	ctx.JSON(http.StatusOK, gin.H{
		"user": user.Username,
		"msg":  "signed up sucefully",
		"jwt":  "123456789",
	})
}

func Signin(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	for _, u := range store.Users {
		if u.Username == user.Username && u.Password == user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"user": u.Username,
				"msg":  "Signed in successfully.",
				"jwt":  "123456789",
			})
			return
		}
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}
