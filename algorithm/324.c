#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
void insertSort(int *nums, int numsSize)
{
    int i = 1, j, val;
    for (; i < numsSize; i++)
    {
        val = nums[i];
        j = i - 1;
        while (j >= 0 && nums[j] < val)
        {
            nums[j + 1] = nums[j];
            j--;
        }
        nums[j + 1] = val;
    }
}

void wiggleSort(int *nums, int numsSize)
{
    insertSort(nums, numsSize);
    int l = 0, r = 0;
    int *s = (int *)malloc(sizeof(int) * numsSize);
    memset(s, 0, sizeof(int) * numsSize);
    for (int i = numsSize / 2; i < numsSize; i++)
    {
        s[r++] = nums[i];
        s[r++] = INT_MAX;
    }
    for (int i = 0; i < numsSize; i++)
    {
        if (s[i] == INT_MAX)
        {
            s[i] = nums[l++];
        }
    }
    for (int i = 0; i < numsSize; i++)
    {
        nums[i] = s[i];
    }
}
int main(int argc, char *argv[])
{
    int a[7] = {1, 1, 2, 1, 2, 2, 1};
    wiggleSort(a, 7);
    printArray(a, 7);
    return 0;
}