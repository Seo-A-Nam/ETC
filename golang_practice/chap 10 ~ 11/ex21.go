// 11.함수 - 역학적 에너지
package main

import "fmt"

const g float32 = 9.8

type kinEnergy func(float32, float32) float32
type potEnergy func(float32, float32) float32

func calMechEnergy(f func(float32, float32) float32, a float32, b float32) float32 {
	result := f(a, b)
	return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)

	kinEnergy := func(m float32, v float32) float32 {
		return (0.5 * m * v * v)
	}
	potEnergy := func(m float32, h float32) float32 {
		return (m * g * h)
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)

	fmt.Printf("%g %g %g\n", ke, pe, ke+pe)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81679/%EC%97%AD%ED%95%99%EC%A0%81-%EC%97%90%EB%84%88%EC%A7%80
