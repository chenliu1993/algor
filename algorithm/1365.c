#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}
void q(int *nums, int start, int end)
{
    if (start >= end)
    {
        return;
    }
    int left = start, right = end - 1, standard = nums[end];
    while (left < right)
    {
        while (left < right && nums[left] < standard)
        {
            left++;
        }
        while (left < right && nums[right] > standard)
        {
            end--;
        }
        swap(&nums[left], &nums[right]);
    }
    if (nums[left] > nums[end])
    {
        swap(&nums[left], &nums[end]);
    }
    else
    {
        left++;
    }
    if (left)
    {
        q(nums, 0, left - 1);
    }
    q(nums, left + 1, end);
}

void qicksort(int *nums, int numsSize)
{
    q(nums, 0, numsSize - 1);
}

// int *smallerNumbersThanCurrent(int *nums, int numsSize, int *returnSize)
// {
//     if (numsSize == 0 || nums == NULL)
//     {
//         return nums;
//     }
//     if (numsSize == 1)
//     {
//         int *rc = (int *)calloc(1, sizeof(int));
//         return rc;
//     }
//     int *rc = (int *)calloc(numsSize, sizeof(int));
//     int count;
//     for (int i = 0; i < numsSize; i++)
//     {
//         count = 0;
//         for (int j = 0; j < numsSize; j++)
//         {
//             if (nums[j] < nums[i] && j != i)
//             {
//                 count++;
//             }
//         }
//         rc[i] = count;
//     }
//     *returnSize = numsSize;
//     return rc;
// }

int cmp(const void *a, const void *b)
{
    return ((*(int **)a)[0] - (*(int **)b)[0]);
}

// int *smallerNumbersThanCurrent(int *nums, int numsSize, int *returnSize)
// {
//     int *data[numsSize];
//     for (int i = 0; i < numsSize; i++)
//     {
//         data[i] = malloc(sizeof(int) * 2);
//         data[i][0] = nums[i], data[i][1] = i;
//     }
//     qsort(data, numsSize, sizeof(int *), cmp);

//     int *ret = malloc(sizeof(int) * numsSize);
//     *returnSize = numsSize;
//     int prev = -1;
//     for (int i = 0; i < numsSize; i++)
//     {
//         if (prev == -1 || data[i][0] != data[i - 1][0])
//         {
//             prev = i;
//         }
//         ret[data[i][1]] = prev;
//     }
//     return ret;
// }

int *smallerNumbersThanCurrent(int *nums, int numsSize, int *returnSize)
{
    if (numsSize == 0 || nums == NULL)
    {
        return nums;
    }
    *returnSize = numsSize;
    if (numsSize == 1)
    {
        int *rc = (int *)calloc(1, sizeof(int));
        return rc;
    }
    int *rc = (int *)calloc(numsSize, sizeof(int));
    int **data = (int **)malloc(numsSize * sizeof(int *));
    for (int i = 0; i < numsSize; i++)
    {
        data[i] = (int *)malloc(2 * sizeof(int));
        memset(data[i], 0, 2 * sizeof(int));
        data[i][0] = nums[i];
        data[i][1] = i;
    }
    // qicksort(nums, numsSize);
    qsort(data, numsSize, sizeof(int *), cmp);
    int smaller = -1;
    for (int i = 0; i < numsSize; i++)
    {
        if (smaller == -1 || data[i][0] != data[i - 1][0])
        {
            smaller = i;
        }
        rc[data[i][1]] = smaller;
    }
    return rc;
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
    int *returnSize = (int *)calloc(0, sizeof(int));
    int a[5] = {8, 1, 2, 2, 3};
    int *b = smallerNumbersThanCurrent(a, 5, returnSize);
    printArray(b, *returnSize);
    return 0;
}