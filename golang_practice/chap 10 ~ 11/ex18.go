// 10.컬렉션 - 중간고사 평균 점수
package main

import "fmt"

func main() {
	m := make(map[string]int)
	var avg float32
	var subject string
	var score int

	for {
		fmt.Scanf("%s", &subject)
		if subject == "0" {
			break
		}
		fmt.Scanf("%d", &score)
		m[subject] = score
	}
	var sum int
	var count int
	for str, _ := range m {
		sum += m[str]
		fmt.Printf("%s %d\n", str, m[str])
		count++
	}
	avg = float32(sum) / float32(count)
	fmt.Printf("%.2f\n", avg)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81411/%EC%A4%91%EA%B0%84%EA%B3%A0%EC%82%AC-%ED%8F%89%EA%B7%A0-%EC%A0%90%EC%88%98
