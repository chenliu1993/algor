#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// int compareFunc(const void *a, const void *b)
// {
//     return (*(int *)a - *(int *)b);
// }
// int *findErrorNums(int *nums, int numsSize, int *returnSize)
// {
//     int *rc = (int *)calloc(2, sizeof(int));
//     int counter = 1, ptr = 0;
//     qsort(nums, numsSize, sizeof(int), compareFunc);
//     for (int i = 1; i < numsSize; i++)
//     {
//         if (nums[i - 1] == nums[i] && !rc[0])
//         {
//             rc[0] = nums[i];
//         }
//     }
//     for (int i = 0; i < numsSize; i++)
//     {
//         if (nums[i] == counter++)
//         {
//             ptr++;
//         }
//     }

//     *returnSize = 2;
//     return rc;
// }

// Hash
int *findErrorNums(int *nums, int numsSize, int *returnSize)
{
    int ans[numsSize + 1];
    *returnSize = 2;
    int *res = malloc(sizeof(int) * *returnSize);
    for (int i = 1; i < numsSize + 1; i++)
    {
        ans[i] = 0;
    }
    for (int i = 1; i < numsSize + 1; i++)
    {
        ans[nums[i - 1]]++;
        if (ans[nums[i - 1]] == 2)
        {
            res[0] = nums[i - 1];
        }
    }
    for (int i = 1; i < numsSize + 1; i++)
    {
        if (ans[i] == 0)
        {
            res[1] = i;
            break;
        }
    }
    return res;
}

void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int a[6] = {3, 2, 3, 4, 6, 5};
    printArray(a, 6);
    int *returnSize = (int *)calloc(1, sizeof(int));
    int *b = findErrorNums(a, 6, returnSize);
    printArray(b, *returnSize);
    return 0;
}