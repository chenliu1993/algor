#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
int coinChange(int *coins, int coinsSize, int amount)
{
    int *dp = (int *)calloc(amount + 1, sizeof(int));
    for (int i = 1; i < amount + 1; i++)
    {
        dp[i] = INT_MAX;
    }
    for (int i = 1; i <= amount; i++)
    {
        for (int j = 0; j < coinsSize; j++)
        {
            if (i >= coins[j])
            {
                dp[i] = MIN(dp[i], dp[i - coins[j]] + 1);
            }
        }
    }
    if (dp[amount] == INT_MAX)
    {
        return -1;
    }
    return dp[amount];
}
int main(int argc, char *argv[])
{
    int a[3] = {1, 2, 5};
    printf("%d\n", coinChange(a, 3, 11));
    return 0;
}