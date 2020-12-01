#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

int compareFunc(const void *a, const void *b)
{
    return *((int *)a) < *((int *)b);
}
int minCount(int *coins, int coinsSize)
{
    qsort(coins, coinsSize, sizeof(int), compareFunc);
    int count = 0;
    for (int i = 0; i < coinsSize; i++)
    {
        if (coins[i] % 2 == 0)
        {
            count += coins[i] / 2;
        }
        else
        {
            count += coins[i] / 2;
            count++;
        }
    }
    return count;
}

int main(int argc, char *argv[])
{
    int a[3] = {4, 2, 1};
    printf("%d\n", minCount(a, 3));
    return 0;
}