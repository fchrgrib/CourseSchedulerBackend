package cdm

import (
	"database/sql"
	"docker_training.com/db"
	"docker_training.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMajor(c *gin.Context) {
	_db := db.Connect()

	var major []models.Major

	row, err := _db.Query("SELECT * FROM major")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"major":  major,
			"status": err,
		})
		return
	}
	defer func(row *sql.Rows) {
		if err := row.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"major":  major,
				"status": err,
			})
			return
		}
	}(row)

	for row.Next() {
		var each models.Major
		if err := row.Scan(&each.Id, &each.MajorName, &each.DepartmentID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"major":  []models.Major{},
				"status": err,
			})
			return
		}
		major = append(major, each)
	}

	c.JSON(http.StatusOK, gin.H{
		"major":  major,
		"status": "ok",
	})
	return
}

func AddOneMajor(c *gin.Context) {
	var major models.Major

	_db := db.Connect()

	if err := c.ShouldBind(&major); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	if _, err := _db.Exec("INSERT INTO major(id, major_name, department_id) VALUES(?,?,?)", major.Id, major.MajorName, major.DepartmentID); err != nil {
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

func AddMultipleMajor(c *gin.Context) {
	var majors []models.Major

	_db := db.Connect()

	if err := c.ShouldBind(&majors); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	for _, major := range majors {
		if _, err := _db.Exec("INSERT INTO major(id, major_name, department_id) VALUES(?,?,?)", major.Id, major.MajorName, major.DepartmentID); err != nil {
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

func DeleteMajor(c *gin.Context) {
	_db := db.Connect()

	id := c.Param("id")

	if _, err := _db.Exec("DELETE FROM major WHERE id = ?", id); err != nil {
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
