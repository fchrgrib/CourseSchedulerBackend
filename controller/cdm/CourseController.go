package cdm

import (
	"database/sql"
	"docker_training.com/db"
	"docker_training.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(c *gin.Context) {
	_db := db.Connect()

	var course []models.Course

	row, err := _db.Query("SELECT * FROM course")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"course": course,
			"status": err,
		})
		return
	}
	defer func(row *sql.Rows) {
		if err := row.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"course": course,
				"status": err,
			})
			return
		}
	}(row)

	for row.Next() {
		var each models.Course
		if err := row.Scan(&each.Id, &each.Name, &each.CreditTotal, &each.Semester, &each.Major, &each.Expectation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"course": []models.Course{},
				"status": err,
			})
			return
		}
		course = append(course, each)
	}

	c.JSON(http.StatusOK, gin.H{
		"course": course,
		"status": "ok",
	})
	return
}

func AddOneCourse(c *gin.Context) {
	var course models.Course

	_db := db.Connect()

	if err := c.ShouldBind(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	var departmentId string
	if err := _db.QueryRow("SELECT department_id FROM major WHERE major_name = ?", course.Major).Scan(&departmentId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	if _, err := _db.Exec("INSERT INTO course(id, course_name, credit_total, semester, major_name, expectation, department_id) VALUES(?,?,?,?,?,?,?)", course.Id, course.Name, course.CreditTotal, course.Semester, course.Major, course.Expectation, departmentId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func AddMultipleCourse(c *gin.Context) {
	var courses []models.Course

	_db := db.Connect()

	if err := c.ShouldBind(&courses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	for _, course := range courses {
		var departmentId string
		if err := _db.QueryRow("SELECT department_id FROM major WHERE major_name = ?", course.Major).Scan(&departmentId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err,
			})
			return
		}
		if _, err := _db.Exec("INSERT INTO course(id, course_name, credit_total, semester, major_name, expectation, department_id) VALUES(?,?,?,?,?,?,?)", course.Id, course.Name, course.CreditTotal, course.Semester, course.Major, course.Expectation, departmentId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func DeleteCourse(c *gin.Context) {
	_db := db.Connect()

	id := c.Param("id")

	if _, err := _db.Exec("DELETE FROM course WHERE id = ?", id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
