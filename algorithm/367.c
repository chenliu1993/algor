#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// int isPerfectSquare(int num)
// {
//     if (num == 0)
//     {
//         return 1;
//     }
//     int i = 1;
//     while (num > 0)
//     {
//         num -= i;
//         i += 2;
//     }
//     return num == 0;
// }

// From Newton
int isPerfectSquare(int num)
{
    if (num < 2)
    {
        return 1;
    }
    long init = (num / 2);
    while (init * init > num)
    {
        init = 0.5 * (init + (num / init));
    }
    return init * init == num;
}

int main(int argc, char *argv[])
{
    printf("%d\n", isPerfectSquare(16));
    return 0;
}