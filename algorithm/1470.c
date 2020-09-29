#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int *shuffle(int *nums, int numsSize, int n, int *returnSize)
{
    // numsSize always equal to 2*n
    if (nums == NULL || n == 1)
    {
        return nums;
    }
    *returnSize = 2 * n;
    // int *rc = (int *)malloc(sizeof(int) * 2 * n);
    // memset(rc, 0, sizeof(int) * 2 * n);
    int j = 0;
    int *rc = (int *)calloc(2 * n, sizeof(int));
    for (int i = 0; i < n; i++)
    {
        rc[j++] = nums[i];
        rc[j++] = nums[i + n];
    }
    return rc;
}
void printArray(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        printf("%d,", arr[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int a[2] = {2, 7};
    int n = 1;
    int *re = (int *)calloc(1, sizeof(int));
    int *rc = shuffle(a, 2, n, re);
    printArray(rc, 2 * n);
    return 0;
}