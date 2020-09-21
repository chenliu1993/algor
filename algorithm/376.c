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

int wiggleMaxLength(int *nums, int numsSize)
{
    if (nums == NULL || numsSize == 0)
    {
        return 0;
    }
    if (numsSize == 1)
    {
        return 1;
    }
    int *sub = (int *)malloc(sizeof(int) * numsSize);
    memset(sub, 0, sizeof(int) * numsSize);
    int rear = 1, cur = 0, old = 0;
    sub[0] = nums[0];
    for (int i = 1; i < numsSize; i++)
    {
        cur = nums[i] - nums[i - 1];
        if ((cur > 0 && old <= 0) || (cur < 0 && old >= 0))
        {
            old = cur;
            sub[rear++] = nums[i];
        }
    }
    printArray(sub, rear);
    free(sub);
    return rear;
}

int main(int argc, char *argv[])
{
    int a[6] = {1, 7, 4, 9, 2, 5};
    printf("%d\n", wiggleMaxLength(a, 6));
    return 0;
}