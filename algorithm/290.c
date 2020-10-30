#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define MAXSIZE 1000
int wordPattern(char *pattern, char *s)
{
    int plen = strlen(pattern);
    char *sub[MAXSIZE];
    int rear = 0;
    sub[rear] = strtok(s, " ");
    while (sub[rear] != NULL)
    {
        rear++;
        sub[rear] = strtok(NULL, " ");
    }
    if (plen != rear)
    {
        return 0;
    }
    for (int i = 0; i < rear; i++)
    {
        for (int j = i + 1; j < rear; j++)
        {
            if ((strcmp(sub[i], sub[j]) == 0 && pattern[i] != pattern[j]) || (strcmp(sub[i], sub[j]) != 0 && pattern[i] == pattern[j]))
            {
                return 0;
            }
        }
    }
    return 1;
}
int main(int argc, char *argv[])
{
    char a[5] = "abba\0";
    char b[17] = "dog cat cat fish\0";
    printf("%d\n", wordPattern(a, b));
    return 0;
}