// 03.연산자 - 간단한 덧셈과 곱셈
package main

import "fmt"

func main() {
	var num1, num2, num3, result int

	fmt.Scanln(&num1, &num2, &num3)
	result = num1*num2 + num3

	fmt.Printf("%d x %d + %d = %d", num1, num2, num3, result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/269813/%EA%B0%84%EB%8B%A8%ED%95%9C-%EB%8D%A7%EC%85%88%EA%B3%BC-%EA%B3%B1%EC%85%88
