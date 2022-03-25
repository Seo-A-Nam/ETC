// 02.변수와 상수 - 잘못된 신상정보
package main

import "fmt"

const (
	name string = "kim"
	RRN  string = "800101-1000000"
	job
)

func main() {
	fmt.Println(name + " " + RRN + " " + job)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/269664/%EC%9E%98%EB%AA%BB%EB%90%9C-%EC%8B%A0%EC%83%81%EC%A0%95%EB%B3%B4
