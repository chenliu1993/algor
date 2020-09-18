#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/** 
 * dp[text1.len+1][text2.len+1]
 * dp[0][:]=0 && dp[:][0]=0
 * if text1[i]==text2[j]:
 *      dp[i][j]=dp[i-1][j-1]+1
 * else:
 *      dp[i][j]=max(dp[i-1][j],dp[i][j-1])
 * dp[text1.len][text2.len]
**/

int max(int a, int b)
{
    return a > b ? a : b;
}

int longestCommonSubsequence(char *text1, char *text2)
{
    if (text1 == NULL || text2 == NULL)
    {
        return 0;
    }
    int slen1 = strlen(text1);
    int slen2 = strlen(text2);
    int **dp = (int **)malloc(sizeof(int *) * (slen1 + 1));
    for (int i = 0; i <= slen1; i++)
    {
        dp[i] = (int *)malloc(sizeof(int) * (slen2 + 1));
        memset(dp[i], 0, sizeof(int) * (slen2 + 1));
    }
    for (int i = 0; i < slen2 + 1; i++)
    {
        dp[0][i] = 0;
    }
    for (int i = 0; i < slen1 + 1; i++)
    {
        dp[i][0] = 0;
    }
    for (int i = 1; i < slen1 + 1; i++)
    {
        for (int j = 1; j < slen2 + 1; j++)
        {
            if (text1[i - 1] == text2[j - 1])
            {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            }
            else
            {
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1]);
            }
        }
    }
    return dp[slen1][slen2];
}