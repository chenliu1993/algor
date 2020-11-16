#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int findNumberOfLIS(int *nums, int numsSize)
{
    int *dp = (int *)calloc(numsSize, sizeof(int));
    dp[0] = 1;
    int *counts = (int *)calloc(numsSize, sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        counts[i] = 1;
    }
    int maxlen = INT_MIN, count = 0;
    for (int i = 1; i < numsSize; i++)
    {
        for (int j = 0; j < i; j++)
        {
            if (nums[i] > nums[j])
            {
                if (dp[j] >= dp[i])
                {
                    dp[i] = dp[j] + 1;
                    counts[i] = counts[j];
                }
                else if (dp[j] + 1 == dp[i])
                {
                    counts[i] += counts[j];
                }
            }
        }
        maxlen = maxlen < dp[i] ? dp[i] : maxlen;
    }
    for (int i = 0; i < numsSize; i++)
    {
        if (maxlen == dp[i])
        {
            count += counts[i];
        }
    }
    printArray(dp, numsSize);
    free(dp);
    return count;
}
int main(int argc, char *argv[])
{
    int a[5] = {1, 3, 5, 4, 7};
    printf("%d\n", findNumberOfLIS(a, 5));
    return 0;
}