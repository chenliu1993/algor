#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX(x, y) (x > y ? x : y)
//dp[i][j]代表第i堆石子，M为j时先手取石子所能获得的最大石子数量
// int stoneGameII(int *piles, int n)
// {
//     int i, x, m, dp[110][110] = {0}, sum = 0;

//     for (i = n - 1; i >= 0; i--)
//     {
//         sum += piles[i];
//         for (m = 1; m <= 50; m++)
//         {
//             if (i + 2 * m >= n)
//             {
//                 dp[i][m] = sum;
//                 continue;
//             }
//             for (x = 1; i + x <= n && x <= 2 * m; x++)
//             {
//                 dp[i][m] = max(dp[i][m], sum - dp[i + x][max(x, m)]);
//             }
//         }
//     }

//     return dp[0][1];
// }

int stoneGameII(int *piles, int pilesSize)
{
    int sum = 0;
    int **dp = (int **)malloc(sizeof(int) * pilesSize);
    for (int i = 0; i < pilesSize; i++)
    {
        dp[i] = (int *)malloc(sizeof(int) * (pilesSize / 2));
        memset(dp[i], 0, sizeof(int) * (pilesSize / 2));
    }
    // return dp[0][1]
    // dp[i][j]+dp[i+x][max(x,j)]=sum
    for (int i = pilesSize - 1; i >= 0; i--)
    {
        sum += piles[i];
        for (int j = 1; j < pilesSize / 2; j++)
        {
            if (i + 2 * j >= pilesSize)
            {
                dp[i][j] = sum;
                continue;
            }
            for (int k = 1; k + i <= pilesSize && k <= 2 * j; k++)
            {
                dp[i][j] = MAX(dp[i][j], sum - dp[i + k][MAX(k, j)]);
            }
        }
    }
    return dp[0][1];
}

int main(int argc, char *argv[])
{
    int a[5] = {2, 7, 9, 4, 4};
    printf("%d\n", stoneGameII(a, 5));
    return 0;
}
