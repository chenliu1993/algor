#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#define MIN(x, y) (x < y ? x : y)
int isUgly(int num)
{
    if (num == 0)
    {
        return 0;
    }
    if (num == 1)
    {
        return 1;
    }
    if (num % 2 == 0 && isUgly(num / 2) ||
        num % 3 == 0 && isUgly(num / 3) ||
        num % 5 == 0 && isUgly(num / 5))
    {
        return 1;
    }
    return 0;
}
int main(int argc, char *argv[])
{
    printf("%d\n", isUgly(2));
    return 0;
}