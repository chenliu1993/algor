#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int *sortArrayByParity(int *A, int ASize, int *returnSize)
{
    int *even = (int *)calloc(ASize, sizeof(int));
    int *odd = (int *)calloc(ASize, sizeof(int));
    int *rc = (int *)calloc(ASize, sizeof(int));
    int evenCnt = 0, oddCnt = 0, ptr = 0;
    for (int i = 0; i < ASize; i++)
    {
        if (A[i] % 2 == 1)
        {
            odd[oddCnt++] = A[i];
        }
        else
        {
            even[evenCnt++] = A[i];
        }
    }
    for (int i = 0; i < evenCnt; i++)
    {
        rc[ptr++] = even[i];
    }
    for (int i = 0; i < oddCnt; i++)
    {
        rc[ptr++] = odd[i];
    }
    *returnSize = ASize;
    return rc;
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
    int a[4] = {3, 1, 2, 4};
    printArray(a, 4);
    int *b = sortArrayByParity(a, 4, returnSize);
    printArray(b, 4);
    return 0;
}