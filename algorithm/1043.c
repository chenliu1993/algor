#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
int maxSumAfterPartitioning(int *arr, int arrSize, int k)
{
    if (arrSize == 0 || arr == NULL)
    {
        return 0;
    }
    int *dp = (int *)calloc(arrSize + 1, sizeof(int));
    dp[0] = 0;
    for (int i = 1; i < arrSize + 1; i++)
    {
        int maxArr = INT_MIN;
        for (int j = i - 1; j >= MAX(i - k, 0); j--)
        {
            maxArr = MAX(maxArr, arr[j]);
            dp[i] = MAX(dp[i], dp[j] + maxArr * (i - j));
        }
    }
    return dp[arrSize];
}
int main(int argc, char *argv[])
{
    int k = 3;
    int a[7] = {1, 15, 7, 9, 2, 5, 10};
    printf("%d\n", maxSumAfterPartitioning(a, 7, k));
    return 0;
}