#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int numWaterBottles(int numBottles, int numExchange)
{
    int res = numBottles, sum = numBottles;
    while (res >= numExchange)
    {
        res -= numExchange;
        sum++;
        res++;
    }
    return sum;
}
int main(int argc, char *argv[])
{
    int numBottles = 9;
    int numExchange = 3;
    printf("%d\n", numWaterBottles(numBottles, numExchange));
    return 0;
}