#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char **argv)
{
    int i = 0;
    while (1)
    {
        printf("%d\n", i);
        i++;
    }
    return (0);
} // ctrl + C 를 입력했을 때 프로그램이 종료된다. (해당 프로세스에 SIGINT 전달)
