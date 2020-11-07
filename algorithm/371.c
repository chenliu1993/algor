#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int getSum(int a, int b)
{
    return !(a && b) ? a | b : getSum(a ^ b, (unsigned int)(a & b) << 1);
}

int main(int argc, char *argv[])
{
    printf("%d\n", getSum(5, 7));
    return 0;
}