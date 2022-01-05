#include <signal.h>
#include <unistd.h>
#include <stdio.h>

// signal 함수를 이용한 시그널 catch
void    sig_handler(int signo);

int     main(int argc, char **argv)
{
    int i = 0;
    signal(SIGINT, SIG_IGN); // SIGINT 신호 발생 시 기본행동(종료) 대신 SIG_IGN(시그널 무시)가 실행되도록 함
    while (1)
    {
        printf("%d\n", i);
        sleep(2);
        i++;
    }
}
/* 이제 Ctrl + C 를 눌러도 키가 아예 먹지 않는다! Ctrl + C 누를 시 SIG_IGN(시그널 무시)가 실행되기 때문.
* 참고 : SIG_DFL : 기본행동을 하도록 한다.
   -> 언제 사용되는가 ? : 1. 최초 시그널을 무시했는 데, 중간에 시그널을 기본행동으로 해야할 필요가 있을 때. 2. 자식 프로세스를 생성했을 때.
   fork를 통해 자식프로세스를 생성하면, 자식 프로세스는 부모의 시그널정책까지를 그대로 복사해서 사용하게 된다. 즉 부모의 특정 시그널에 정책이 SIG_IGN 이였다면, 자식도 그대로 그 정책을 따른다.
   때로, 자식의 시그널 정책을 달리할 필요가 있을 것이다. 이 경우 사용할 수 있다.
*/

    