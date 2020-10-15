#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int checkGood(int x)
{
    int res = 0;
    while (x > 0)
    {
        if (x % 10 == 3 || x % 10 == 4 || x % 10 == 7)
            return 0;
        else if (x % 10 == 2 || x % 10 == 5 || x % 10 == 6 || x % 10 == 9)
            res++;
        x /= 10;
    }
    return (res) ? 1 : 0;
}

int rotatedDigits(int N)
{
    int count = 0;
    for (int i = 1; i <= N; i++)
    {
        if (checkGood(i))
        {
            count++;
        }
    }
    return count;
}

int main(int argc, char *argv[])
{
    printf("%d\n", rotatedDigits(10));
    return 0;
}