#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MIN(x, y) (x < y ? x : y)
int minFallingPathSum(int **A, int ASize, int *AColSize)
{
    int m = ASize;
    int n = AColSize[0];
    int **dp = (int **)malloc(m * sizeof(int *));
    for (int i = 0; i < m; i++)
    {
        dp[i] = (int *)malloc(n * sizeof(int));
        memset(dp[i], 0, sizeof(int));
    }
    for (int i = 0; i < n; i++)
    {
        dp[0][i] = A[0][i];
    }
    for (int i = 1; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (j < 1)
            {
                if (j + 1 < n)
                {
                    dp[i][j] = MIN(dp[i - 1][j], dp[i - 1][j + 1]) + A[i][j];
                }
                else
                {
                    dp[i][j] = dp[i - 1][j] + A[i][j];
                }
            }
            else if (j >= n - 1)
            {
                if (j - 1 >= 0)
                {
                    dp[i][j] = MIN(dp[i - 1][j - 1], dp[i - 1][j]) + A[i][j];
                }
                else
                {
                    dp[i][j] = dp[i - 1][j] + A[i][j];
                }
            }
            else
            {
                dp[i][j] = MIN(MIN(dp[i - 1][j], dp[i - 1][j + 1]), dp[i - 1][j - 1]) + A[i][j];
            }
        }
    }
    int maxsum = INT_MAX;
    for (int i = 0; i < n; i++)
    {
        maxsum = maxsum > dp[m - 1][i] ? dp[m - 1][i] : maxsum;
    }
    return maxsum;
}

int main(int argc, char *argv[])
{
    int asize = 3;
    int **a = (int **)malloc(asize * sizeof(int *));
    int a1[3] = {1, 2, 3};
    int a2[3] = {4, 5, 6};
    int a3[3] = {7, 8, 9};
    a[0] = a1;
    a[1] = a2;
    a[2] = a3;
    int *acolsize = (int *)malloc(asize * sizeof(int));
    for (int i = 0; i < asize; i++)
    {
        acolsize[i] = asize;
    }
    printf("%d\n", minFallingPathSum(a, asize, acolsize));
    return 0;
}