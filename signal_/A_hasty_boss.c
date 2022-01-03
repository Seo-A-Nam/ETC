#include <unistd.h>
#include <signal.h>
#include <stdio.h>
#include <string.h>
#include <sys/types.h>

#define COFFEE 0
#define CAKE 1
int totalOrder[2];

// 매출이 얼마인 지 5초마다 한번씩 확인하는 프로그램.
// program that prints how much is the sales in every 5 seconds.

void    salesCheck(int sig)
{
    printf("\n[Total orders of today!]");
    printf("Coffee : %d, Cake : %d\n", totalOrder[COFFEE], totalOrder[CAKE]);
    alarm(5);
}

int     main()
{
    int order = 0;
    memset(totalOrder, 0, sizeof(int) * 2);
    signal(SIGALRM, salesCheck);
    signal(SIGTERM, salesCheck);

    alarm(5);
    while (1)
    {
        printf("Enter order (0 : coffee, 1 : cake) :");
        scanf(" %d", &order);
        if (order == 0 || order == 1)
            totalOrder[order]++;
    }
    return (0);
}