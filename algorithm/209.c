#include <stdio.h>
#include <limits.h>

// brute force
// int minSubArrayLen(int s, int *nums, int numsSize)
// {
//     int minlen = INT_MAX, sum = 0;
//     for (int i = 0; i < numsSize; i++)
//     {
//         sum = 0;
//         for (int j = i; j < numsSize; j++)
//         {
//             sum += nums[j];
//             if (sum >= s)
//             {
//                 printf("here %d, %d, %d\n", sum, j, i);
//                 minlen = minlen > (j - i + 1) ? (j - i + 1) : minlen;
//                 break;
//             }
//         }
//     }
//     if (minlen == INT_MAX)
//     {
//         minlen = 0;
//     }
//     return minlen;
// }

// self realization, not tested, sliding window.
int minSubArrayLen(int s, int *nums, int numsSize)
{
    if (numsSize == 0)
    {
        return 0;
    }
    int front, end, minlen, sum;
    front = end = 0;
    minlen = INT_MAX;
    for (; end < numsSize; end++)
    {
        for (; end < numsSize; end++)
        {
            sum += nums[end];
            if (sum >= s)
            {
                minlen = minlen > (end - front + 1) ? (end - front + 1) : minlen;

                end++;
                break;
            }
        }
        if (end == numsSize)
        {
            break;
        }
        sum = sum + nums[end];
        for (; front <= end; front++)
        {
            sum -= nums[front];
            if (sum < s)
            {
                minlen = minlen > (end - front + 1) ? (end - front + 1) : minlen;
                front++;
                break;
            }
        }
        if (front > end)
        {
            front = end;
        }
    }
    return minlen == INT_MAX ? 0 : minlen;
}

// sliding window
// int minSubArrayLen(int s, int *nums, int numsSize)
// {
//     if (nums == NULL)
//     {
//         return 0;
//     }
//     int front, end, minlen, sum;
//     front = end = sum = 0;
//     minlen = numsSize;
//     while (end < numsSize)
//     {
//         sum += nums[end];
//         while (sum >= s && end >= front)
//         {
//             minlen = minlen > (end - front + 1) ? (end - front + 1) : minlen;
//             sum -= nums[front++];
//         }
//         end++;
//         if (end - front == numsSize && sum < s)
//         {
//             minlen = 0;
//         }
//     }
//     return minlen;
// }

int main(int argc, char *argv[])
{
    int a[6] = {2, 3, 1, 2, 4, 3};
    int s = 7;
    printf("%d\n", minSubArrayLen(s, a, 6));
    return 0;
}