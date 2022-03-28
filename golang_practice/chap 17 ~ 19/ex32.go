// 18. 채널 - 메시지 전송
package main

import (
	"fmt"
	"time"
)

func main() {
	var tmp string
	var flag int

	ch := make(chan string)
	sec := 10
	go sendMessage(ch)

	for {
		if sec == 0 || flag == 1 {
			break
		}
		select {
		case tmp = <-ch:
			fmt.Printf("%s 메시지가 발송되었습니다.\n", tmp)
			flag = 1
		default:
			time.Sleep(time.Second)
			fmt.Printf("%d 초 안에 메시지를 입력하시오.\n", sec)
			sec--
		}
	}
	close(ch)
}

func sendMessage(ch chan string) {
	var msg string
	fmt.Scanln(&msg)
	ch <- msg
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/389394/%EB%A9%94%EC%8B%9C%EC%A7%80-%EC%A0%84%EC%86%A1
