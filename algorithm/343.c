#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x > y ? x : y)
int integerBreak(int n)
{
    int cur;
    int *dp = (int *)calloc(n + 1, sizeof(int));
    dp[0] = dp[1] = 0;
    for (int i = 2; i < n + 1; i++)
    {
        cur = 0;
        for (int j = 1; j < i; j++)
        {
            cur = MAX(cur, MAX(j * (i - j), j * dp[i - j]));
        }
        dp[i] = cur;
    }
    return dp[n];
}
int main(int argc, char *argv[])
{
    printf("%d\n", integerBreak(10));
    return 0;
}