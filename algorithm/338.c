#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int *countBits(int num, int *returnSize)
{
    int *count = (int *)malloc((num + 1) * sizeof(int));
    memset(count, 0, (num + 1) * sizeof(int));
    int i = 0, b = 1;
    while (b <= num)
    {
        while (i < b && i + b <= num)
        {
            count[i + b] = count[i] + 1;
            i++;
        }
        i = 0;
        b <<= 1;
    }
    *returnSize = num + 1;
    return count;
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
    int *a = countBits(5, returnSize);
    printArray(a, *returnSize);
    return 0;
}
