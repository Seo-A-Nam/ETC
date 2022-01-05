#include <sys/types.h>
#include <signal.h>
#include <stdlib.h>
#include <stdio.h>

/*
#include <sys/types.h>
#include <signal.h>

int     kill(pid_t pid, int sig);
// pid : 프로세스의 아이디
// sig : 시그널 번호
*/

// kill 함수를 이용한 시그널 보내기
int     main(int argc, char **argv)
{
    int pid;
    int sig_num;

    if (argc != 3)
    {
        printf("usage %s [pid] [signum]\n", argv[0]);
        return (1);
    }
    // 실행인자로 pid 번호와 전송할 signal 번호를 만들어서 이를 해당 pid로 보낸다.
    pid = atoi(argv[1]);
    sig_num = atoi(argv[2]);
    if (!kill(pid, sig_num))
    {
        perror("Signal Send error");
        return (1);
    }
    return (0);
}
// 프로세스에 원하는 시그널을 보내는 간단한 프로그램
// 실행시 인자로 pid, sig 차례로 넣어줌