// 06.반복문 for - 구구단
package main

import "fmt"

func main() {
	var dan int

	fmt.Scanln(&dan)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/80283/%EA%B5%AC%EA%B5%AC%EB%8B%A8
