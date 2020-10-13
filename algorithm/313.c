#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define MIN(x, y) ((x < y) ? x : y)

int nthSuperUglyNumber(int n, int *primes, int primesSize)
{
    int cur = INT_MAX;
    int *dp = (int *)malloc(n * sizeof(int));
    memset(dp, INT_MAX, n * sizeof(int));
    dp[0] = 1;
    int *counter = (int *)malloc(primesSize * sizeof(int));
    memset(counter, 0, primesSize * sizeof(int));
    for (int i = 1; i < n; i++)
    {
        cur = INT_MAX;
        for (int j = 0; j < primesSize; j++)
        {
            cur = MIN(cur, primes[j] * dp[counter[j]]);
        }
        dp[i] = cur;
        for (int j = 0; j < primesSize; j++)
        {
            if (dp[i] == primes[j] * dp[counter[j]])
            {
                counter[j]++;
            }
        }
    }
    return dp[n - 1];
}

int main(int argc, char *argv[])
{
    int a[4] = {2, 7, 13, 19};
    printf("%d\n", nthSuperUglyNumber(12, a, 4));
    return 0;
}