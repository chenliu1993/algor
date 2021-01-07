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

int compareFunc(const void *a, const void *b)
{
    return (*(int *)a) > (*(int *)b);
}
int *sequentialDigits(int low, int high, int *returnSize)
{
    int *res = malloc(0);
    (*returnSize) = 0;
    for (int i = 1; i <= 9; i++)
    {
        int num = i;
        for (int j = i + 1; j <= 9; j++)
        {
            num = num * 10 + j;
            if (num >= low && num <= high)
            {
                (*returnSize)++;
                res = realloc(res, (*returnSize) * sizeof(int));
                res[(*returnSize) - 1] = num;
            }
        }
    }
    qsort(res, (*returnSize), sizeof(int), compareFunc);
    return res;
}
int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    int low = 1000;
    int high = 13000;
    int *res = sequentialDigits(low, high, returnSize);
    printArray(res, (*returnSize));
    return 0;
}