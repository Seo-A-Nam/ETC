# Overview
* Program to check if the pending signal arrvies in right order when unblokced.
* i will type signals in turn on server input.

* 시그널 unblock 시 보류된 시그널들이 알맞은 순서대로 도착하는 지 확인하기 위한 프로그램
* sever 파일을 실행한 후 여러 숫자(1~4)를 입력한다, 시그널 언블록 시 이 순서가 유지되며 도달하는 지 확인한다.

# Purpose of this experiment
* I wanted to make mini telecommunication program with signal in C
* But unfortunately i found people have issues. it was the problem that one client might have crash while getting multiple signals in different server at the same moment.
* So i come up with a idea of signal blocking. cuz it might prevent crash and get pending signals which have arrived on blocked state.

# Result
* it didn't get duplicate pending signals when it's unblocked
* while it turned out that it gets signals in order in real-time when it's unblocked

<img width="1142" alt="스크린샷 2022-01-03 오후 3 34 49" src="https://user-images.githubusercontent.com/65381957/147904737-d9663efb-e8a8-4272-8354-d70495571828.png">

<img width="1142" alt="스크린샷 2022-01-03 오후 3 47 48" src="https://user-images.githubusercontent.com/65381957/147905273-26da6d1a-1643-47d0-adb0-8ff624f4746c.png">

<br>

# Conclusion
* I'll need other way to make client get message from multi server which has sent message in same moment.
* Should i need to split blocked times in really small unit of time, and then switch those block task really fast? (inspired by TSS system)