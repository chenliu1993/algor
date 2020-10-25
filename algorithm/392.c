#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int isSubsequence(char *s, char *t)
{
    int slen = strlen(s);
    int tlen = strlen(t);
    int **dp = (int **)malloc((tlen + 1) * sizeof(int *));
    for (int i = 0; i < tlen + 1; i++)
    {
        dp[i] = (int *)malloc(26 * sizeof(int));
        memset(dp[i], 0, 26 * sizeof(int));
    }
    for (int i = 0; i < 26; i++)
    {
        dp[tlen][i] = tlen;
    }
    for (int i = tlen - 1; i >= 0; i--)
    {
        for (int j = 0; j < 26; j++)
        {
            if (t[i] == j + 'a')
            {
                dp[i][j] = i;
            }
            else
            {
                dp[i][j] = dp[i + 1][j];
            }
        }
    }
    int add = 0;
    for (int i = 0; i < slen; i++)
    {
        if (dp[add][s[i] - 'a'] == tlen)
        {
            return 0;
        }
        add = dp[add][s[i] - 'a'] + 1;
    }
    return 1;
}

int main(int argc, char *argv[])
{
    char s[4] = "abc\0";
    char t[7] = "ahbgdc\0";
    printf("%d\n", isSubsequence(s, t));
    return 0;
}