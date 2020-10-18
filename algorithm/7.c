#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <limits.h>

long myPow(int tgt, int x)
{
    long sum = 1;
    for (int i = 0; i < x; i++)
    {
        sum *= tgt;
    }
    return sum;
}
int reverse(int x)
{
    if (x == INT_MIN)
    {
        return 0;
    }
    int flag = 1, res, rear = 0, tmp;
    long sum = 0;
    int *rc = (int *)calloc(100, sizeof(int));
    if (x < 0)
    {
        flag = -1;
        x = -1 * x;
    }
    while (x > 0)
    {
        res = x % 10;
        x = x / 10;
        rc[rear++] = res;
    }
    tmp = rear - 1;
    for (int i = 0; i < rear; i++)
    {
        sum += rc[i] * myPow(10, tmp--);
    }
    sum = flag < 0 ? flag * sum : sum;
    if (sum > INT_MAX || sum < INT_MIN)
    {
        return 0;
    }
    return sum;
}

int main(int argc, char *argv[])
{
    printf("%d\n", reverse(1534236469));
    return 0;
}