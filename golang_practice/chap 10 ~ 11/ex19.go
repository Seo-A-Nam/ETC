// 11.함수 - 오름차순 정렬
package main

import "fmt"

func bubbleSort(nums ...int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func inputNums() []int {
	var nums []int
	var temp, n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&temp)
		nums = append(nums, temp)
	}
	return nums
}

func outputNums(nums ...int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nums[i])
	}
}

func main() {
	nums := inputNums()
	bubbleSort(nums...)
	outputNums(nums...)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81417/%EC%98%A4%EB%A6%84%EC%B0%A8%EC%88%9C-%EC%A0%95%EB%A0%AC
