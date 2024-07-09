package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(ctx *gin.Context)
}
