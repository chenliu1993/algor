#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAX(x, y) (x > y ? x : y)
// brute force, time overlimit
// int maxProfit(int *prices, int pricesSize)
// {
//     int maxprice = INT_MIN;
//     int **dp = (int **)malloc(pricesSize * sizeof(int *));
//     for (int i = 0; i < pricesSize; i++)
//     {
//         dp[i] = (int *)malloc(pricesSize * sizeof(int));
//         memset(dp[i], 0, pricesSize * sizeof(int));
//     }
//     for (int i = 0; i < pricesSize; i++)
//     {
//         for (int j = 0; j < pricesSize; j++)
//         {
//             if (j > i)
//             {
//                 dp[i][j] = prices[j] - prices[i];
//                 maxprice = maxprice < dp[i][j] ? dp[i][j] : maxprice;
//             }
//         }
//     }
//     if (maxprice < 0)
//     {
//         return 0;
//     }
//     return maxprice;
// }

int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize == 0 || prices == NULL)
    {
        return 0;
    }
    int maxprice = INT_MIN;
    int *dp = (int *)calloc(pricesSize, sizeof(int));
    for (int i = 0; i < pricesSize; i++)
    {
        dp[i] = INT_MIN;
    }
    dp[0] = 0;
    for (int i = 1; i < pricesSize; i++)
    {
        for (int j = i - 1; j >= 0; j--)
        {
            dp[i] = MAX(prices[i] - prices[j], dp[i]);
        }
    }
    for (int i = 0; i < pricesSize; i++)
    {
        maxprice = maxprice < dp[i] ? dp[i] : maxprice;
    }
    return maxprice;
}

int main(int argc, char *argv[])
{
    int p[6] = {7, 1, 5, 3, 6, 4};
    printf("%d\n", maxProfit(p, 6));
    return 0;
}