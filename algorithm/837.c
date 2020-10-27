#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MIN(x, y) (x < y ? x : y)
// double new21Game(int N, int K, int W)
// {
//     int *dp = (int *)malloc(10001 * sizeof(int));
//     memset(dp, 0, 10001 * sizeof(int));
//     dp[0] = 0;
//     dp[1] = 1;
//     for (int i = 2; i < 10001; i++)
//     {
//         for (int j = 1; j < W + 1; j++)
//         {
//             if (i > j)
//             {
//                 dp[i] += dp[i - j];
//             }
//         }
//     }
//     printArray(dp, N);
//     return (double)dp[K] / dp[N];
// }

void printArray(double *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%f,", a[i]);
    }
    printf("\n");
}

// double new21Game(int N, int K, int W)
// {
//     if (K == 0)
//     {
//         return 1.00000;
//     }
//     double *dp = (double *)calloc(K + W, sizeof(double));
//     for (int i = K; i <= N && i < K + W; i++)
//     {
//         dp[i] = 1.00000;
//     }
//     for (int i = K - 1; i >= 0; i--)
//     {
//         for (int j = 1; j < W + 1; j++)
//         {
//             dp[i] += dp[i + j] / W;
//         }
//     }
//     printArray(dp, K + W);
//     return dp[0];
// }

double new21Game(int N, int K, int W)
{
    if (K == 0)
    {
        return 1.00000;
    }
    double *dp = (double *)calloc(K + W, sizeof(double));
    for (int i = K; i <= N && i < K + W; i++)
    {
        dp[i] = 1.00000;
    }
    dp[K - 1] = 1.0 * MIN(N - K + 1, W) / W;
    for (int i = K - 2; i >= 0; i--)
    {
        dp[i] = dp[i + 1] - (dp[i + W + 1] - dp[i + 1]) / W;
    }
    return dp[0];
}

int main(int argc, char *argv[])
{
    int n = 10;
    int k = 1;
    int w = 10;
    printf("%f\n", new21Game(n, k, w));
    return 0;
}