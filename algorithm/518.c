#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int change(int amount, int *coins, int coinsSize)
{
    int *dp = (int *)calloc((amount + 1), sizeof(int));
    dp[0] = 1;
    for (int i = 1; i < coinsSize + 1; i++)
    {
        int w = coins[i - 1];
        for (int j = w; j < amount + 1; j++)
        {
            dp[j] += dp[j - w];
        }
    }
    return dp[amount];
}

int main(int argc, char *argv[])
{
    int coinsSize = 3;
    int coins[3] = {1, 2, 5};
    int amount = 5;
    printf("%d\n", change(amount, coins, coinsSize));
    return 0;
}