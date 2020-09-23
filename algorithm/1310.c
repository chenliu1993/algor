#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void printArray(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        printf("%d,", arr[i]);
    }
    printf("\n");
}
// brute force will make time overload
// int *xorQueries(int *arr, int arrSize, int **queries, int queriesSize, int *queriesColSize, int *returnSize)
// {
//     int *rc = (int *)malloc(sizeof(int) * queriesSize);
//     memset(rc, 0, sizeof(int) * queriesSize);
//     int rear = 0;
//     for (int i = 0; i < queriesSize; i++)
//     {
//         rc[rear] = arr[queries[i][0]];
//         for (int j = queries[i][0] + 1; j <= queries[i][1]; j++)
//         {
//             rc[rear] ^= arr[j];
//         }
//         rear++;
//     }
//     return rc;
// }

// XOR[Li, Ri] = arr[Li] xor arr[Li+1] xor ... xor arr[Ri]
// A xor A =0;
// XOR[0, Ri] = XOR[0, Li-1] xor XOR[Li, Ri]
// XOR[0, Ri] xor XOR[0, Li-1] = XOR[Li, Ri]

int *xorQueries(int *arr, int arrSize, int **queries, int queriesSize, int *queriesColSize, int *returnSize)
{
    int xor ;
    int *prefixXor = (int *)malloc(sizeof(int) * arrSize);
    memset(prefixXor, 0, sizeof(int) * arrSize);

    int *retArr = (int *)malloc(sizeof(int) * queriesSize);
    memset(retArr, 0, sizeof(int) * queriesSize);

    prefixXor[0] = arr[0];
    for (int i = 1; i < arrSize; i++)
    {
        prefixXor[i] = prefixXor[i - 1] ^ arr[i];
    }

    for (int i = 0; i < queriesSize; i++)
    {
        /* XOR[0, Ri] xor XOR[0, Li-1] = XOR[Li, Ri] */
        retArr[i] = prefixXor[queries[i][1]] ^ ((queries[i][0] != 0) ? prefixXor[queries[i][0] - 1] : 0);
    }

    free(prefixXor);

    *returnSize = queriesSize;
    return retArr;
}

int main(int argc, char *argv[])
{
    int arr[4] = {1, 3, 4, 8};
    int **q = (int **)malloc(sizeof(int *) * 4);
    for (int i = 0; i < 4; i++)
    {
        q[i] = (int *)malloc(sizeof(int) * 2);
        memset(q[i], 0, sizeof(int) * 2);
    }
    int s[4] = {2, 2, 2, 2};
    int q1[2] = {0, 1};
    int q2[2] = {1, 2};
    int q3[2] = {0, 3};
    int q4[2] = {3, 3};
    int *rc = xorQueries(arr, 4, q, 4, s, s);
    printArray(rc, 4);
    return 0;
}