#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int sumR(int *nums, int end)
{
    int sum = 0;
    for (int i = 0; i <= end; i++)
    {
        sum += nums[i];
    }
    return sum;
}
int *runningSum(int *nums, int numsSize, int *returnSize)
{
    int *rc = (int *)calloc(numsSize, sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        rc[i] = sumR(nums, i);
    }
    *returnSize = numsSize;
    return rc;
}
int main(int argc, char *argv[])
{
    int a[5] = {1, 1, 1, 1, 1};
    printArray(a, 5);
    int *returnSize = (int *)calloc(1, sizeof(int));
    int *b = runningSum(a, 5, returnSize);
    printArray(b, 5);
    return 0;
}