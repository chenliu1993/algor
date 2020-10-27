#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x < y ? y : x)
int PredictTheWinner(int *nums, int numsSize)
{
    int **dp = (int **)malloc(numsSize * sizeof(int *));

    for (int i = 0; i < numsSize; i++)
    {
        dp[i] = (int *)malloc(numsSize * sizeof(int));
        memset(dp[i], 0, numsSize * sizeof(int));
        dp[i][i] = nums[i];
    }
    for (int i = numsSize - 2; i >= 0; i--)
    {
        for (int j = i + 1; j < numsSize; j++)
        {
            dp[i][j] = MAX(nums[i] - dp[i + 1][j], nums[j] - dp[i][j - 1]);
        }
    }
    if (dp[0][numsSize - 1] >= 0)
    {
        return 1;
    }
    return 0;
}

int main(int argc, char *argv[])
{
    int a[3] = {1, 5, 2};
    printf("%d\n", PredictTheWinner(a, 3));
    return 0;
}