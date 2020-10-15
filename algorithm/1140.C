#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int stoneGameII(int *piles, int n)
{
    int i, x, m, dp[110][110] = {0}, sum = 0;

    for (i = n - 1; i >= 0; i--)
    {
        sum += piles[i];
        for (m = 1; m <= 50; m++)
        {
            if (i + 2 * m >= n)
            {
                dp[i][m] = sum;
                continue;
            }
            for (x = 1; i + x <= n && x <= 2 * m; x++)
            {
                dp[i][m] = max(dp[i][m], sum - dp[i + x][max(x, m)]);
            }
        }
    }

    return dp[0][1];
}
