// 07.조건문 if/else - 7과 9의 배수
package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%9 == 0 {
			fmt.Printf("%d ", i)
		} else if i%9 != 0 && i%7 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/80286/7%EA%B3%BC-9%EC%9D%98-%EB%B0%B0%EC%88%98
