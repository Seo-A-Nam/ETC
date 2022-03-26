// 13.구조체와 메소드 - 역학적 에너지
package main

import "fmt"

const g float32 = 9.8

type energy struct {
	m float32
	h float32
	v float32
}

func (e energy) ke() float32 {
	return (0.5) * e.m * e.v * e.v
}

func (e energy) pe() float32 {
	return g * e.m * e.h
}

func main() {
	var n int

	fmt.Scanln(&n)
	items := make([]energy, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&items[i].m, &items[i].v, &items[i].h)
	}
	for i := 0; i < n; i++ {
		fmt.Println(items[i].ke(), items[i].pe(), items[i].ke()+items[i].pe())
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/316489/%EC%97%AD%ED%95%99%EC%A0%81-%EC%97%90%EB%84%88%EC%A7%802
