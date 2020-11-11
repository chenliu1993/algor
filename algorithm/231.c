#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
int isPowerOfTwo(int n)
{
    if (n == 0)
        return 0;
    long x = n;
    return (x & (x - 1)) == 0;
}
int main(int argc, char *argv[])
{
    printf("%d\n", isPowerOfTwo(16));
    return 0;
}