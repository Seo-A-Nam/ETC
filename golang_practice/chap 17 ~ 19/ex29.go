// 17. 고루틴 - 고루틴 실습
package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a int, b int, result *int, w *sync.WaitGroup) {
	defer w.Done()

	*result = a + b
	time.Sleep(time.Second)
}

func main() {
	var num1, num2 int = 10, 5
	var result int

	wait := new(sync.WaitGroup)
	wait.Add(1)
	go add(num1, num2, &result, wait)

	wait.Wait()
	fmt.Println(result)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/388600/%EA%B3%A0%EB%A3%A8%ED%8B%B4-%EC%8B%A4%EC%8A%B5
