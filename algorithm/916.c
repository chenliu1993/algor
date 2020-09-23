#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX(a, b) (((a) > (b)) ? (a) : (b))

// 发现可以将B中字母合并成单词，不过要注意单个单词中有多个某一相同字母，需要记录该字母出现的最大值，复杂度降为n*n
char **wordSubsets(char **A, int ASize, char **B, int BSize, int *returnSize)
{
    char **res = (char **)malloc(sizeof(char *) * ASize);
    *returnSize = 0;
    int numB[26] = {0}; // 存放B中出现的字母
    for (int i = 0; i < BSize; i++)
    {
        int tmp[26] = {0};
        for (int j = 0; j < strlen(B[i]); j++)
        {
            tmp[B[i][j] - 'a']++;
        }
        for (int k = 0; k < 26; k++)
        {
            numB[k] = MAX(numB[k], tmp[k]); // 确保B中每个单词都能满足，例如e，oo，出现多次，则刷新numB数组
        }
    }
    for (int i = 0; i < ASize; i++)
    {
        int flag = 1;
        int numA[26] = {0};
        for (int k = 0; k < strlen(A[i]); k++)
        {
            numA[A[i][k] - 'a']++;
        }
        for (int j = 0; j < 26; j++)
        {
            if (numB[j] != 0 && numA[j] < numB[j])
            {
                flag = 0;
                break;
            }
        }
        if (flag == 1)
        {
            res[*returnSize] = (char *)malloc(sizeof(char) * (strlen(A[i]) + 1));
            strcpy(res[*returnSize], A[i]);
            (*returnSize)++;
        }
    }
    return res;
}