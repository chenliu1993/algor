#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX(x, y) (x < y ? y : x)
/**
 * if (s[i]==s[j]){
 *  dp[i][j]=MAX(dp[i][j],dp[i+1][j-1]+2);
 * }
 * */
// int longestPalindrome(char *s)
// {
//     int slen = strlen(s);
//     if (slen == 0 || slen == 1)
//     {
//         return s;
//     }
//     int **dp = (int **)malloc((slen + 1) * sizeof(int *));
//     for (int i = 0; i < slen; i++)
//     {
//         dp[i] = (int *)malloc((slen + 1) * sizeof(int));
//         memset(dp[i], 0, sizeof(int) * (slen + 1));
//     }
//     char *res = (char *)malloc(sizeof(char) * (slen + 1));
//     for (int i = 0; i <= slen; i++)
//     {
//         dp[i][i] = 1;
//     }
//     for (int i = 0, j = 1; j < slen; i++, j++)
//     {
//         if (s[i] == s[j])
//         {
//             dp[i][j] = 1;
//         }
//     }

//     for (int k = 2; k <= slen; k++)
//     {
//         for (int i = 0, j = k; j <= slen; i++, j++)
//         {
//             if (dp[i + 1][j - 1] == 1 && s[i] == s[j])
//             {
//                 dp[i][j] = 1;
//             }
//         }
//     }

//     int start, end;
//     for (int i = 0; i <= slen; i++)
//     {
//         for (int j = i; j <= slen; j++)
//         {
//             if (dp[i][j] == 1 && (j - i) > (end - start))
//             {
//                 //printf("i is %d\tj is %d\n",i,j);
//                 start = i;
//                 end = j;
//             }
//         }
//     }
//     //printf("the start is %d\n,the end is %d\n",start,end);
//     if (end <= slen)
//     {
//         res[end + 1] = '\0';
//     }
//     while (start > 0)
//     {
//         res++;
//         start--;
//     }

//     return res;
// }
char *longestPalindrome(char *s)
{
    int slen = strlen(s);
    if (slen == 0 || slen == 1)
    {
        return s;
    }
    int left, right, len = 0, start = 0, count;
    for (int i = 0; s[i] != '\0'; i += count)
    {
        count = 1;
        left = i - 1;
        right = i + 1;
        while (s[right] != '\0' && s[i] == s[right])
        {
            right++;
            count++;
        }
        while (left >= 0 && s[right] != '\0' && s[left] == s[right])
        {
            left--;
            right++;
        }
        if (right - left - 1 > slen)
        {
            start = left + 1;
            len = right - left - 1;
        }
    }
    s[start + len] = '\0';
    return s + start;
}

int main(int argc, char *argv[])
{
    char a[6] = "babad\0";
    printf("%s\n", longestPalindrome(a));
    return 0;
}