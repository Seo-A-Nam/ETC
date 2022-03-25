// 10.컬렉션 - 역행렬
package main

import "fmt"

func main() {
	var mat = [2][2]int{{7, 3}, {5, 2}}

	x, y, w, z := mat[0][0], mat[0][1], mat[1][0], mat[1][1]
	d := mat[0][0]*mat[1][1] - mat[1][0]*mat[0][1]

	if d == 0 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
		coef := 1 / (x*z - y*w)
		rmat := [2][2]int{{coef * z, -1 * coef * y}, {-1 * coef * w, coef * x}}
		fmt.Printf("%d %d\n%d %d\n", rmat[0][0], rmat[0][1], rmat[1][0], rmat[1][1])
	}
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/81404/%EC%97%AD%ED%96%89%EB%A0%AC
