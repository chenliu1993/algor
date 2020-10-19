#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
int dominantIndex(int *nums, int numsSize)
{
    int maxnum = INT_MIN, idx;
    for (int i = 0; i < numsSize; i++)
    {
        if (maxnum < nums[i])
        {
            maxnum = nums[i];
            idx = i;
        }
    }
    for (int i = 0; i < numsSize; i++)
    {
        if (maxnum < 2 * nums[i] && idx != i)
        {
            printf("%d,%d\n", maxnum, nums[i]);
            return -1;
        }
    }
    return idx;
}

int main(int argc, char *argv)
{
    int a[4] = {3, 6, 1, 0};
    printf("%d\n", dominantIndex(a, 4));
    return 0;
}