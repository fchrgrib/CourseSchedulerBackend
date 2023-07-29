package models

type Course struct {
	Id          string `json:"id" gorm:"column:id;primaryKey;index;type:int"`
	Name        string `json:"name" gorm:"column:course_name"`
	CreditTotal int    `json:"credit_total" gorm:"column:credit_total"`
	Semester    int    `json:"semester"`
	Major       string `json:"major_name" gorm:"column:major"`
	Expectation string `json:"expectation" gorm:"column:expectation"`
}

type Major struct {
	Id           string `json:"id"`
	DepartmentID string `json:"department_id"`
	MajorName    string `json:"major_name"`
}

type Department struct {
	Id             string `json:"id"`
	DepartmentName string `json:"department_name"`
}
