#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <signal.h>

// but a single signal handler may be used for multiple signals.
static void	write_nbr_fd(int n, int fd)
{
	char	c;

	c = (n % 10) + '0';
	if (n >= 10)
		write_nbr_fd(n / 10, fd);
	write(fd, &c, 1);
}

void		ft_putnbr_fd(int n, int fd)
{
	if (n == -2147483648)
	{
		write(fd, "-2147483648", 11);
		return ;
	}
	else if (n < 0)
	{
		write(fd, "-", 1);
		n *= -1;
	}
	write_nbr_fd(n, fd);
}

void signalCheck(int signo)
{
    write(1, "signo : ", 8);
    ft_putnbr_fd(signo, 1);
    write(1, "\n", 1);
    sleep(1);
}

int main()
{
    sigset_t new;

    printf("Client PID : %d\n", getpid());
    
    printf("\nBlocking signals\n");
    sigemptyset(&new);
    if (signal(1, signalCheck) == SIG_ERR)
    {
        perror("signal SIGINT"); exit(1);
    } // it always make error when i put this in a loop of i = 1 ~ 4
    if (signal(2, signalCheck) == SIG_ERR)
    {
        perror("signal SIGINT"); exit(1);
    }
    if (signal(3, signalCheck) == SIG_ERR)
    {
        perror("signal SIGINT"); exit(1);
    }
    if (signal(4, signalCheck) == SIG_ERR)
    {
        perror("signal SIGINT"); exit(1);
    }
    for (int i = 0; i < 4; i++)
    {
        sigaddset(&new, i);
    }
    sigprocmask(SIG_BLOCK, &new, NULL);
    // ============ Unblock after 5 secs ==============
    sleep(30);
    printf("UnBlocking signals\n");
    sigprocmask(SIG_UNBLOCK, &new, (sigset_t *)NULL);
    sleep(30);
    return (0);
}