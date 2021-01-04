#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

void printGrid(char **a, int aSize, int aColSize)
{
    for (int i = 0; i < aSize; i++)
    {
        for (int j = 0; j < aColSize; j++)
        {
            printf("%c,", a[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

int countHamming(int a)
{
    int rc = 0;
    while (a != 0)
    {
        rc += (a % 2);
        a >>= 1;
    }
    return rc;
}
char **readBinaryWatch(int num, int *returnSize)
{
    char **res = (char **)malloc(500 * sizeof(char *));
    int hour = 0;
    int minute = 0;
    (*returnSize) = 0;
    for (int i = 0; i < 12; i++)
    {
        for (int j = 0; j < 60; j++)
        {
            if (countHamming(i) + countHamming(j) == num)
            {
                res[(*returnSize)] = (char *)calloc(6, sizeof(char));
                sprintf(res[*returnSize], "%d:%02d", i, j);
                (*returnSize)++;
            }
        }
    }
    return res;
}

int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    char **res = readBinaryWatch(1, returnSize);
    printGrid(res, *returnSize, 4);
    return 0;
}