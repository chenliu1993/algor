#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MIN(x, y) (x < y ? x : y)
int minimumTotal(int **triangle, int triangleSize, int *triangleColSize)
{
    int **dp = (int **)malloc(triangleSize * sizeof(int *));
    for (int i = 0; i < triangleSize; i++)
    {
        dp[i] = (int *)malloc(triangleColSize[i] * sizeof(int));
        memset(dp[i], 0, triangleColSize[i] * sizeof(int));
    }
    dp[0][0] = triangle[0][0];
    for (int i = 1; i < triangleSize; i++)
    {
        for (int j = 0; j < triangleColSize[i]; j++)
        {
            if (j - 1 >= 0 && j < triangleColSize[i - 1])
            {
                dp[i][j] = MIN(dp[i - 1][j], dp[i - 1][j - 1]) + triangle[i][j];
            }
            else if (j < 1)
            {
                dp[i][j] = dp[i - 1][j] + triangle[i][j];
            }
            else if (j >= triangleColSize[i - 1])
            {
                dp[i][j] = dp[i - 1][j - 1] + triangle[i][j];
            }
        }
    }
    int minlen = INT_MAX;
    for (int i = 0; i < triangleColSize[triangleSize - 1]; i++)
    {
        minlen = MIN(minlen, dp[triangleSize - 1][i]);
    }
    return minlen;
}
int main(int argc, char *argv[])
{
    int trangleSize = 4;
    int a1[1] = {2};
    int a2[2] = {3, 4};
    int a3[3] = {6, 5, 7};
    int a4[4] = {4, 1, 8, 3};
    int **a = (int **)malloc(trangleSize * sizeof(int *));
    a[0] = a1;
    a[1] = a2;
    a[2] = a3;
    a[3] = a4;
    int b[4] = {1, 2, 3, 4};
    printf("%d\n", minimumTotal(a, trangleSize, b));
    return 0;
}