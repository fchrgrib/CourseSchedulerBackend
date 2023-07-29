package algorithm

import (
	"docker_training.com/models"
	"docker_training.com/utils"
)

func MaximumCourses(course []models.Course, maximumCredits int, minimumCredits int) (int, []models.Course) {
	if maximumCredits == minimumCredits {
		maximum, listCourse := utils.MaxScoreSKS(course, minimumCredits)
		return int(maximum), listCourse
	}

	scoreNow, listNow := utils.MaxScoreSKS(course, maximumCredits)
	scoreBottom, listBottom := MaximumCourses(course, maximumCredits-1, minimumCredits)
	if int(scoreNow) < scoreBottom {
		return scoreBottom, listBottom
	} else {
		return int(scoreNow), listNow
	}
}
