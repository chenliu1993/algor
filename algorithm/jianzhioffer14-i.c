#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
// same as 343
#define MAX(x, y) (x < y ? y : x)
int cuttingRope(int n)
{
    int *dp = (int *)calloc(n + 1, sizeof(int));
    dp[0] = 0;
    dp[1] = 0;
    for (int i = 2; i < n + 1; i++)
    {
        int cur = INT_MIN;
        for (int j = i - 1; j > 0; j--)
        {
            cur = MAX(cur, MAX(j * dp[i - j], j * (i - j)));
        }
        dp[i] = cur;
    }
    return dp[n];
}

int main(int argc, char *argv[])
{
    printf("%d\n", cuttingRope(8));
    return 0;
}