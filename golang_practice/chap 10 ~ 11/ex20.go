// 11.함수 - 아이패드를 사주는 조건
package main

import "fmt"

func inputNums() []int {
	nums := make([]int, 5)
	var temp int

	for i := 0; i < 5; i++ {
		fmt.Scanln(&temp)
		nums = append(nums, temp)
	}
	return nums
}

func calExam(arr []int) (int, int, int) {
	var total, over90, under70 int

	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] >= 90 {
			over90++
		} else if arr[i] < 70 {
			under70++
		}
		total += arr[i]
	}
	return total, over90, under70
}

func printResult(f func([]int) (int, int, int), scores []int) {
	var result bool = true

	total, over90, under70 := f(scores)
	if total < 400 {
		fmt.Println("총점이 400점 미만입니다.")
		result = false
	} else if over90 < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	} else if under70 > 0 {
		fmt.Println("70점 미만 과목이 있습니다.")
		result = false
	}

	if result == true {
		fmt.Println("아이패드를 살 수 있습니다.")
	} else {
		fmt.Println("아이패드를 살 수 없습니다.")
	}
}

func main() {
	scores := inputNums()
	printResult(calExam, scores)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81680/%EC%95%84%EC%9D%B4%ED%8C%A8%EB%93%9C%EB%A5%BC-%EC%82%AC%EC%A3%BC%EB%8A%94-%EC%A1%B0%EA%B1%B4
