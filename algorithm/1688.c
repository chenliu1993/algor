#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
void count(int n, int *sum)
{
    if (n == 0 || n == 1)
    {
        return;
    }
    if (n % 2 == 0)
    {
        (*sum) += (n / 2);
        count(n / 2, sum);
    }
    else
    {
        (*sum) += ((n - 1) / 2);
        count(1 + ((n - 1) / 2), sum);
    }
}
int numberOfMatches(int n)
{
    int *res = (int *)calloc(1, sizeof(int));
    count(n, res);
    return (*res);
}

int main(int argc, char *argv[])
{
    printf("%d\n", numberOfMatches(7));
    return 0;
}