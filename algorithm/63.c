#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
/**
 * dp[i] is the subarray that ends with nums[i] while can be divided by 3 in maximum.
 * dp[i] (dp[i-1]+nums[i]), dp[i-1]
 **/

int uniquePathsWithObstacles(int **obstacleGrid, int obstacleGridSize, int *obstacleGridColSize)
{
    int n = obstacleGridSize;
    int m = obstacleGridColSize[0];
    int latter = 1;
    int **dp = (int **)malloc(sizeof(int *) * n);
    for (int i = 0; i < n; i++)
    {
        dp[i] = (int *)malloc(sizeof(int) * m);
        memset(dp[i], 0, sizeof(int) * m);
    }
    for (int i = 0; i < n; i++)
    {
        if (obstacleGrid[i][0])
        {
            latter = 0;
        }
        dp[i][0] = latter;
    }
    latter = 1;
    for (int i = 0; i < m; i++)
    {
        if (obstacleGrid[0][i])
        {
            latter = 0;
        }
        dp[0][i] = latter;
    }
    for (int i = 1; i < n; i++)
    {
        for (int j = 1; j < m; j++)
        {
            if (!obstacleGrid[i][j])
            {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
            else
            {
                dp[i][j] = 0;
            }
        }
    }
    return dp[n - 1][m - 1];
}
int main(int argc, char *argv[])
{
    int *a[3];
    for (int i = 0; i < 3; i++)
    {
        memset(a[i], 0, sizeof(int) * 3);
    }
    int *a1 = {0, 0, 0};
    int *a2 = {0, 1, 0};
    int *a3 = {0, 0, 0};
    a[0] = a1;
    a[1] = a2;
    a[2] = a3;
    int acolsize[3] = {3, 3, 3};
    printf("%d\n", uniquePathsWithObstacles(a, 3, acolsize));
    return 0;
}