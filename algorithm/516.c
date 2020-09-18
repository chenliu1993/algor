#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int max(int a, int b)
{
    return a > b ? a : b;
}
int longestPalindromeSubseq(char *s)
{
    int slen = strlen(s);
    if (slen == 1)
    {
        return 1;
    }
    else if (slen == 0)
    {
        return 0;
    }
    int maxlen = 0;
    int **dp = (int **)malloc(sizeof(int *) * slen);
    for (int i = 0; i < slen; i++)
    {
        dp[i] = (int *)malloc(sizeof(int) * slen);
        memset(dp[i], 0, sizeof(int) * slen);
        dp[i][i] = 1;
    }

    for (int i = slen - 1; i >= 0; i--)
    {
        for (int j = i + 1; j < slen; j++)
        {
            if (s[i] == s[j])
            {
                dp[i][j] = dp[i + 1][j - 1] + 2;
            }
            else
            {
                dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);
            }
        }
    }
    return dp[0][slen - 1];
}

int main(int argc, char *argv[])
{
    char a[6] = "bbbab";
    printf("%d\n", longestPalindromeSubseq(a));
    return 0;
}