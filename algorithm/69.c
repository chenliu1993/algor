#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int mySqrt(int x)
{
    if (x == 0 || x == 1)
    {
        return x;
    }
    unsigned int l = 1, r = x;
    while (l < r)
    {
        int mid = (l + r) / 2;
        if (x / mid == mid)
        {
            return mid;
        }
        else if (x / mid > mid)
        {
            l = mid + 1;
        }
        else
        {
            r = mid;
        }
    }
    return l - 1;
}
int main(int argc, char *argv[])
{
    printf("%d\n", mySqrt(5));
    return 0;
}