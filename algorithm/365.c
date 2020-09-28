#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int gcd(int a, int b)
{
    int m = 0;
    while (b > 0)
    {
        m = a % b;
        a = b;
        b = m;
    }
    printf("gcd is %d\n", a);
    return a;
}
int canMeasureWater(int x, int y, int z)
{
    if (x == 0 && y == 0)
    {
        return z == 0;
    }
    if (z > x + y)
    {
        return 0;
    }
    int g = y == 0 ? gcd(y, x) : gcd(x, y);
    if (z % g == 0)
    {
        return 1;
    }
    return 0;
}
int main(int argc, char *argv[])
{
    printf("%d\n", canMeasureWater(2, 0, 2));
    return 0;
}