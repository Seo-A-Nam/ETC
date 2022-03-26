// 13.구조체와 메소드 - 성적 저장 프로그램
package main

import "fmt"

type student struct {
	name   string
	gender string
	score  map[string]int
}

func newStudent() *student {
	s := student{}
	s.score = map[string]int{}
	return &s
}

func main() {
	var n, m, grade int
	var gender, name, subject string
	var s [](*student)

	fmt.Scanln(&n, &m)

	for i := 0; i < n; i++ {
		stu := newStudent()
		s = append(s, stu)
		fmt.Scanln(&name, &gender)
		s[i].name, s[i].gender = name, gender
		for j := 0; j < m; j++ {
			fmt.Scanln(&subject, &grade)
			s[i].score[subject] = grade
		}
	}
	for _, temp := range s {
		fmt.Println("----------")
		fmt.Printf("%s %s\n", temp.name, temp.gender)
		for index, val := range temp.score {
			fmt.Println(index, val)
		}
	}
	fmt.Println("----------")
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/316444/%EC%84%B1%EC%A0%81-%EC%A0%80%EC%9E%A5-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%A8
