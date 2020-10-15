#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

// This have a limitation about overflow (int ,long, ...)
// int smallestRepunitDivByK(int K)
// {
//     int count = 1;
//     unsigned long long tgt = 1;
//     while (tgt % K)
//     {
//         count++;
//         tgt = tgt * 10 + 1;
//         if (tgt > ULONG_MAX)
//         {
//             count = -1;
//             break;
//         }
//     }
//     return count;
// }

// class Solution
// {
// public:
//     int smallestRepunitDivByK(int K)
//     {
//         vector<bool> seen(K, false);
//         int n = 0;
//         int c = 0;
//         while (!seen[n])
//         {
//             seen[n] = true;
//             n = (10 * n + 1) % K;
//             ++c;
//         }
//         return (n == 0) ? c : -1;
//     }
// };
int smallestRepunitDivByK(int K)
{
    int *kc = (int *)calloc(K, sizeof(int));
    int count = 0, i = 0;
    while (!kc[i])
    {
        kc[i] = 1;
        i = (i * 10 + 1) % K;
        count++;
    }
    return (i == 0) ? count : -1;
}

int main(int argc, char *argv[])
{
    printf("%d\n", smallestRepunitDivByK(23));
    return 0;
}