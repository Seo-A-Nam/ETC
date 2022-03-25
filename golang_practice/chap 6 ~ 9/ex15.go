// 09.제어문 - 두 수를 더하면 99
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			if i+j == 99 && i%10 == j/10 {
				fmt.Printf("%02d + %02d = 99\n", i, j)
			}
		}
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/274619/%EB%91%90-%EC%88%98%EB%A5%BC-%EB%8D%94%ED%95%98%EB%A9%B4-99
