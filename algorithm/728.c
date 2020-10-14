#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int isSelfDevided(int tgt)
{
    if (tgt == 0)
        return 0;
    int tmp = tgt;
    while (tgt != 0)
    {
        int y = tgt % 10;
        if (y == 0 || tmp % y != 0)
            return 0;
        tgt /= 10;
    }
    return 1;
}

int *selfDividingNumbers(int left, int right, int *returnSize)
{
    int *tmp = (int *)calloc(10000, sizeof(int));
    int ptr = 0;
    for (int i = left; i <= right; i++)
    {
        if (isSelfDevided(i))
        {
            tmp[ptr++] = i;
        }
    }
    int *rc = (int *)calloc(ptr, sizeof(int));
    for (int i = 0; i < ptr; i++)
    {
        rc[i] = tmp[i];
    }
    *returnSize = ptr;
    return rc;
}

int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    int *b = selfDividingNumbers(1, 22, returnSize);
    printArray(b, *returnSize);
    return 0;
}
