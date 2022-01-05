#include <signal.h>
#include <unistd.h>
#include <stdio.h>

/*
#include <signal.h>
typedef void (*sighandler_t)(int);

sighandler_t signal(int signum, sighandler_t handler); s

// signum : 제어할 시그널 번호
// sighandler_t : signum을 받았을 때. 호출할 함수
*/

// signal 함수를 이용한 시그널 catch
void    sig_handler(int signo);

int     main(int argc, char **argv)
{
    int i = 0;
    signal(SIGINT, (void *)sig_handler); // SIGINT 신호 발생 시 기본행동(종료) 대신 sig_handler 실행.
    while (1)
    {
        printf("%d\n", i);
        sleep(2);
        i++;
    }
}

void    sig_handler(int signo)
{
    printf("I Received SIGINT(%d)\n", SIGINT);
}

/* ctrl + C 입력시 SIGINT가 발생하여 기본행동으로 종료하게 되는 데, 
    이때 signal 함수를 이용해서 SIGINT에 대해서 sig_handler라는 함수를 실행하도록 했다.
    -> 이제 ctrl+c를 입력하게 되면 프로세스가 종료되는 대신 sig_hanlder를 실행한다. */
    