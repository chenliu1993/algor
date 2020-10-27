#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int findTargetSumWays(int *nums, int numsSize, int S)
{
    int **dp = (int **)malloc(numsSize * sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        dp[i] = (int *)malloc(2001 * sizeof(int));
        memset(dp[i], 0, 2001 * sizeof(int));
    }
    dp[0][nums[0] + 1000] = 1;
    dp[0][-nums[0] + 1000] = 1;
    for (int i = 1; i < numsSize; i++)
    {
        for (int j = -1000; j <= 1000; j++)
        {
            // if (dp[i - 1][j + 1000] > 0)
            // {
            dp[i][j + 1000] = dp[i - 1][j - nums[i] + 1000] + dp[i - 1][j + nums[i] + 1000];
            // dp[i][j + nums[i] + 1000] += dp[i - 1][j + 1000];
            // dp[i][j - nums[i] + 1000] += dp[i - 1][j + 1000];
            // }
        }
    }
    return S > 1000 ? 0 : dp[numsSize - 1][S + 1000];
}
int main(int argc, char *argv[])
{
    int a[5] = {1, 1, 1, 1, 1};
    int S = 3;
    printf("%d\n", findTargetSumWays(a, 5, S));
    return 0;
}
