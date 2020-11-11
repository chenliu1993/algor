#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
/**
 * 如何获取二进制中最右边的 1：x & (-x)。
 * 如何将二进制中最右边的 1 设置为 0：x & (x - 1)。
**/
int isPowerOfTwo(int n)
{
    if (n == 0)
        return 0;
    long x = n;
    return (x & (-1 * x)) == x;
}
int main(int argc, char *argv[])
{
    printf("%d\n", isPowerOfTwo(16));
    return 0;
}