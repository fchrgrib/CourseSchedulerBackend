package main

import (
	"docker_training.com/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "hello world",
		})
	})

	routers.CMDRouters(router)
	routers.CSRouters(router)

	if err := router.Run(); err != nil {
		return
	}
}

//package main
//
//import (
//	"docker_training.com/algorithm"
//	"docker_training.com/models"
//	"fmt"
//)
//
//// 8 -> A
//// 7 -> AB
//// 6 -> B
//// 5 -> BC
//// 4 -> C
//// 2 -> D
//// 0 -> E
//type matkul struct {
//	score int
//	sks   int
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func getMaxScore(Matkul []matkul, bobotSks int) (maxScore int, listMatkul []int) {
//	lenMatkul := len(Matkul)
//	// bikin table dp[lenMatkul][bobotSks]
//	dp := make([][]int, lenMatkul+1)
//	for i := range dp {
//		dp[i] = make([]int, bobotSks+1)
//	}
//
//	// isi table dp
//	for i := 0; i <= lenMatkul; i++ {
//		for w := 0; w <= bobotSks; w++ {
//			if i == 0 || w == 0 {
//				dp[i][w] = 0
//			} else if Matkul[i-1].sks <= w {
//				// Perbandingan antara jika ambil matkul ke-i atau tidak
//				dp[i][w] = max(Matkul[i-1].score*Matkul[i-1].sks+dp[i-1][w-Matkul[i-1].sks], dp[i-1][w])
//			} else {
//				dp[i][w] = dp[i-1][w]
//			}
//		}
//	}
//
//	// Cari matkul yang diambil dengan algoritma backtrack
//	res := dp[lenMatkul][bobotSks]
//	w := bobotSks
//	listMaxMatkul := make([]int, 0)
//
//	for i := lenMatkul; i > 0 && res > 0; i-- {
//		if res == dp[i-1][w] {
//			continue
//		} else {
//			listMaxMatkul = append(listMaxMatkul, i-1)
//
//			res = res - Matkul[i-1].score*Matkul[i-1].sks
//			w = w - Matkul[i-1].sks
//		}
//	}
//
//	return dp[lenMatkul][bobotSks], listMaxMatkul
//}
//
//func main() {
//	// Matkul
//	//Matkul := []matkul{
//	//	{score: 8, sks: 2}, // 0
//	//	{score: 8, sks: 2}, // 1
//	//	{score: 8, sks: 2}, // 2
//	//	{score: 8, sks: 2}, // 3
//	//	{score: 8, sks: 2}, // 4
//	//	{score: 8, sks: 3}, // 5
//	//	{score: 8, sks: 2}, // 6
//	//	{score: 7, sks: 3}, // 7
//	//	{score: 2, sks: 1}, // 8
//	//	{score: 2, sks: 1}, // 9
//	//}
//	//
//	//// Bobot sks
//	//bobotSks := 17
//	//
//	//maxScore, listMatkul := getMaxScore(Matkul, bobotSks)
//	//IP := float64(maxScore) / float64(2*bobotSks)
//	//fmt.Println("Max Score:", IP)
//	//fmt.Println("List Matkul:", listMatkul)
//
//	course := []models.Course{
//		{
//			Id:          "AAA",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAB",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAC",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAD",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAE",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAF",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAG",
//			Name:        "name",
//			CreditTotal: 3,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAH",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "AB",
//		},
//		{
//			Id:          "AAI",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//		{
//			Id:          "AAJ",
//			Name:        "name",
//			CreditTotal: 2,
//			Semester:    4,
//			Major:       "Informatics",
//			Expectation: "A",
//		},
//	}
//
//	maxs, list := algorithm.MaximumCourses(course, 20, 15)
//	credit := 0
//	for _, value := range list {
//		credit += value.CreditTotal
//	}
//
//	fmt.Printf("Maximum Score: %.2f\n", float32(maxs)/float32(credit))
//	fmt.Println(credit)
//
//}
//
