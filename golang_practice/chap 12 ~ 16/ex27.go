// 15. defer와 panic() - 중간고사 평균 점수2
package main

import "fmt"

func average() float64 {

	var num int
	fmt.Scanln(&num)

	if num <= 0 {
		panic("잘못된 과목 수입니다.")
	}
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 {
			panic("잘못된 점수입니다.")
		}
		total += score
	}
	avg := float64(total) / float64(num)
	return avg
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			main()
		}
	}()
	result := average()
	fmt.Println(result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/387049/%EC%A4%91%EA%B0%84%EA%B3%A0%EC%82%AC-%ED%8F%89%EA%B7%A0-%EC%A0%90%EC%88%982
