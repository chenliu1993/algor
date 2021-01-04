#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int fib(int n)
{
    if (n == 0 || n == 1)
    {
        return n;
    }
    int *memo = (int *)calloc(n + 1, sizeof(int));
    memo[0] = 0;
    memo[1] = 1;
    for (int i = 2; i <= n; i++)
    {
        memo[i] = memo[i - 1] + memo[i - 2];
    }
    return memo[n];
}

int main(int argc, char *argv[])
{
    printf("%d\n", fib(4));
    return 0;
}