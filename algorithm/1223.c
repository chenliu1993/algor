#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int dieSimulator(int n, int *rollMax, int rollMaxSize)
{
    int divider = (int)pow(10, 7) + 9;
    int ***dp = (int ***)malloc(n * sizeof(int **));
    for (int i = 0; i < n; i++)
    {
        dp[i] = (int **)malloc(7 * sizeof(int *));
        for (int j = 0; j < 7; j++)
        {
            dp[i][j] = (int *)malloc(16 * sizeof(int));
            memset(dp[i][j], 0, 16 * sizeof(int));
        }
    }
    for (int i = 0; i < 7; i++)
    {
        dp[0][i][1] = 1;
    }
    for (int i = 1; i < n; i++)
    {
        // current point
        for (int j = 1; j < 7; j++)
        {
            // last point
            for (int k = 1; k < 7; k++)
            {
                if (j == k)
                {
                    for (int t = 1; t < rollMax[k - 1]; t++)
                    {
                        dp[i][j][t + 1] = dp[i - 1][k][t];
                        dp[i][j][t + 1] = dp[i][j][t + 1] % divider;
                    }
                }
                else
                {
                    for (int t = 1; t <= rollMax[k - 1]; t++)
                    {
                        dp[i][j][1] += dp[i - 1][k][t];
                        dp[i][j][1] = dp[i][j][1] % divider;
                    }
                }
            }
        }
    }
    int sum = 0;
    for (int i = 1; i < 7; i++)
    {
        for (int j = 1; j <= rollMax[i - 1]; j++)
        {
            sum = (sum + dp[n - 1][i][j]) % divider;
        }
    }
    return sum;
}

int main(int argc, char *argv[])
{
    int roll[] = {1, 1, 1, 2, 2, 3};
    printf("%d\n", dieSimulator(3, roll, 6));
    return 0;
}