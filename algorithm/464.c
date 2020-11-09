#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include <limits.h>
#include <string.h>

int canIWin(int maxChoosableInteger, int desiredTotal)
{
    if (maxChoosableInteger >= desiredTotal)
    {
        return 1;
    }
    if ((((maxChoosableInteger + 1) * maxChoosableInteger) / 2) < desiredTotal)
    {
        return 0;
    }
    int dplen = (1 << maxChoosableInteger) - 1;
    int *dp = (int *)malloc(dplen * sizeof(int));
    memset(dp, 0, dplen * sizeof(int));
    for (int i = 0; i < dplen; i++)
    {
        dp[i] = -1;
    }
    return dfs(maxChoosableInteger, desiredTotal, 0, dp);
}

int dfs(int maxChoosableInteger, int desiredTotal, int state, int *dp)
{
    if (dp[state] != -1)
    {
        return dp[state];
    }
    for (int i = 1; i <= maxChoosableInteger; i++)
    {
        int tmp = (1 << (i - 1));
        if ((tmp & state) == 0)
        {
            if (desiredTotal - i <= 0 || !dfs(maxChoosableInteger, desiredTotal - i, tmp | state, dp))
            {
                dp[state] = 1;
                return 1;
            }
        }
    }
    dp[state] = 0;
    return 0;
}

int main(int argc, char *argv[])
{
    int maxChoosableInteger = 10;
    int desiredTotal = 11;
    printf("%d\n", canIWin(maxChoosableInteger, desiredTotal));
    return 0;
}