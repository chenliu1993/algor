#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int minSteps(int n)
{
    int buf = 0;
    int curA = 1;
    int op = 0;
    while (n > curA)
    {
        if (n % curA == 0)
        {
            buf = curA;
            curA = curA * 2;
            op += 2;
        }
        else
        {
            op++;
            curA += buf;
        }
    }
    return op;
}

int main(int argc, char *argv[])
{
    printf("min steps %d\n", minSteps(10));
    return 0;
}