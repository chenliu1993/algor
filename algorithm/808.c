#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MAXZERO(x) (x > 0 ? x : 0)
double soupServings(int N)
{
    N = N / 25 + (N % 25 > 0 ? 1 : 0);
    if (N >= 500)
    {
        return 1.0;
    }
    double **dp = (double **)malloc((N + 1) * sizeof(double *));
    for (int i = 0; i < N + 1; i++)
    {
        dp[i] = (double *)malloc((N + 1) * sizeof(double));
        memset(dp[i], 0, (N + 1) * sizeof(double));
    }
    for (int k = 0; k <= 2 * N; k++)
    {
        for (int i = 0; i <= N; i++)
        {
            int j = k - i;
            if (j < 0 || j > N)
            {
                continue;
            }
            double ans = 0.0;
            if (i == 0)
            {
                ans = 1.0;
            }
            if (i == 0 && j == 0)
            {
                ans = 0.5;
            }
            if (i > 0 && j > 0)
            {
                ans = 0.25 * (dp[MAXZERO(i - 4)][j] + dp[MAXZERO(i - 3)][MAXZERO(j - 1)] + dp[MAXZERO(i - 2)][MAXZERO(j - 2)] + dp[MAXZERO(i - 1)][MAXZERO(j - 3)]);
            }
            dp[i][j] = ans;
        }
    }
    return dp[N][N];
}

int main(int argc, char *argv[])
{
    printf("%f\n", soupServings(50));
    return 0;
}