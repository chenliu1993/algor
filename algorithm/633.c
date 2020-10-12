#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
int judgeSquareSum(int c)
{
    int i = 0, j = (int)sqrt(c);
    while (i <= j)
    {
        if ((long long)i * i + (long long)j * j == c)
            return 1;
        while ((long long)i * i + (long long)j * j > c)
            j--;
        while ((long long)i * i + (long long)j * j < c)
            i++;
    }
    return 0;
}

int main(int argc, char *argv[])
{
    printf("%d\n", judgeSquareSum(5));
    return 0;
}