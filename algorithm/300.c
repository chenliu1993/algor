#include <stdio.h>
#include <stdlib.h>

int lengthOfLIS(int *nums, int numsSize)
{
    int maxlen = 0;
    if (nums == NULL || numsSize == 0)
    {
        return 0;
    }
    int *res = (int *)malloc(sizeof(int) * numsSize);
    for (int i = 0; i < numsSize; i++)
    {
        res[i] = 1;
        for (int j = 0; j < i; j++)
        {
            if (nums[i] > nums[j] && res[j] + 1 > res[i])
            {
                res[i] = res[j] + 1;
            }
        }
        maxlen = maxlen > res[i] ? maxlen : res[i];
    }
    free(res);
    return maxlen;
}

int main(int argc, char *argv[])
{
    int a[8] = {10, 9, 2, 5, 3, 7, 101, 18};
    printf("%d\n", lengthOfLIS(a, 8));
    return 0;
}