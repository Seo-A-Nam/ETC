// 05.콘솔 출력과 입력 함수 - 정돈된 표
package main

import "fmt"

func main() {
	name := "이름"
	major := "전공학과"
	grade := "학년"
	fmt.Printf("%-8s%-14s%5s\n", name, major, grade)

	name = "유현수"
	major = "전자공학"
	grade = "3"
	fmt.Printf("%-8s%-14s%5s\n", name, major, grade)

	name = "김윤욱"
	major = "컴퓨터공학"
	grade = "4"
	fmt.Printf("%-8s%-14s%5s\n", name, major, grade)

	name = "김나영"
	major = "미술교육학"
	grade = "2"
	fmt.Printf("%-8s%-14s%5s\n", name, major, grade)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/270261/%EC%A0%95%EB%8F%88%EB%90%9C-%ED%91%9C
