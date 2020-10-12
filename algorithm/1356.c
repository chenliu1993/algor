#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int transfer(int x)
{
    int p = 1, y = 0, yushu, count = 0;
    while (1)
    {
        yushu = x % 2;
        if (yushu)
        {
            count++;
        }
        x /= 2;
        y += yushu * p;
        p *= 10;
        if (x < 2)
        {
            if (x)
            {
                count++;
            }
            y += x * p;
            break;
        }
    }
    return count;
}
int compareFunc(const void *a, const void *b)
{
    if (transfer(*(int *)a) == transfer(*(int *)b))
    {
        return (*(int *)a - *(int *)b);
    }
    else
    {
        return transfer(*(int *)a) - transfer(*(int *)b);
    }
}
int *sortByBits(int *arr, int arrSize, int *returnSize)
{
    qsort(arr, arrSize, sizeof(int), compareFunc);
    return arr;
}
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    int a[11] = {1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1};
    int *b = sortByBits(a, 11, returnSize);
    printArray(b, 11);
    return 0;
}