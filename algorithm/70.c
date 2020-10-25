#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int climbStairs(int n)
{
    int *dp = (int *)malloc((n + 1) * sizeof(int));
    memset(dp, 0, (n + 1) * sizeof(int));
    dp[0] = 1;
    dp[1] = 1;
    for (int i = 2; i < n + 1; i++)
    {
        dp[i] = dp[i - 1] + dp[i - 2];
    }
    return dp[n];
}

int main(int arrgc, char *argv[])
{
    printf("%d\n", climbStairs(2));
    return 0;
}