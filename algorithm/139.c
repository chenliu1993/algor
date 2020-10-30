#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int inDict(char *s, int left, int right, char **wordDict, int wordDictSize)
{
    char *word = (char *)calloc(right - left + 2, sizeof(char));
    for (int i = left; i <= right; i++)
    {
        word[i] = s[i];
    }
    word[right + 1] = '\0';
    for (int i = 0; i < wordDictSize; i++)
    {
        if (strcmp(word, wordDict[i]) == 0)
        {
            return 1;
        }
    }
    return 0;
}
int wordBreak(char *s, char **wordDict, int wordDictSize)
{
    int slen = strlen(s);
    int *dp = (int *)calloc(slen + 1, sizeof(int));
    dp[0] = 1;
    for (int i = 1; i <= slen; i++)
    {
        for (int j = 0; j < wordDictSize; j++)
        {
            int len = strlen(wordDict[j]);
            int k = i - len;
            if (k < 0)
            {
                continue;
            }
            dp[i] = (dp[k] && !strncmp(s + k, wordDict[j], len)) || dp[i];
        }
    }
    return dp[slen];
}
