#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#define MAX(x, y) (x < y ? y : x)
int maxSubArray(int *nums, int numsSize)
{
    int *dp = (int *)calloc(numsSize, sizeof(int));
    dp[0] = nums[0];
    int maxsum = INT_MIN;
    for (int i = 1; i < numsSize; i++)
    {
        dp[i] = MAX(dp[i - 1] + nums[i], nums[i]);
    }
    for (int i = 0; i < numsSize; i++)
    {
        maxsum = maxsum < dp[i] ? dp[i] : maxsum;
    }
    free(dp);
    return maxsum;
}
int main(int argc, char *argv[])
{
    int a[9] = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    printf("%d\n", maxSubArray(a, 9));
    return 0;
}