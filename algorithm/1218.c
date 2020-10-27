#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x > y ? x : y)
int longestSubsequence(int *arr, int arrSize, int difference)
{
    int *dp = (int *)malloc(arrSize * sizeof(int));
    memset(dp, 0, arrSize * sizeof(int));
    dp[0] = 1;
    int maxlen = 1;
    for (int i = 1; i < arrSize; i++)
    {
        dp[i] = 1;
        for (int j = 0; j < i; j++)
        {
            if (arr[i] - arr[j] == difference)
            {
                dp[i] = MAX(dp[i], dp[j] + 1);
            }
        }
        maxlen = MAX(maxlen, dp[i]);
    }
    return maxlen;
}

int main(int argc, char *argv[])
{
    int a[9] = {1, 5, 7, 8, 5, 3, 4, 2, 1};
    int difference = -2;
    printf("%d\n", longestSubsequence(a, 9, difference));
    return 0;
}