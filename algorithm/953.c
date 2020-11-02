#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int isAlienSorted(char **words, int wordsSize, char *order)
{
    int *map = (int *)calloc(26, sizeof(int));
    int od = 0, m, n;
    for (int i = 0; i < 26; i++)
    {
        map[order[i] - 'a'] = od++;
    }
    for (int i = 0; i < wordsSize - 1; i++)
    {
        int flen = strlen(words[i]);
        int llen = strlen(words[i + 1]);
        for (m = 0, n = 0; n < llen && m < flen; n++, m++)
        {
            if (map[words[i][m] - 'a'] == map[words[i + 1][n] - 'a'])
            {
                continue;
            }
            else if (map[words[i][m] - 'a'] > map[words[i + 1][n] - 'a'])
            {
                return 0;
            }
            else
            {
                break;
            }
        }
        if (m < flen && n >= llen)
        {
            return 0;
        }
    }
    return 1;
}
int main(int argc, char *argv[])
{
    char order[27] = "worldabcefghijkmnpqstuvxyz\0";
    char **words = (char **)malloc(3 * sizeof(char));
    char w1[5] = "word\0";
    char w2[6] = "world\0";
    char w3[4] = "row\0";
    words[0] = w1;
    words[1] = w2;
    words[2] = w3;
    printf("%d\n", isAlienSorted(words, 3, order));
    return 0;
}