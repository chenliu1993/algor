#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x > y ? x : y)

// write from the comments
// https://leetcode-cn.com/problems/last-stone-weight-ii/solution/cyu-yan-0-1bei-bao-wen-ti-er-wei-dp-by-zen-jie/
int lastStoneWeightII(int *stones, int stonesSize)
{
    if (stonesSize == 1)
        return stones[0];
    int sum = 0, target;
    for (int i = 0; i < stonesSize; i++)
    {
        sum += stones[i];
    }
    target = sum / 2;
    int **dp = (int **)malloc(stonesSize * sizeof(int *));
    for (int i = 0; i < stonesSize; i++)
    {
        dp[i] = (int *)malloc((target + 1) * sizeof(int));
        memset(dp[i], 0, (target + 1) * sizeof(int));
    }

    for (int i = 0; i < target + 1; i++)
    {
        if (i >= stones[0])
        {
            dp[0][i] = stones[0];
        }
    }

    for (int i = 1; i < stonesSize; i++)
    {
        for (int j = 0; j < target + 1; j++)
        {
            dp[i][j] = dp[i - 1][j];
            if (j >= stones[i])
            {
                dp[i][j] = MAX(dp[i - 1][j], dp[i - 1][j - stones[i]] + stones[i]);
            }
        }
    }
    return sum - 2 * dp[stonesSize - 1][target];
}
int main(int argc, char *argv[])
{
    int a[6] = {2, 7, 4, 1, 8, 1};
    printf("%d\n", lastStoneWeightII(a, 6));
    return 0;
}