#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MAXSIZE 500

int findLucky(int *arr, int arrSize)
{
    int maxnum = INT_MIN;
    int *rc = (int *)calloc(MAXSIZE, sizeof(int));
    for (int i = 0; i < arrSize; i++)
    {
        rc[arr[i] - 1]++;
        printf("%d\n", rc[arr[i] - 1]);
    }
    for (int i = 0; i < MAXSIZE; i++)
    {
        if (rc[i] == i + 1)
        {
            maxnum = maxnum < i + 1 ? i + 1 : maxnum;
        }
    }
    if (maxnum == INT_MIN)
    {
        return -1;
    }
    return maxnum;
}
int main(int argc, char *argv[])
{
    int a[6] = {1, 2, 2, 3, 3, 3};
    printf("%d\n", findLucky(a, 6));
    return 0;
}
