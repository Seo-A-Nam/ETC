// 의사코드를 섞어서 대략적으로라도 짜본다.


// 모든 client들은 노드0의 좌표를 기본적으로 가지고 있다. (전역이나 지연변수 초기화 등등으로 선언)

/*
type t_data struct {
	flag			int // 나는 지금 몇 번째로 참가하는 노드인가?
	node_ip[3]		string
	node_cord[4][3]	int // (x, y, z) 형식의 좌표가 4개.
	buf_cord[3]		int // 새로 구한 좌표의 값을 넣어줄 버퍼 -- 포인터(new)로 구현?
}
*/

func get_node_N(msg *t_data) {
	// 현재 있는 노드 중 4개를 뽑아 노드 N의 delay 를 측정 <- 이건 서버측에서 msg 보내주기 전에 미리 뽑아서 초기화시켜줘야함.(매번)
	for i:=0; i < 4; i++ {
		// 저장변수[i] = ping 때리고 rtt 가져옴. (go-ping 라이브러리 사용하기)
	}
}

func get_node3(msg *t_data) {
	// 임의의 노드 3이 노드 0, 노드 1, 그리고 노드 2와의 delay 를 측정
	for i:=0; i < 3; i++ {
		// 저장변수[i] = ping 때리고 rtt 가져옴. (go-ping 라이브러리 사용하기)
   }
	// ~~~ 공식 계산 ~~~
	// x3, y3, z3에 좌표값 넣고 서버에 해당 구조체 리턴
}

func get_node2(msg *t_data) {
	for i:=0; i < 2; i++ {
		 // 저장변수[i] = ping 때리고 rtt 가져옴. (go-ping 라이브러리 사용하기)
	}
	// ~~~ 공식 계산 ~~~
	// x2, y2에 좌표값 넣고 서버에 해당 구조체 리턴 (z2 자리에는 0 넣기)
}

func get_node1(msg *t_data)  {
	// x1 := ping 때리고 rtt 가져옴. (go-ping 라이브러리 사용하기)
	// 지금 rtt를 x1에 넣고 서버에 해당 구조체 리턴 (y1, z1 자리에는 0 넣기)
}

// 보통노드(client)가 패킷을 막 받았을 경우의 코드
func check_packet(msg *t_data)
{
	
	if t_data.flag == 1 { // 지금 참여하는 노드는 1번째 노드입니다. -> (x1, 0, 0) 좌표 가져야함
		get_node1(msg)
	} else if t_data.flag == 2 { // 지금 참여하는 노드는 2번째 노드입니다. -> (x2, y2, 0) 좌표 가져야함
		get_node2(msg)
	} else if t_data.flag == 3 { // 지금 참여하는 노드는 3번째 노드입니다. (검증노드(노드0), 노드1, 노드2) 다 해서 노드가 총 3개가 있다.
		get_node3(msg)
	} else { // 노드 랜덤으로 4개 뽑아서 ping 치고 계산함.
		get_node_N(msg)
	}
}

// 이 모든 코드들은 고루틴으로 병렬적으로 실행하고, 노드 1, 2, 3이 만들어질 때까지는 wait을 걸어준다.
// 검증노드(서버측 에서는) 시작부터 모든 노드의 ip주소를 하드코딩으로 갖고있다.
// 전달되는 msg 구조체 외에, 전체 노드의 ip-주소 정보 쌍을 저장하는 map 데이터는 따로 저장한다. 