#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
int maxProfit(int *prices, int pricesSize)
{
    int **dp = (int **)malloc(pricesSize * sizeof(int *));
    for (int i = 0; i < pricesSize; i++)
    {
        dp[i] = (int *)malloc(2 * sizeof(int));
        memset(dp[i], 0, 2 * sizeof(int));
    }
    dp[0][0] = 0;
    dp[0][1] = -prices[0];
    for (int i = 1; i < pricesSize; i++)
    {
        dp[i][0] = MAX(dp[i - 1][0], dp[i - 1][1] + prices[i]);
        dp[i][1] = MAX(dp[i - 1][0] - prices[i], dp[i - 1][1]);
    }
    return MAX(dp[pricesSize - 1][0], dp[pricesSize - 1][1]);
}

int main(int argc, char *argv[])
{
    int pricesSize = 6;
    int prices[6] = {7, 1, 5, 3, 6, 4};
    printf("%d\n", maxProfit(prices, pricesSize));
    return 0;
}