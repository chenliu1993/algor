#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
#define MAX(x, y) (x < y ? y : x)
int maxProfit(int *prices, int pricesSize)
{
    int *dp = (int *)calloc(pricesSize, sizeof(int));
    int minval;
    for (int i = 1; i < pricesSize; i++)
    {
        minval = INT_MAX;
        for (int j = 0; j < i; j++)
        {
            minval = MIN(minval, prices[j]);
        }
        dp[i] = MAX(dp[i - 1], prices[i] - minval);
    }
    return dp[pricesSize - 1];
}

int main(int argc, char *argv[])
{
    int a[6] = {7, 1, 5, 3, 6, 4};
    printf("%d\n", maxProfit(a, 6));
    return 0;
}