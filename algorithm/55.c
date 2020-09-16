#include <stdio.h>

int canJump(int *nums, int numsSize)
{
    // the longest point we can get to.
    int max = nums[0];
    for (int i = 1; i < numsSize; i++)
    {
        if (i > max)
        {
            return 0;
        }
        if (i + nums[i] > max)
        {
            max = i + nums[i];
        }
    }
    return 1;
}

int main(int argc, char *argv[])
{
    int a[5] = {2, 3, 1, 1, 4};
    printf("%d\n", canJump(a, 5));
    return 0;
}