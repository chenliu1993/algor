#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int divisorGame(int N)
{
    int *dp = (int *)calloc(1001, sizeof(int));
    // when n==1, alice will fall
    dp[1] = 0;
    // when n==2, alice will win
    dp[2] = 1;
    for (int i = 3; i < N + 1; i++)
    {

        dp[i] = 0;
        for (int j = 1; j < i; j++)
        {
            if (i % j == 0 && dp[i - j] == 0)
            {
                dp[i] = 1;
                break;
            }
        }
    }
    return dp[N];
}
int main(int argc, char *argv[])
{
    printf("%d\n", divisorGame(3));
    return 0;
}