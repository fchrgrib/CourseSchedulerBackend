package cs

import (
	"database/sql"
	"docker_training.com/algorithm"
	"docker_training.com/db"
	"docker_training.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CourseSelectController(c *gin.Context) {
	var courses []models.Course
	var courseSelect models.CourseSelect

	_db := db.Connect()

	if err := c.ShouldBind(&courseSelect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"course": courses,
			"score":  0,
			"status": err,
		})
		return
	}
	var row *sql.Rows
	var err error

	if courseSelect.Choose == "no" {
		row, err = _db.Query("SELECT * FROM course WHERE semester <= ? AND major_name = ?", courseSelect.Semester, courseSelect.Major)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"course": courses,
				"score":  0,
				"status": err,
			})
			return
		}
	} else {
		var departmentId string
		if err := _db.QueryRow("SELECT department_id FROM major WHERE major_name = ?", courseSelect.Major).Scan(&departmentId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err,
			})
			return
		}

		row, err = _db.Query("SELECT * FROM course WHERE semester <= ? AND department_id = ?", courseSelect.Semester, departmentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"course": courses,
				"score":  0,
				"status": err,
			})
			return
		}
	}

	for row.Next() {
		var each models.Course
		if err := row.Scan(&each.Id, &each.Name, &each.CreditTotal, &each.Semester, &each.Major, &each.Expectation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"course": courses,
				"score":  0,
				"status": err,
			})
			return
		}
		courses = append(courses, each)
	}

	maxScore, listCourse := algorithm.MaximumCourses(courses, courseSelect.MaximumCredits, courseSelect.MinimumCredits)
	creditTotal := 0

	for _, value := range listCourse {
		creditTotal += value.CreditTotal
	}

	c.JSON(http.StatusOK, gin.H{
		"course": listCourse,
		"score":  strconv.FormatFloat(float64(maxScore)/float64(creditTotal), 'f', 2, 64),
		"status": "ok",
	})
	return
}
