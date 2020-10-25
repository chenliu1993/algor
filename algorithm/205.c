#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int isIsomorphic(char *s, char *t)
{
    int slen = strlen(s);
    int *map = (int *)calloc(128, sizeof(int));
    int *ismap = (int *)calloc(128, sizeof(int));
    for (int i = 0; i < slen; i++)
    {
        if (map[s[i]] != 0 || ismap[t[i]])
        {
            if (map[s[i]] != t[i])
            {
                return 0;
            }
        }
        else
        {
            map[s[i]] = t[i];
            ismap[t[i]] = 1;
        }
    }
    return 1;
}

int main(int argc, char *argv[])
{
    char s[4] = "egg\0";
    char t[4] = "add\0";
    printf("%d\n", isIsomorphic(s, t));
    return 0;
}