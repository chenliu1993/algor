#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// void swap(int *a, int *b)
// {
//     int temp = *a;
//     *a = *b;
//     *b = temp;
// }
// void q(int *nums, int left, int right)
// {
//     if (left >= right)
//     {
//         return;
//     }
//     int start = left, end = right - 1, standard = nums[right];
//     while (start < end)
//     {
//         while (start < end && nums[start] < standard)
//         {
//             start++;
//         }
//         while (start < end && nums[end] > standard)
//         {
//             end--;
//         }
//         swap(&nums[start], &nums[end]);
//     }
//     if (nums[start] >= nums[right])
//     {
//         swap(&nums[start], &nums[right]);
//     }
//     else
//     {
//         start++;
//     }
//     if (start)
//     {
//         q(nums, left, start - 1);
//     }
//     q(nums, start + 1, right);
// }
// void qicksort(int *nums, int numsSize)
// {
//     q(nums, 0, numsSize - 1);
// }

int findBigger(int *nums, int numsSize, int tgt)
{
    for (int i = 0; i < numsSize; i++)
    {
        if (nums[i] >= tgt)
        {
            return numsSize - i;
        }
    }
    return 0;
}
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
    return (*(int *)a - *(int *)b);
}
int specialArray(int *nums, int numsSize)
{
    qsort(nums, numsSize, sizeof(int), compareFunc);
    printArray(nums, numsSize);
    for (int i = 0; i < numsSize + 1; i++)
    {
        if (i == findBigger(nums, numsSize, i))
        {
            return i;
        }
    }
    return -1;
}

int main(int argc, char *argv[])
{
    int a[5] = {3, 6, 7, 7, 0};
    printf("%d\n", specialArray(a, 5));
    return 0;
}