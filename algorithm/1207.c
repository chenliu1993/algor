#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
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
int uniqueOccurrences(int *arr, int arrSize)
{
    int *count = (int *)calloc(2002, sizeof(int));
    int *rc = (int *)calloc(arrSize, sizeof(int));
    for (int i = 0; i < arrSize; i++)
    {
        count[arr[i] + 1000]++;
    }
    for (int i = 0; i < 2002; i++)
    {
        if (count[i] > 0 && rc[count[i]] > 0)
        {
            return 0;
        }
        else
        {
            rc[count[i]] = 1;
        }
    }
    return 1;
}

int main(int argc, char *argv[])
{
    int a[10] = {-3, 0, 1, -3, 1, 1, 1, -3, 10, 0};
    printf("%d\n", uniqueOccurrences(a, 10));
    return 0;
}