#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int newN(int n)
{
    int rc, sum = 0;
    while (n != 0)
    {
        rc = n % 10;
        sum += rc * rc;
        n = n / 10;
    }
    return sum;
}

int isHappy(int n)
{
    int *map = (int *)calloc(1000, sizeof(int));
    n = newN(n);
    while (n != 1)
    {
        if (map[n] == 1)
        {
            return 0;
        }
        else
        {
            map[n] = 1;
        }
        n = newN(n);
    }
    return 1;
}

int main(int argc, char *argv[])
{
    printf("%d\n", isHappy(82));
    return 0;
}