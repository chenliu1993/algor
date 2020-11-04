#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int isAnagram(char *s, char *t)
{
    int *map = (int *)calloc(26, sizeof(int));
    int slen = strlen(s);
    int tlen = strlen(t);
    if (slen != tlen)
    {
        return 0;
    }
    for (int i = 0; i < slen; i++)
    {
        map[s[i] - 'a']++;
    }
    for (int i = 0; i < tlen; i++)
    {
        if (map[t[i] - 'a'] != 0)
        {
            map[t[i] - 'a']--;
        }
        else
        {
            return 0;
        }
    }
    for (int i = 0; i < 26; i++)
    {
        if (map[i] != 0)
        {
            return 0;
        }
    }
    return 1;
}

int main(int argc, char *argv[])
{
    char s[8] = "anagram\0";
    char t[8] = "nagaram\0";
    printf("%d\n", isAnagram(s, t));
    return 0;
}