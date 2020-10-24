#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#define MAX(x, y) (x > y ? x : y)
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int compareFunc(const void *a, const void *b)
{
    return (*(int *)a) > (*(int *)b);
}
// int *largestDivisibleSubset(int *nums, int numsSize, int *returnSize)
// {
//     if (numsSize == 0)
//     {
//         return NULL;
//     }
//     int maxnum = INT_MIN, maxcnt = INT_MIN, maxidx, rear = 0;
//     qsort(nums, numsSize, sizeof(int), compareFunc);
//     int *rc = (int *)malloc(sizeof(int) * numsSize);
//     memset(rc, 0, sizeof(int) * numsSize);
//     int *count = (int *)malloc(sizeof(int) * numsSize);
//     memset(count, 0, sizeof(int) * numsSize);
//     for (int i = numsSize - 1; i >= 1; i--)
//     {
//         for (int j = i - 1; j >= 0; j--)
//         {
//             if (!(nums[i] % nums[j]))
//             {
//                 count[i]++;
//             }
//         }
//     }
//     for (int i = 0; i < numsSize; i++)
//     {
//         if (maxcnt < count[i])
//         {
//             maxcnt = count[i];
//             maxnum = nums[i];
//             maxidx = i;
//         }
//     }
//     rc[rear++] = maxnum;
//     if (maxidx == 0)
//     {
//         goto RETURN;
//     }
//     for (int i = maxidx - 1; i >= 0; i--)
//     {
//         if (!(maxnum % nums[i]))
//         {
//             rc[rear++] = nums[i];
//             maxnum = nums[i];
//         }
//     }
// RETURN:
//     free(count);
//     rc = (int *)realloc(rc, rear * sizeof(int));
//     *returnSize = rear;
//     return rc;
// }
int *largestDivisibleSubset(int *nums, int numsSize, int *returnSize)
{
    *returnSize = 0;
    if (numsSize == 0)
        return NULL;
    qsort(nums, numsSize, sizeof(int), compareFunc);
    int dp[numsSize];
    memset(dp, 0, sizeof dp);
    for (int i = 0; i < numsSize; ++i)
    {
        *returnSize = MAX(*returnSize, dp[i] + 1);
        for (int j = i + 1; j < numsSize; ++j)
        {
            if (nums[j] % nums[i] == 0)
            {
                dp[j] = MAX(dp[j], dp[i] + 1);
            }
        }
    }
    int *res = malloc(*returnSize * sizeof(int));
    int t = 0;
    for (int i = numsSize - 1; i >= 0; --i)
    {
        if (t == 0 && dp[i] == *returnSize - 1)
        {
            res[t++] = nums[i];
            continue;
        }
        if (t > 0 && res[t - 1] % nums[i] == 0 && dp[i] == *returnSize - 1 - t)
        {
            res[t++] = nums[i];
        }
    }
    return res;
}

int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));
    int a[4] = {4, 8, 10, 240};
    int *b = largestDivisibleSubset(a, 4, returnSize);
    if (b == NULL)
    {
        printf("here\n");
        return 0;
    }
    printArray(b, *returnSize);
    return 0;
}