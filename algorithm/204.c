#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int countPrimes(int n)
{
    int *isPrime = (int *)calloc(n, sizeof(int));
    for (int i = 2; i < n; i++)
    {
        isPrime[i] = 1;
    }
    for (int i = 2; i * i < n; i++)
    {
        if (!isPrime[i])
            continue;
        for (int j = i * i; j < n; j += i)
        {
            isPrime[j] = 0;
        }
    }
    int count = 0;
    for (int i = 2; i < n; i++)
    {
        if (isPrime[i])
            count++;
    }
    return count;
}
int main(int argc, char *argv[])
{
    printf("%d\n", countPrimes(2));
    return 0;
}