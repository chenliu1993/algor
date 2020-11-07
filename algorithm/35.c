#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int searchInsert(int *nums, int numsSize, int target)
{
    int position = -1;
    for (int i = 0; i < numsSize; i++)
    {
        if (nums[i] == target)
        {
            return i;
        }
        else if (target < nums[i])
        {
            position = i;
            break;
        }
    }
    if (position == -1)
    {
        position = numsSize;
    }
    return position;
}

int main(int argc, char *argv[])
{
    int a[4] = {1, 3, 5, 6};
    int target = 2;
    printf("%d\n", searchInsert(a, 4, target));
    return 0;
}