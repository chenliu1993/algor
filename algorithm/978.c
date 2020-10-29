#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MAX(x, y) (x < y ? y : x)
// int maxTurbulenceSize(int *A, int ASize)
// {
//     if (ASize == 1)
//     {
//         return 1;
//     }
//     int **dp = (int **)malloc(ASize * sizeof(int));
//     for (int i = 0; i < ASize; i++)
//     {
//         dp[i] = (int *)malloc(ASize * sizeof(int));
//         memset(dp[i], 0, ASize * sizeof(int));
//         dp[i][i] = 1;
//         if (i != ASize - 1)
//         {
//             dp[i][i + 1] = dp[i][i] + 1;
//         }
//     }
//     for (int i = ASize - 1; i >= 2; i--)
//     {
//         for (int j = i; j < ASize; j++)
//         {
//             if ((A[j - 2] < A[j - 1] && A[j - 1] > A[j]) || (A[j - 2] > A[j - 1] && A[j - 1] < A[j]))
//             {
//                 dp[i][j] = MAX(dp[i][j - 1] + 1, dp[i][j]);
//             }
//         }
//     }
//     return dp[2][ASize - 1];
// }

int maxTurbulenceSize(int *A, int ASize)
{
    if (ASize == 1)
    {
        return ASize;
    }
    int *dp = (int *)calloc(ASize, sizeof(int));
    dp[0] = 1;
    for (int i = 1; i < ASize; i++)
    {
        if (A[i - 1] == A[i])
        {
            dp[i] = 1;
        }
    }
    int maxlen = 0;
    for (int j = 1; j < ASize; j++)
    {
        if ((j > 1 && A[j - 2] < A[j - 1] && A[j - 1] > A[j]) || (j > 1 && A[j - 2] > A[j - 1] && A[j - 1] < A[j]))
        {
            dp[j] = dp[j - 1] + 1;
        }
        else if (A[j - 1] == A[j])
        {
            dp[j] = 1;
        }
        else
        {
            // This is the key
            dp[j] = 2;
        }
    }
    for (int i = 0; i < ASize; i++)
    {
        maxlen = maxlen < dp[i] ? dp[i] : maxlen;
    }
    return maxlen;
}

int main(int argc, char *argv[])
{
    int a[2] = {9, 9};
    printf("%d\n", maxTurbulenceSize(a, 2));
    return 0;
}