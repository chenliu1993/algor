#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int validMountainArray(int *A, int ASize)
{
    if (ASize < 3)
    {
        return 0;
    }
    int i = 0;
    while (i < ASize - 1 && A[i] < A[i + 1])
    {
        i++;
    }
    if (i == 0 || i == ASize - 1)
    {
        return 0;
    }
    while (i < ASize - 1 && A[i] > A[i + 1])
    {
        i++;
    }
    return i == ASize - 1;
}
int main(int argc, char *argv[])
{
    int a[4] = {0, 3, 2, 1};
    printf("%d\n", validMountainArray(a, 4));
    return 0;
}
