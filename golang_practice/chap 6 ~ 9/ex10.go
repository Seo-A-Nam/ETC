// 06.반복문 for - 빛나는 이등변삼각형
package main

import "fmt"

func main() {
	var n, i, j int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
		for j = 0; j <= i; j++ {
			if j == i {
				fmt.Print("* ")
			} else {
				fmt.Print("o ")
			}
		}
		fmt.Print("\n")
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81414/%EB%B9%9B%EB%82%98%EB%8A%94-%EC%9D%B4%EB%93%B1%EB%B3%80%EC%82%BC%EA%B0%81%ED%98%95
