#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int thirdMax(int *nums, int numsSize)
{
    long first = LONG_MIN, second = LONG_MIN, third = LONG_MIN;
    if (numsSize == 1)
    {
        return *nums;
    }
    else if (numsSize == 2)
    {
        return nums[0] < nums[1] ? nums[1] : nums[0];
    }
    for (int i = 0; i < numsSize; i++)
    {
        if (nums[i] > first)
        {
            third = second;
            second = first;
            first = nums[i];
        }
        else if (nums[i] < first && nums[i] > second)
        {
            third = second;
            second = nums[i];
        }
        else if (nums[i] < second && nums[i] > third)
        {
            third = nums[i];
        }
    }
    if (third == LONG_MIN)
    {
        return first;
    }
    return third;
}
int main(int argc, char *argv[])
{
    int a[3] = {3, 2, 1};
    printf("%d\n", thirdMax(a, 3));
    return 0;
}