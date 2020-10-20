#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MIN(x, y) (x < y ? x : y)
int numSquares(int n)
{
    int *dp = (int *)calloc(n + 1, sizeof(int));
    for (int i = 1; i <= n; i++)
    {
        dp[i] = i;
    }
    for (int i = 1; i <= n; i++)
    {
        for (int j = 1; i - j * j >= 0; j++)
        {

            dp[i] = MIN(dp[i], dp[i - j * j] + 1);
        }
    }
    return dp[n];
}

int main(int argc, char *argv[])
{
    printf("%d\n", numSquares(13));
    return 0;
}