#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int numDecodings(char *s)
{
    int slen = strlen(s);
    int *dp = (int *)calloc(slen, sizeof(int));
    if (s[0] == '0')
    {
        dp[0] = 0;
    }
    else
    {
        dp[0] = 1;
    }
    if (slen == 1)
    {
        return dp[0];
    }
    for (int i = 1; i < slen; i++)
    {
        if (s[i] == '0')
        {
            if (s[i - 1] == '1' || s[i - 1] == '2')
            {
                if (i < 2)
                {
                    dp[i] = 1;
                    continue;
                }
                dp[i] = dp[i - 2];
            }
            else
            {
                return 0;
            }
        }
        else if (s[i - 1] == '1')
        {
            if (i < 2)
            {
                dp[i] = 2;
                continue;
            }
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        else if (s[i - 1] == '2' && s[i] < '7')
        {
            if (i < 2)
            {
                dp[i] = 2;
                continue;
            }
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        else
        {
            dp[i] = dp[i - 1];
        }
    }
    return dp[slen - 1];
}