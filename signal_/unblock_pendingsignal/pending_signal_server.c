#include <unistd.h>
#include <signal.h>
#include <stdio.h>
#include <errno.h>


int main()
{
    int targetPID = 0, signo = 0;

    while (1)
    {
        printf("Enter PID of client and the signo you want to send :");
        scanf(" %d", &targetPID);
        scanf(" %d", &signo);
        // kill()은 성공 시 0 반환, 실패 시 -1 반환
        if (kill(targetPID, signo) == -1) 
        { // kill()로 시그널 보내는 것에 실패했을 경우 예외 처리
            switch (errno)
            {
                case EPERM :
                    printf("Not enough permission!\n"); 
                    break;

                case ESRCH : 
                    printf("Cannot find the process %d\n", targetPID);
                    break;
                case EINVAL :
                    printf("Invalid signo\n");
                    break;
                default :
                    break;
            }
        }
        else // kill()로 SIGKILL 시그널 보내는 데 성공한 경우
            printf("bang! -> %d to pid [%d]\n", signo, targetPID);
    }
} 