#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <signal.h>

/* SIGINT라는 시그널을 받을 경우, SIGINT를 잡았다!라는 
출력메시지와 함께 잡은 시그널 ‘(시그널 정보)’를 출력하는 프로그램 */
/* Program that prints signal and the message, if getting a signal SIGINT */

void    handler(int signo) // ctrl + C 입력 시마다 호출 됨
{
    printf("Caught SIGINT!\n");
    psignal(signo, "Recieved signal"); // Recieved signal : Interrupt 출력
}

int     main()
{
    if (signal(SIGINT, handler) == SIG_ERR) 
    {
        fprintf(stderr, "Cannot handle SIGINT!"); // 핸들러 등록 실패 시
        exit(EXIT_FAILURE);
    }
    for (;;) pause();
    return (0);
}


