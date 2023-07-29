package cdm

import (
	"database/sql"
	"docker_training.com/db"
	"docker_training.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDepartment(c *gin.Context) {
	_db := db.Connect()

	var department []models.Department

	row, err := _db.Query("SELECT * FROM department")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"department": department,
			"status":     err,
		})
		return
	}
	defer func(row *sql.Rows) {
		if err := row.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"department": department,
				"status":     err,
			})
			return
		}
	}(row)

	for row.Next() {
		var each models.Department
		if err := row.Scan(&each.Id, &each.DepartmentName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"department": []models.Department{},
				"status":     err,
			})
			return
		}
		department = append(department, each)
	}

	c.JSON(http.StatusOK, gin.H{
		"department": department,
		"status":     "ok",
	})
	return
}

func AddOneDepartment(c *gin.Context) {
	var department models.Department

	_db := db.Connect()

	if err := c.ShouldBind(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	if _, err := _db.Exec("INSERT INTO department(id, department_name) VALUES(?,?)", department.Id, department.DepartmentName); err != nil {
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

func AddMultipleDepartment(c *gin.Context) {
	var departments []models.Department

	_db := db.Connect()

	if err := c.ShouldBind(&departments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	for _, department := range departments {
		if _, err := _db.Exec("INSERT INTO department(id, department_name) VALUES(?,?)", department.Id, department.DepartmentName); err != nil {
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

func DeleteDepartment(c *gin.Context) {
	_db := db.Connect()

	id := c.Param("id")

	if _, err := _db.Exec("DELETE FROM department WHERE id = ?", id); err != nil {
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
