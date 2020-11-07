#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
char *pushDominoes(char *dominoes)
{
    int left = 0, mid = 0, right = 0;
    for (int i = 0; i < strlen(dominoes); i++)
    {
        if (dominoes[i] != '.')
        {
            left = right;
            right = i;
            if (dominoes[left] == '.' && dominoes[right] == 'L')
            {
                for (int j = left; j < right; j++)
                {
                    dominoes[j] = 'L';
                }
            }
            else if (dominoes[left] == dominoes[right])
            {
                for (int j = left + 1; j < right; j++)
                {
                    dominoes[j] = dominoes[left];
                }
            }
            else if (dominoes[left] == 'R' && dominoes[right] == 'L')
            {
                if (right - left < 2)
                {
                    continue;
                }
                for (int j = left + 1; j < right; j++)
                {
                    if (j < (right + left) / 2)
                    {
                        dominoes[j] = dominoes[left];
                    }
                    if (j > (right + left) / 2)
                    {
                        dominoes[j] = dominoes[right];
                    }
                }
                dominoes[(left + right) / 2] = dominoes[left];
                if ((right - left) % 2 == 0)
                {
                    dominoes[(left + right) / 2] = '.';
                }
            }
        }
    }
    if (dominoes[right] == 'R')
    {
        for (int j = right; j < strlen(dominoes); j++)
        {
            dominoes[j] = 'R';
        }
    }
    return dominoes;
}

int main(int argc, char *argv[])
{
    char a[16] = ".L.R...LR..L..\0";
    char *b = pushDominoes(a);
    printf("%s\n", b);
    return 0;
}