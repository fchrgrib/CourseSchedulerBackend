package routers

import (
	"docker_training.com/controller/cs"
	"github.com/gin-gonic/gin"
)

func CSRouters(router *gin.Engine) {
	router.POST("/course_select", cs.CourseSelectController)
}
