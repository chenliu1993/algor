#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 1310.c
int xorOperation(int n, int start)
{
    int sum = start;
    for (int i = 1; i < n; i++)
    {
        sum ^= (start + 2 * i);
    }
    return sum;
}
int main(int argc, char *argv[])
{
    printf("%d\n", xorOperation(10, 5));
    return 0;
}