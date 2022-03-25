// 04.자료형 - 강제 형 변환
package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Scanln(&num1, &num2, &num3)

	a := float32(num1)
	b := uint(num2)
	c := int64(num3)
	fmt.Printf("%T, %f\n%T, %d\n%T, %d\n", a, a, b, b, c, c)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/270251/%EA%B0%95%EC%A0%9C-%ED%98%95-%EB%B3%80%ED%99%98
