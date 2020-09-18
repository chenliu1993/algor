#include <stdio.h>
#include <stdlib.h>

/**
 * dp[i] 0~i biggest value.
 * dp[i] depends nums[i] is stolen or not.
 * dp[i]=max(dp[i-1], dp[i-2]+nums[i]);
 **/
int max(int a, int b)
{
    return a > b ? a : b;
}

int rob(int *nums, int numsSize)
{
    if (nums == NULL || numsSize == 0)
    {
        return 0;
    }
    if (numsSize == 1)
        return nums[0];
    int *dp1 = (int *)malloc(sizeof(int) * numsSize);
    memset(dp1, 0, sizeof(int) * numsSize);
    int *dp2 = (int *)malloc(sizeof(int) * numsSize);
    memset(dp2, 0, sizeof(int) * numsSize);
    dp1[0] = 0;
    dp1[1] = nums[1];
    dp2[0] = nums[0];
    dp2[1] = nums[0];
    for (int i = 2; i < numsSize; i++)
    {
        dp1[i] = max(dp1[i - 1], dp1[i - 2] + nums[i]);
        dp2[i] = max(dp2[i - 1], dp2[i - 2] + nums[i]);
    }
    return max(dp1[numsSize - 1], dp2[numsSize - 2]);
}

int main(int argc, char *argv[])
{
    int a[3] = {2, 3, 2};
    printf("%d\n", rob(a, 3));
    return 0;
}