// 16.에러처리 - 중간고사 평균 점수3
package main

import (
	"fmt"
)

func inputSubNum() (int, error) {
	var num int
	_, err := fmt.Scanln(&num)
	if num <= 0 || err != nil {
		return 0, fmt.Errorf("잘못된 과목 수입니다.")
	}
	return num, nil
}

func average(num int) (float64, error) {
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 || score > 100 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}
		total += score
	}
	avg := float64(total) / float64(num)
	return avg, nil
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	num, err := inputSubNum()
	if err != nil {
		panic(err)
	}
	result, err := average(num)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/388044/%EC%A4%91%EA%B0%84%EA%B3%A0%EC%82%AC-%ED%8F%89%EA%B7%A0-%EC%A0%90%EC%88%983
