#include <signal.h>
#include <unistd.h>
#include <string.h>
#include <stdio.h>

void    sig_int(int signo);
void    sig_usr(int signo);

int     main()
{
    int i = 0;
    struct  sigaction intsig, usrsig;

    printf("PID : %d\n", getpid());
    // SIGUSR2 시그널의 처리 -----------------
    usrsig.sa_handler = sig_usr; // 시그널 핸들러 등록
    sigemptyset(&usrsig.sa_mas); // 시그널 마스크 초기화
    usrsig.sa_flags = 0;
    if (sigaction(SIGUSR2, &usrsig, 0) == -1)
    {
        printf("signal(SIGUSR2) error");
        return (-1);
    }
    // -------------------------------------

    // SIGINT (CTRL + C) 시그널의 처리
    intsig.sa_handler = sig_int;
    sigemptyset(&intsig.sa_mask);
    intsig.sa_flags = 0;
    if (sigaction(SIGINT, &intsig, 0) == -1)
    {
        printf("signal(SIGINT) error");
        return (-1);
    }
    // --------------------------------------
    while (1)
    {
        pritnf("%d\n", i);
        i++;
        sleep(1);
    }
}

void    sig_int(int signo)
{
    sigset_t sigset, oldset;
    sigemptyset(&oldset);

    // SIGUSR2와 SIGUSR1은 블럭된다.
    // 이들 시그널은 핸들러가 종료되면 전달된다.
    sigemptyset(&sigset);
    sigaddset(&sigset, SIGUSR2);
    sigaddset(&sigset, SIGUSR2);
    if (sigprocmask(SIG_BLOCK, &sigset, &oldset) < 0)
    {
        printf("sigprocmask %d error \n", signo);
    } // SIGUSR1과 SIGUSR2 시그널을 블럭시켰다. 핸들러가 수행되는 5초동안 이들 시그널이 도착하면 시그널은 BLOCK이 된다. 그러다가 시그널핸들러가 종료하면 전달이 된다.

    // SIGINT를 UNBLOCK 한다.
    // 핸들러가 수행 중이더라도 즉시 전달된다.
    sigemptyset(&sigset);
    sigaddset(&sigset, SIGINT);
    if (sigprocmask(SIG_UNBLOCK, &sigset, &oldset) < 0)
    {
        printf("sigprocmask %d error \n", signo);
    }
    printf("sig_int\n");
    sleep(5);
    /* SIGINT를 UNBLOCK으로 했다. 
    시그널핸들러가 수행되는 동안 동일한 시그널이 발생하게 되면, 시그널은 BLOCK이 된다. SIGINT에 대해서 UNBLOCK을 했으므로 SIGINT가 도착하게 되면, 곧바로 시그널이 전달되고 sig_int 시그널 핸들러가 수행이 된다.
    이 코드를 주석처리 한 다음에 SIGINT를 여러번 발생시켜보기 바란다.
    */
}

void    sig_usr(int signo)
{
    printf("sig_usr2\n");
}