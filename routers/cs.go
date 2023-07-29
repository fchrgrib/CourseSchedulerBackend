package routers

import (
	"docker_training.com/controller/cs"
	"github.com/gin-gonic/gin"
)

func CSRouters(router *gin.Engine)  {
	router.GET("/course_select",cs.CourseSelectController)
}
