// 15. defer와 panic() - 엘리베이터
package main

import "fmt"

func main() {
	var person []string
	var name string

	for {
		fmt.Scanln(&name)
		if name == "0" {
			break
		} else {
			person = append(person, name)
		}
	}

	for _, str := range person {
		defer fmt.Println(str)
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/387048/%EC%97%98%EB%A6%AC%EB%B2%A0%EC%9D%B4%ED%84%B0
