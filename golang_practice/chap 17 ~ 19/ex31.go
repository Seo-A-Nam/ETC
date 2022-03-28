// 18. 채널 - 동기 채널 실습
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 20; i++ {
			done <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i := 0; i < 20; i++ {
		<-done
		fmt.Printf("수신한 데이터 : %d\n", i)
	}
	time.Sleep(time.Second * 3)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/389824/%EB%8F%99%EA%B8%B0-%EC%B1%84%EB%84%90-%EC%8B%A4%EC%8A%B5
