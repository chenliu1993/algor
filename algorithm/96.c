#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int numTrees(int n)
{
    int *dp = (int *)calloc(n + 1, sizeof(int));
    dp[0] = 1;
    dp[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        for (int j = 1; j <= i; j++)
        {
            dp[i] += dp[j - 1] * dp[i - j];
        }
    }
    return dp[n];
}

int main(int argc, char *argv[])
{
    printf("%d\n", numTrees(3));
    return 0;
}