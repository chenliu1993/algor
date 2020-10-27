#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
// int numberOfArithmeticSlices(int *A, int ASize)
// {
//     int sum = 0;
//     int *dp = (int *)calloc(ASize, sizeof(int));
//     for (int i = 2; i < ASize; i++)
//     {
//         if (A[i] - A[i - 1] == A[i - 1] - A[i - 2])
//         {
//             dp[i] = 1 + dp[i - 1];
//             sum += dp[i];
//         }
//     }
//     return sum;
// }
int numberOfArithmeticSlices(int *A, int ASize)
{
}

int main(int argc, char *argv[])
{
    int a[4] = {1, 2, 3, 4};
    printf("%d\n", numberOfArithmeticSlices(a, 4));
    return 0;
}