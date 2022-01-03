#include <unistd.h>
#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

// 처음 실행 5초 동안 시그널을 Block 해준 뒤, 다시 Unblock 해줌.
// Block signals(SIGINT, SIGQUIT) for the first 5 secs when it's initiated. and then Unblock those after the while.
void handler_SIGINT(int signo)
{
    printf("Recieved signal : %s\n", strsignal(signo));
}

int main()
{
    sigset_t new;

    if (signal(SIGINT, handler_SIGINT) == SIG_ERR)
    {
        perror("signal SIGINT"); exit(1);
    }
    sigemptyset(&new);
    sigaddset(&new, SIGINT);
    sigaddset(&new, SIGQUIT);
    sigprocmask(SIG_BLOCK, &new, NULL);
    // ============ Unblock after 5 secs ==============
    sleep(5);
    printf("UnBlocking signals\n");
    sigprocmask(SIG_UNBLOCK, &new, (sigset_t *)NULL);
    return (0);
}