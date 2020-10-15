#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX(x, y) (x > y ? x : y)
// int stoneGame(int *piles, int pilesSize)
// {
//     int *rc = (int *)calloc(2, sizeof(int));
//     int left, right;
//     left = 0;
//     right = pilesSize - 1;
//     while (left < right)
//     {
//         if (piles[left] > piles[right])
//         {
//             rc[0] = piles[left++];
//         }
//         else
//         {
//             rc[0] = piles[right--];
//         }
//         if (piles[left] > piles[right])
//         {
//             rc[1] = piles[left++];
//         }
//         else
//         {
//             rc[1] = piles[right--];
//         }
//     }
//     if (left == right)
//     {
//         rc[1] += piles[left];
//     }
//     return rc[0] > rc[1] ? 1 : 0;
// }

int stoneGame(int *piles, int pilesSize)
{
    int dp[pilesSize][pilesSize];
    for (int i = 0; i < pilesSize; i++)
    {
        dp[i][i] = piles[i];
    }
    for (int i = pilesSize - 2; i >= 0; i--)
    {
        for (int j = i + 1; j < pilesSize; j++)
        {
            dp[i][j] = MAX(piles[i] - dp[i + 1][j], piles[j] - dp[i][j - 1]);
        }
    }
    return dp[0][pilesSize - 1];
}

int main(int argc, char *argv[])
{
    int a[4] = {5, 3, 4, 5};
    printf("%d\n", stoneGame(a, 4));
    return 0;
}