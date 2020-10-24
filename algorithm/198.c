#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#define MAX(x, y) (x < y ? y : x)
int rob(int *nums, int numsSize)
{
    if (numsSize == 0 || nums == NULL)
    {
        return 0;
    }
    if (numsSize == 1)
    {
        return 1;
    }
    int *dp = (int *)calloc(numsSize, sizeof(int));
    dp[0] = nums[0];
    dp[1] = MAX(nums[0], nums[1]);
    for (int i = 2; i < numsSize; i++)
    {
        dp[i] = MAX(dp[i - 1], dp[i - 2] + nums[i]);
    }
    return dp[numsSize - 1];
}

int main(int argc, char *argv[])
{
    int a[4] = {1, 2, 3, 1};
    printf("%d\n", rob(a, 4));
    return 0;
}