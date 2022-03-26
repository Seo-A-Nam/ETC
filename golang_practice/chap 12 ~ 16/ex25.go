// 14.인터페이스 - 직육면체와 원기둥

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	volume() float32
}

type Cylinder struct {
	radius, height float32
}

type Cube struct {
	len1, len2, len3 float32
}

func (c Cylinder) area() float32 {
	return c.radius * math.Pi * 2 * (c.radius + c.height)
}

func (c Cylinder) volume() float32 {
	return c.radius * c.radius * math.Pi * c.height
}

func (q Cube) area() float32 {
	return 2 * (q.len1*q.len2 + q.len1*q.len3 + q.len2*q.len3)
}

func (q Cube) volume() float32 {
	return q.len1 * q.len2 * q.len3
}

func main() {
	cy1 := Cylinder{10, 10}
	cy2 := Cylinder{4.2, 15.6}
	cu1 := Cube{10.5, 20.2, 20}
	cu2 := Cube{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(m ...geometry) {
	for _, val := range m {
		fmt.Printf("%.2f, %.2f\n", val.area(), val.volume())
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/320105/%EC%A7%81%EC%9C%A1%EB%A9%B4%EC%B2%B4%EC%99%80-%EC%9B%90%EA%B8%B0%EB%91%A5
