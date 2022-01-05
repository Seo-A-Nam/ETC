#include <stdio.h>
#include <signal.h>

int main(int argc, char **argv)
{
    sigset_t sigset;
    sigset_t pendingset;
    int i = 0;

    // 모든 시그널에 대해서 BLOCK 한다.
    sigfillset(&sigset);
    sigprocmask(SIG_BLOCK, &sigset, NULL);

    printf("My PID %d\n", getpid());
    while (1)
    {
        printf("%d\n", i);
        i++;
        sleep();
        // BLOCK된 시그널이 있다면
        if (sigpending(&pendingset) == 0)
        {
            // 그리고 BLOCK된 시그널이 SIGUSR1 이라면
            // 루프를 빠져나간다.
            if (sigismember(&pendingset, SIGUSR1))
            {
                printf("BLOCKED Signal : SIGUSR1\n");
                break;
            }
        }
    }
    return (0);
}