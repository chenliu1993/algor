#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
int numRollsToTarget(int d, int f, int target)
{
    int **dp = (int **)malloc((31) * sizeof(int *));
    for (int i = 0; i < 31; i++)
    {
        dp[i] = (int *)malloc((1000 + 1) * sizeof(int));
        memset(dp[i], 0, (1000 + 1) * sizeof(int));
    }
    for (int i = 1; i < MIN(target, f) + 1; i++)
    {
        dp[1][i] = 1;
    }
    for (int i = 2; i < d + 1; i++)
    {
        for (int j = i; j < f * d + 1; j++)
        {
            for (int k = 1; j - k >= 0 && k <= f; k++)
            {
                dp[i][j] = (dp[i][j] + dp[i - 1][j - k]) % 1000000007;
            }
        }
    }
    return dp[d][target];
}
int main(int argc, char *argv[])
{
    printf("%d\n", numRollsToTarget(30, 30, 500));
    return 0;
}