#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x < y ? x : y)
// int longestArithSeqLength(int *A, int ASize)
// {
//     int *dp = (int *)calloc(ASize, sizeof(int));
//     for (int i = 0; i < ASize; i++)
//     {
//         dp[i] = 1;
//     }
//     int maxlen = INT_MIN;
//     for (int i = 2; i < ASize; i++)
//     {
//         if (A[i] - A[i - 1] == A[i - 1] - A[i - 2])
//         {
//             dp[i] = MAX(dp[i], dp[i - 1] + 1);
//         }
//         maxlen = MAX(maxlen, dp[i]);
//     }
//     return maxlen;
// }

int longestArithSeqLength(int *A, int ASize)
{
    int **dp = (int **)malloc(ASize * sizeof(int *));
    for (int i = 0; i < ASize; i++)
    {
        dp[i] = (int *)calloc(ASize, sizeof(int));
    }
    for (int i = 0; i < ASize; i++)
    {
        for (int j = 0; j < ASize; j++)
        {
            dp[i][j] = 2;
        }
    }
    int ans = 0;
    for (int i = 0; i < ASize; i++)
    {
        for (int j = i + 1; j < ASize; j++)
        {
            int tgt = A[i] - (A[j] - A[i]);
            for (int k = i - 1; k >= 0; k--)
            {
                if (A[k] == tgt)
                {
                    dp[i][j] = dp[k][i] + 1;
                    break;
                }
            }
            ans = MAX(ans, dp[i][j]);
        }
    }
    return ans;
}

int main(int argc, char *argv[])
{
    int a[5] = {9, 4, 7, 2, 10};
    printf("%d\n", longestArithSeqLength(a, 5));
    return 0;
}