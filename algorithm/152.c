#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
#define MAX(x, y) (x < y ? y : x)
#define MIN(x, y) (x > y ? y : x)
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int maxProduct(int *nums, int numsSize)
{
    if (nums == NULL || numsSize == 0)
    {
        return 0;
    }
    else if (numsSize == 1)
    {
        return nums[0];
    }
    int *dp = (int *)calloc(numsSize, sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        dp[i] = nums[i];
    }
    int maxproduct = 1, minproduct = 1, maxval = INT_MIN;
    for (int i = 0; i < numsSize; i++)
    {
        if (nums[i] < 0)
        {
            int temp = maxproduct;
            maxproduct = minproduct;
            minproduct = temp;
        }
        maxproduct = MAX(maxproduct * nums[i], nums[i]);
        minproduct = MIN(minproduct * nums[i], nums[i]);
        maxval = MAX(maxval, maxproduct);
    }
    free(dp);
    return maxval;
}
int main(int argc, char *argv[])
{
    int a[4] = {2, -1, 1, 1};
    printf("%d\n", maxProduct(a, 4));
    return 0;
}