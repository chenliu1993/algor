/** 
 * This is important 
 * */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

#define MAX(a, b) (a < b ? b : a)
#define MIN(a, b) (a > b ? b : a)
#define SIZE 128

int g[SIZE];

void InitHash(char *text, int textSize)
{
    memset(g, 0, sizeof(int) * SIZE);
    unsigned int i;
    for (int i = 0; i < textSize; ++i)
    {
        g[text[i]]++;
    }
}

int maxRepOpt1(char *text)
{
    if (text == NULL || strlen(text) == 0)
    {
        return 0;
    }
    int slen = strlen(text), i = 0;
    InitHash(text, slen);
    int maxlen = INT_MIN, j = i;
    while (j < slen)
    {
        char ch = text[i];
        int mark = 0;
        int index = -1;
        while (j < slen && text[j] == ch || !mark && g[ch] > 0)
        {
            if (text[j] != ch)
            {
                mark = 1;
                index = j;
            }
            g[text[j++]]--;
        }
        if (index == -1 || g[ch] > 0)
        {
            maxlen = MAX(maxlen, j - i);
        }
        else
        {
            maxlen = MAX(maxlen, j - i - 1);
        }
        while (i < j)
        {
            g[text[i++]]++;
        }
        if (index != -1)
        {
            j = index, i = j;
        }
    }
    return maxlen;
}

int main(int argc, char *argv[])
{
    char text[6] = "ababa\0";
    printf("%d\n", maxRepOpt1(text));
    return 0;
}