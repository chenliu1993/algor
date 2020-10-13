#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define base 1337

int mypow(int a, int k)
{
    if (k == 0)
    {
        return 1;
    }
    a %= base;
    if (k % 2 == 1)
    {
        return (a * mypow(a, k - 1)) % k;
    }
    else
    {
        int sub = mypow(a, k / 2);
        return (sub * sub) % 1337;
    }
}
int superPow(int a, int *b, int bSize)
{
    if (!bSize)
    {
        return 1;
    }
    int rear = bSize - 1;
    int last = b[rear--];
    int part1 = mypow(a, last);
    int part2 = mypow(superPow(a, b, rear), 10);
    return part1 * part2;
}

int main(int argc, char *argv[])
{
    int b[2] = {1, 0};
    printf("%d\n", superPow(2, b, 2));
    return 0;
}