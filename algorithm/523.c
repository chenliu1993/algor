#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// brute force will lead to time overload.
int checkSubarraySum(int *nums, int numsSize, int k)
{
    int can = 0;
    int *dp = (int *)malloc(sizeof(int) * (numsSize + 1));
    memset(dp, 0, sizeof(int) * (numsSize + 1));
    if (nums == NULL || numsSize < 2)
    {
        goto RETURN;
    }
    dp[0] = 0;
    dp[1] = dp[0] + nums[0];
    for (int i = 1; i < numsSize; i++)
    {
        dp[i + 1] = dp[i] + nums[i];
        for (int j = 0; j < i; j++)
        {
            if (k == 0)
            {
                if ((dp[i + 1] - dp[j]) == 0)
                {
                    can = 1;
                    goto RETURN;
                }
            }
            else
            {
                if ((dp[i + 1] - dp[j]) % k == 0)
                {
                    can = 1;
                    goto RETURN;
                }
            }
        }
    }
RETURN:
    free(dp);
    return can;
}

int main(int argc, char *argv[])
{
    int a[5] = {23, 2, 6, 4, 7};
    printf("%d\n", checkSubarraySum(a, 5, 6));
    return 0;
}