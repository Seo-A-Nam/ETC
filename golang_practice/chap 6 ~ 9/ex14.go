// 09.제어문 - 구구단2
package main

import "fmt"

func main() {
	for i := 2; i <= 9; i++ {
		if i%2 == 1 {
			for j := 1; j <= i; j++ {
				fmt.Printf("%d x %d = %d\n", i, j, i*j)
			}
			fmt.Println()
		}
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/274212/%EA%B5%AC%EA%B5%AC%EB%8B%A82
