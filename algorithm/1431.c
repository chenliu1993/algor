#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int maxCandies(int *a, int asize, int idx, int append)
{
    int sum = a[idx] + append;
    for (int i = 0; i < asize; i++)
    {
        if (i != idx && sum < a[i])
        {
            return 0;
        }
    }
    return 1;
}
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int *kidsWithCandies(int *candies, int candiesSize, int extraCandies, int *returnSize)
{
    int *rc = (int *)calloc(candiesSize, sizeof(int));
    for (int i = 0; i < candiesSize; i++)
    {
        if (maxCandies(candies, candiesSize, i, extraCandies))
        {
            rc[i] = 1;
        }
    }
    *returnSize = candiesSize;
    return rc;
}

int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    int a[5] = {2, 3, 5, 1, 3};
    int ex = 3;
    int *b = kidsWithCandies(a, 5, ex, returnSize);
    printArray(b, *returnSize);
    return 0;
}