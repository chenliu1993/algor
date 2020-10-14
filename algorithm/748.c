#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char *shortestCompletingWord(char *licensePlate, char **words, int wordsSize)
{
    int *map = (int *)calloc(26, sizeof(int)), *t = (int *)calloc(26, sizeof(int));
    int *num = (int *)calloc(wordsSize, sizeof(int)), j, k = 0;
    while (*licensePlate)
    {
        if (*licensePlate >= 'A' && *licensePlate <= 'Z')
            map[*licensePlate - 'A']++;
        else if (*licensePlate >= 'a' && *licensePlate <= 'z')
            map[*licensePlate - 'a']++;
        licensePlate++;
    }
    for (int i = 0; i < wordsSize; i++)
    {
        num[i] = strlen(words[i]);
        memset(t, 0, sizeof(int) * 26);
        for (int l = 0; l < num[i]; l++)
            t[words[i][l] - 'a']++;
        // to count words[i] has chars smaller than map have.
        for (j = 0; j < 26; j++)
            if (map[j] > t[j])
            {
                num[i] = 0;
                break;
            }
    }
    j = 1000;
    for (int i = 0; i < wordsSize; i++)
        if (num[i] && num[i] < j)
        {
            j = num[i];
            k = i;
        }
    free(map);
    free(t);
    free(num);
    return words[k];
}