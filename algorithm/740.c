#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define MAX(x, y) (x > y ? x : y)

int deleteAndEarn(int *nums, int numsSize)
{

    int *dp = (int *)calloc(10001, sizeof(int));
    int *count = (int *)calloc(10001, sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        count[nums[i]]++;
    }
    dp[1] = count[1];
    for (int i = 2; i < 10001; i++)
    {
        dp[i] = MAX(dp[i - 1], dp[i - 2] + i * count[i]);
    }
    return dp[10000];
}
int main(int argc, char *argv[])
{
    int a[3] = {3, 4, 2};
    printf("%d\n", deleteAndEarn(a, 3));
    return 0;
}