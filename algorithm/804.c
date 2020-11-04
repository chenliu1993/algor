#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int uniqueMorseRepresentations(char **words, int wordsSize)
{
    char dict[26][5] = {".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."};

    char morse[100][60] = {0};
    int count = 0;
    for (int i = 0; i < wordsSize; i++)
    {
        char tmp[60] = {0};
        int flag = 0;
        for (int j = 0; j < strlen(words[i]); j++)
        {
            strcat(tmp, dict[words[i][j] - 'a']);
        }
        for (int k = 0; k < count; k++)
        {
            if (strcmp(morse[k], tmp) == 0)
            {
                flag = 1;
                break;
            }
        }
        if (flag == 0)
        {
            strcpy(morse[count], tmp);
            count++;
        }
    }
    return count;
}
