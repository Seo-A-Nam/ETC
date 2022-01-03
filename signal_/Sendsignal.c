#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <errno.h>

/* 어떤 프로세스를 KILL할 지 고르는 킬러 프로그램. 
pid 입력을 받아서 SIGKILL 시그널을 날린다. */
// Killer program choosing what process to kill. get targetPID as a input and then send SIGKILL signal

int main()
{
    int targetPID = 0;

    while (1)
    {
        printf("Enter PID to kill :");
        scanf(" %d", &targetPID);
        // kill()은 성공 시 0 반환, 실패 시 -1 반환
        if (kill(targetPID, SIGKILL) == -1) 
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
            printf("Bang! -> %d\n", targetPID);
    }
} 