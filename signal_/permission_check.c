#include <stdio.h>
#include <unistd.h>
#include <signal.h>

// Null signal을 날려서 signal 전송 권한이 있는 지 확인하는 프로그램
// program that checks if we have signal sending permission by sending a Null signal.
int     main()
{
    int check;
    int targetPID = 0;

    printf("Which process do you want to check permission? : ");
    scanf(" %d", &targetPID);
    check = kill(targetPID, 0); // targetPID로 Null signal 보낸다
    if (check) // 실패 : 권한 없음
        fprintf(stderr, "Lack of permission\n");
    else // check == 0 -> 성공 : 권한 있음
        printf("We have permission on pid : [%d]\n", targetPID);
}
//