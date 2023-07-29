package routers

import (
	"docker_training.com/controller/cdm"
	"github.com/gin-gonic/gin"
)

func CMDRouters(router *gin.Engine) {
	router.GET("/course", cdm.GetCourse)
	router.GET("/department", cdm.GetDepartment)
	router.GET("/major", cdm.GetMajor)

	router.POST("/add_course", cdm.AddOneCourse)
	router.POST("/add_department", cdm.AddOneDepartment)
	router.POST("/add_major", cdm.AddOneMajor)

	router.POST("/add_courses", cdm.AddMultipleCourse)
	router.POST("/add_departments", cdm.AddMultipleDepartment)
	router.POST("/add_majors", cdm.AddMultipleMajor)

	router.DELETE("/course/:id", cdm.DeleteCourse)
	router.DELETE("/department/:id", cdm.DeleteDepartment)
	router.DELETE("/major/:id", cdm.DeleteMajor)
}
