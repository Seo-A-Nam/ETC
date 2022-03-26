// 12.클로저 - 동전 정리
package main

import "fmt"

func calCoin(value int) func() int {
	sum := value
	return func() int { //클로저
		return sum
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
		return
	}
	add10 := calCoin(10 * coin10)
	add50 := calCoin(50 * coin50)
	add100 := calCoin(100 * coin100)
	add500 := calCoin(500 * coin500)
	totalmoney := add10() + add50() + add100() + add500()
	fmt.Println(totalmoney)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/173738/%EB%8F%99%EC%A0%84-%EC%A0%95%EB%A6%AC
