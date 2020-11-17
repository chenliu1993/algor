#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int subtractProductAndSum(int n)
{
    int rc = 0, multi = 1, sum = 0;
    while (n != 0)
    {
        rc = n % 10;
        multi *= rc;
        sum += rc;
        n = n / 10;
    }
    return multi - sum;
}
int main(int argc, char *argv[])
{
    printf("%d\n", subtractProductAndSum(4421));
    return 0;
}