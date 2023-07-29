package utils

import "docker_training.com/models"

func MaxScoreSKS(courses []models.Course, creditsTotal int) (float32, []models.Course) {
	lenCourses := len(courses)

	if lenCourses == 0 || creditsTotal == 0 {
		return 0, []models.Course{}
	}

	dp := make([][]float32, lenCourses+1)
	for i := range dp {
		dp[i] = make([]float32, creditsTotal+1)
	}

	for i := 0; i <= lenCourses; i++ {
		for w := 0; w <= creditsTotal; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if courses[i-1].CreditTotal <= w {
				dp[i][w] = max(ParseIndex(courses[i-1].Expectation)*float32(courses[i-1].CreditTotal)+dp[i-1][w-courses[i-1].CreditTotal], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	res := dp[lenCourses][creditsTotal]
	w := creditsTotal
	listMaxCourse := make([]models.Course, 0)

	for i := lenCourses; i > 0 && res > 0; i-- {
		if res == dp[i-1][w] {
			continue
		} else {
			listMaxCourse = append(listMaxCourse, courses[i-1])

			res = res - ParseIndex(courses[i-1].Expectation)*float32(courses[i-1].CreditTotal)
			w = w - courses[i-1].CreditTotal
		}
	}

	return dp[lenCourses][creditsTotal], listMaxCourse
}

func ParseIndex(index string) float32 {
	if index == "A" {
		return 4
	}
	if index == "AB" {
		return 3.5
	}
	if index == "B" {
		return 3
	}
	if index == "BC" {
		return 2.5
	}
	if index == "C" {
		return 2
	}
	if index == "D" {
		return 1
	}
	if index == "E" {
		return 0
	}
	return 0
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
