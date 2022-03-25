// 07.조건문 if/else - 두 수의 차
package main

import "fmt"

func main() {
	var num1, num2, result int

	fmt.Scanf("%d %d", &num1, &num2)

	if num1 > num2 {
		result = num1 - num2
	} else {
		result = num2 - num1
	}

	fmt.Printf("%d\n", result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/273749/%EB%91%90-%EC%88%98%EC%9D%98-%EC%B0%A8
