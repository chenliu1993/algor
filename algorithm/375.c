#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
#define MIN(x, y) (x < y ? x : y)
// int calculate(int left, int right)
// {
//     if (left >= right)
//     {
//         return 0;
//     }
//     int minres = INT_MAX;
//     for (int i = (left + right) / 2; i <= right; i++)
//     {
//         int res = i + MAX(calculate(i + 1, right), calculate(left, i - 1));
//         minres = MIN(res, minres);
//     }
//     return minres;
// }
// int getMoneyAmount(int n)
// {
//     return calculate(1, n);
// }

int getMoneyAmount(int n)
{
    int **dp = (int **)malloc((n + 1) * sizeof(int *));
    for (int i = 0; i < n + 1; i++)
    {
        dp[i] = (int *)malloc((n + 1) * sizeof(int));
        memset(dp[i], 0, (n + 1) * sizeof(int));
    }
    for (int len = 2; len <= n; len++)
    {
        for (int i = 1; i + len - 1 <= n; i++)
        {
            int j = i + len - 1;
            dp[i][j] = INT_MAX;
            for (int k = i; k <= j; k++)
            {
                int leftval = k <= 1 ? 0 : dp[i][k - 1];
                int rightval = k + 1 > j ? 0 : dp[k + 1][j];
                dp[i][j] = MIN(dp[i][j], k + MAX(leftval, rightval));
            }
        }
    }
    return dp[1][n];
}

int main(int argc, char *argv[])
{
    printf("%d\n", getMoneyAmount(8));
    return 0;
}