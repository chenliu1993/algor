#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
// char **uncommonFromSentences(char *A, char *B, int *returnSize)
// {
//     const char delim[2] = " ";
//     char *tokenA, *tokenB;
//     char ASum[202][202] = {0};
//     char BSum[202][202] = {0};
//     tokenA = strtok(A, delim);
//     tokenB = strtok(B, delim);
//     int i = 0;
//     while (tokenA != NULL)
//     {
//         strcpy(ASum[i++], tokenA);
//         tokenA = strtok(NULL, delim);
//     }
//     i = 0;
//     while (tokenB != NULL)
//     {
//         strcpy(BSum[i++], tokenB);
//         tokenB = strtok(NULL, delim);
//     }
// }

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
typedef struct
{
    char *string;
    int freq;
} word;

void getwords(char *string, int *wordscount, word **wordslist);

char **uncommonFromSentences(char *A, char *B, int *returnSize)
{
    char **ans = malloc(sizeof(char *) * 200);
    int i, wordcount = 0;
    word **wordlist = malloc(sizeof(word *) * 200);

    getwords(A, &wordcount, wordlist);
    getwords(B, &wordcount, wordlist);

    for (i = 0, *returnSize = 0; i < wordcount; i++)
    {
        if (wordlist[i]->freq == 1)
        {
            ans[(*returnSize)++] = wordlist[i]->string;
        }
    }
    return ans;
}

void getwords(char *string, int *wordscount, word **wordlist)
{
    char *temp;
    int count, flag = 0, found = 0;

    while (*string != '\0')
    {
        if (flag == 0)
        {
            while (!isalpha(*string))
            { //过非字母characters
                string++;
            }
            if (*string == '\0')
            { //到底出while loop
                break;
            }
            else
            { //否则创建单词的空字符串
                flag = 1;
                temp = malloc(sizeof(char) * 200);
                count = 0;
            }
        }
        temp[count++] = *string;
        string++;

        if (isalpha(*string) == 0)
        { //检查下一字符是否是字母，如果不是，加'\0'终止此单词的字符串写入
            temp[count] = '\0';
            //查找对应字符串是否在库中
            flag = 0;
            for (int i = 0; i < (*wordscount); i++)
            {
                if (strcmp(wordlist[i]->string, temp) == 0)
                {
                    //如果对应字符串已在， 库中的对应字符串频率加一
                    wordlist[i]->freq++;
                    found = 1;
                    break;
                }
            }

            if (!found)
            { //否则， 添加进库
                wordlist[(*wordscount)] = malloc(sizeof(word));
                wordlist[(*wordscount)]->string = temp;
                wordlist[(*wordscount)++]->freq = 1;
            }
            found = 0;
        }
    }
}