#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int findMaxForm(char **strs, int strsSize, int m, int n)
{
    int **dp = (int **)malloc((m + 1) * sizeof(int *));
    for (int i = 0; i < m + 1; i++)
    {
        dp[i] = (int *)calloc((n + 1), sizeof(int));
    }
    for (int i = 0; i < strsSize; i++)
    {
        int curZeros = 0;
        int curOnes = 0;
        char *s = strs[i];
        while (*s != '\0')
        {
            if (*s == '0')
            {
                curZeros++;
            }
            else
            {
                curOnes++;
            }
            ++s;
        }
        for (int j = m; j >= curZeros; j--)
        {
            for (int k = n; k >= curOnes; k--)
            {
                int tmp = dp[j - curZeros][k - curOnes] + 1;
                if (tmp > dp[j][k])
                {
                    dp[j][k] = tmp;
                }
            }
        }
    }
    return dp[m][n];
}
int main(int argc, char *argv[])
{
    char **strs = (char **)malloc(5 * sizeof(char *));
    char s0[3] = "10\0";
    char s1[5] = "0001\0";
    char s2[7] = "111001\0";
    char s3[2] = "1\0";
    char s4[2] = "0\0";
    strs[0] = s0;
    strs[1] = s1;
    strs[2] = s2;
    strs[3] = s3;
    strs[4] = s4;
    int strsSize = 5;
    int m = 5;
    int n = 3;
    printf("%d\n", findMaxForm(strs, strsSize, m, n));
    return 0;
}