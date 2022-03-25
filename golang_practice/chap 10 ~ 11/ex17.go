// 10.컬렉션 - 가장 긴 이름
package main

import "fmt"

func main() {
	names := make([]string, 3)
	var name string

	for name != "1" {
		fmt.Scanln(&name)
		names = append(names, name)
	}
	var result string
	var max_len int
	for _, str := range names {
		if max_len < len(str) {
			max_len = len(str)
			result = str
		}
	}
	fmt.Println(result, len(result))
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81407/%EA%B0%80%EC%9E%A5-%EA%B8%B4-%EC%9D%B4%EB%A6%84
