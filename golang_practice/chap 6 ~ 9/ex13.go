// 08.분기문 switch - 안좋은 계산기
package main

import "fmt"

func main() {
	var sel int
	var num1, num2, result float32

	fmt.Scanln(&sel)
	fmt.Scanf("%f %f", &num1, &num2)

	switch sel {
	case 1:
		result = num1 + num2
	case 2:
		if num1 > num2 {
			result = num1 - num2
		} else {
			result = num2 - num1
		}
	case 3:
		result = num1 * num2
	case 4:
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("잘못된입력입니다.")
		}
	default:
		fmt.Println("잘못된입력입니다.")
	}
	fmt.Printf("%.1f", result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/80289/%EC%95%88%EC%A2%8B%EC%9D%80-%EA%B3%84%EC%82%B0%EA%B8%B0
