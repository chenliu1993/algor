#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
int check(char a, char b)
{
    if (a == 'z' && b == 'a')
    {
        return 1;
    }
    return b - a == 1;
}
int findSubstringInWraproundString(char *p)
{
    int plen = strlen(p);
    if (p == NULL || plen == 0)
    {
        return 0;
    }
    int *dp = (int *)calloc(26, sizeof(int));
    int count = 1, sum = 0, idx;
    dp[p[0] - 'a'] = 1;
    for (int i = 1; i < plen; i++)
    {
        idx = p[i] - 'a';
        if (check(p[i - 1], p[i]))
        {
            count++;
        }
        else
        {
            count = 1;
        }
        dp[idx] = MAX(dp[idx], count);
    }
    for (int i = 0; i < 26; i++)
    {
        sum += dp[i];
    }
    free(dp);
    return sum;
}
int main(int argc, char *argv[])
{
    char p[4] = "zab\0";
    printf("%d\n", findSubstringInWraproundString(p));
    return 0;
}