# Course Web Backend

## Overview
This is backend for course web that build with go programming language and mysql as database
the backend contain algorithm to solve user input to give a maximum score that user can have.
i use dynamic programming to solve the problem

## Framework And Tech
- Gin (Framework)
- Docker
- Mysql

## What is Dynamic Programming?
Dynamic programming is a computational technique used to solve problems by breaking them down into smaller overlapping subproblems and efficiently solving each subproblem only once, storing its solution for future reference. It involves solving these subproblems in a bottom-up manner, starting from the simplest ones and progressively building solutions for more complex ones. By avoiding redundant computations through memoization or tabulation, dynamic programming significantly improves the efficiency of solving complex problems, making it particularly useful for optimization, combinatorial, and decision-making challenges across various fields such as computer science, mathematics, economics, and engineering.

## Algorithm Analysis
i compare the value from the bottom to the top to get the maximum score with algorithm below i use dynamic programming with the maximum score of minimum credits first and compare with minimum credits+1 until maximum credits to get the maximum of maximum score.

## How to Run Program
- ensure you have docker desktop in your device
- run docker desktop
- pull the project in this link https://github.com/fchrgrib/CourseSchedulerBackend.git
- open the project with your IDE or command prompt and make a command `docker compose build`
- after that you can run the project with command `docker compose up -d` in terminal
- you can see the program in http://localhost:8080

```go
func MaximumCourses(course []models.CourseDB, maximumCredits int, minimumCredits int) (int, []models.CourseDB) {
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
```

## Reference
1. Docker Documentation: https://docs.docker.com/
2. Gin Official Documentation: https://gin-gonic.com/docs/
3. Golang Official Website: https://golang.org/
4. MySQL Documentation: https://dev.mysql.com/doc/
