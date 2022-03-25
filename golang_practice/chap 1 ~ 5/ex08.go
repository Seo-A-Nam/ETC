// 05.콘솔 출력과 입력 함수 - 신상정보 입력과 출력
package main

import "fmt"

func main() {
	var id, name string
	var height float32

	fmt.Scanf("%s", &id)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &height)

	fmt.Printf("주민등록번호 앞자리는 %s, 뒷자리는 %s, 이름은 %s입니다.\n그리고 키는 %.2f입니다.", id[:6], id[7:], name, height)
}

// https://edu.goorm.io/learn/lecture/2010/%ED%95%9C-%EB%88%88%EC%97%90-%EB%81%9D%EB%82%B4%EB%8A%94-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EC%B4%88/lesson/270291/%EC%8B%A0%EC%83%81%EC%A0%95%EB%B3%B4-%EC%9E%85%EB%A0%A5%EA%B3%BC-%EC%B6%9C%EB%A0%A5
