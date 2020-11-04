#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int checkPossibility(int *nums, int numsSize)
{
    if (numsSize < 3)
    {
        return 1;
    }
    int cnt = 0, temp;
    for (int i = 0; i < numsSize - 1; i++)
    {
        if (nums[i + 1] < nums[i])
        {
            temp = i;
            cnt++;
        }
        if (cnt >= 2)
        {
            break;
        }
    }
    if (cnt >= 2)
    {
        return 0;
    }
    if (temp == 0 || cnt == 0 || temp == numsSize - 2)
    {
        return 1;
    }
    if (cnt == 1)
    {
        if (nums[temp + 1] >= nums[temp - 1])
        {
            return 1;
        }
        else
        {
            if (nums[temp + 2] > nums[temp])
                return 1;
            else
                return false;
        }
    }
    return -1;
}
int main(int argc, char *argv[])
{
    int a[5] = {1, 2, 3, 4, 5};
    printf("%d\n", checkPossibility(a, 5));
    return 0;
}