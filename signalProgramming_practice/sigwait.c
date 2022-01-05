#include <signal.h>
#include <stdio.h>

/* int sigwait(const sigset_t *set, int *sig);
* // 동기적 시그널 처리 - 다른 작업들을 중단하고 시그널을 기다리겠다..
* // sigwait은 set에 등록된 시그널이 발생할 때까지 기다린다. 여기에 sigset_t라는 데이터 타입이 등장하는 데, 여기에는 시그널에 대응되는 bit 값이 설정되어있다.
* // sigwait을 다루기 위해서는 시그널을 객체로 다루는 방법에 대한 지식이 필요하다. 
* // sigaction 함수를 이용하면 시그널을 객체단위로 다룰 수 있는데, sigaction을 다루면서 sigwait의 사용방법에 대해서 설명하도록 하겠다.
*/