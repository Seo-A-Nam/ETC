// 18. 채널 - 고루틴 실습2
package main

import "fmt"

func add(a int, b int, ch chan int) {
	ch <- a + b
}

func main() {
	var num1, num2 int
	ch := make(chan int, 1)

	fmt.Scanln(&num1, &num2)

	go add(num1, num2, ch)
	fmt.Println(<-ch)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/388997/%EA%B3%A0%EB%A3%A8%ED%8B%B4-%EC%8B%A4%EC%8A%B52
