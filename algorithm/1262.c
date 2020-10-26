#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x < y ? y : x)

int maxSumDivThree(int *nums, int numsSize)
{
    int **dp = (int **)malloc((1 + numsSize) * sizeof(int *));
    for (int i = 0; i < numsSize + 1; i++)
    {
        dp[i] = (int *)calloc(3, sizeof(int));
    }
    dp[0][1] = INT_MIN;
    dp[0][2] = INT_MIN;
    for (int i = 1; i < numsSize + 1; i++)
    {
        if (nums[i - 1] % 3 == 0)
        {
            dp[i][0] = MAX(dp[i - 1][0], dp[i - 1][0] + nums[i - 1]);
            dp[i][1] = MAX(dp[i - 1][1], dp[i - 1][1] + nums[i - 1]);
            dp[i][2] = MAX(dp[i - 1][2], dp[i - 1][2] + nums[i - 1]);
        }
        else if (nums[i - 1] % 3 == 2)
        {
            dp[i][0] = MAX(dp[i - 1][0], dp[i - 1][1] + nums[i - 1]);
            dp[i][1] = MAX(dp[i - 1][1], dp[i - 1][2] + nums[i - 1]);
            dp[i][2] = MAX(dp[i - 1][2], dp[i - 1][0] + nums[i - 1]);
        }
        else if (nums[i - 1] % 3 == 1)
        {
            dp[i][0] = MAX(dp[i - 1][0], dp[i - 1][2] + nums[i - 1]);
            dp[i][1] = MAX(dp[i - 1][1], dp[i - 1][0] + nums[i - 1]);
            dp[i][2] = MAX(dp[i - 1][2], dp[i - 1][1] + nums[i - 1]);
        }
    }
    return dp[numsSize][0];
}
int main(int argc, char *argv[])
{
    int a[5] = {3, 6, 5, 1, 8};
    printf("%d\n", maxSumDivThree(a, 5));
    return 0;
}