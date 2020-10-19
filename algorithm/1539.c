#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int findKthPositive(int *arr, int arrSize, int k)
{
    int i = 0;
    int *count = (int *)calloc(2000, sizeof(int));
    for (int i = 0; i < arrSize; i++)
    {
        count[arr[i] - 1]++;
    }
    for (int i = 0; i < 2000; i++)
    {
        if (!k)
        {
            break;
        }
        if (!count[i])
        {
            k--;
        }
    }
    return i;
}
// int findKthPositive(int *arr, int arrSize, int k)
// {
//     int ans[2001] = {0}, cnt = 0, temp;
//     for (int i = 0; i < arrSize; i++)
//     {
//         ans[arr[i]] = 1;
//     }
//     for (int i = 1; i < 2001; i++)
//     {
//         if (ans[i] != 0)
//         {
//             continue;
//         }
//         if (ans[i] == 0)
//         {
//             cnt++;
//         }
//         if (cnt == k)
//         {
//             temp = i;
//             break;
//         }
//     }
//     return temp;
// }

int main(int argc, char *argv[])
{
    int a[5] = {2, 3, 4, 7, 11};
    int k = 5;

    printf("%d\n", findKthPositive(a, 5, k));
    return 0;
}
