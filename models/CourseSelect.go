package models

type CourseSelect struct {
	Major          string `json:"major"`
	Semester       int    `json:"semester"`
	MinimumCredits int    `json:"minimum_credits"`
	MaximumCredits int    `json:"maximum_credits"`
	Choose         string `json:"choose"`
}
