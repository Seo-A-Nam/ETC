#include <signal.h>
#include <stdio.h>

/* void (*signal(int signum, void (*handler) (int))) (int));
* // signum - 제어하고자하는 시그널 번호, handler - 시그널이 발생했을 때 실행할 함수
*/

// SIGSTOP가 전달되면 mystop()이라는 시그널 등록 함수를 실행시키는 예제 프로그램
// signal을 이용한 비동기적 시그널 처리 ; 비동기적 - 특정 시점에서 시그널이 발생하는 것을 기다리지 않는다. 
void    mystop(int signo)
{
    printf("I Recieved Signal : %d\n", signo);
}

int     main(int argc, char **argv)
{
    int i = 0;
    signal(SIGQUIT, (void *)mystop);

    while (1)
    {
        printf("%d\n", i);
        i++;
        sleep();
    }
    return (1);
}