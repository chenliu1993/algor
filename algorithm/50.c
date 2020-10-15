#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
double myPow(double x, int n)
{
    if (n == 0)
    {
        return 1;
    }
    double rc;
    long m = n;
    if (m < 0)
    {
        x = 1 / x;
        m = -1 * m;
    }
    if (m % 2)
    {
        return x * myPow(x, m - 1);
    }
    else
    {
        rc = myPow(x, m / 2);
        return rc * rc;
    }
}

int main(int argc, char *argv[])
{
    printf("%f\n", myPow(1, INT_));
    return 0;
}