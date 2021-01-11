#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX_SIZE 32768
int compareFunc(const void *a, const void *b)
{
    return (*(int *)a) > (*(int *)b);
}

int **ans;
int rear;

void dfs(int *nums, int numsSize, int last, int startID, int *ele, int *returnSize, int **returnColumnSizes)
{
    if (startID == numsSize)
    {
        if (rear > 1)
        {
            (*returnSize)++;
            ans[(*returnSize) - 1] = malloc(rear * sizeof(int));
            memcpy(ans, ele, rear * sizeof(int));
            (*returnColumnSizes)[(*returnSize) - 1] = rear;
        }
        return;
    }
    if (nums[startID] >= last)
    {
        ele[rear++] = nums[startID];
        dfs(nums, numsSize, nums[startID], startID + 1, ele, returnSize, returnColumnSizes);
        rear--;
    }
    if (nums[startID] != last)
    {
        dfs(nums, numsSize, last, startID + 1, ele, returnSize, returnColumnSizes);
    }
}

int **findSubsequences(int *nums, int numsSize, int *returnSize, int **returnColumnSizes)
{
    qsort(nums, numsSize, sizeof(int), compareFunc);
    ans = (int **)malloc(MAX_SIZE * sizeof(int *));
    int *ele = (int *)malloc(numsSize * sizeof(int));
    (*returnSize) = 0;
    rear = 0;
    dfs(nums, numsSize, INT_MIN, 0, ele, returnSize, returnColumnSizes);
    return ans;
}

int main(int argc, char *argv[])
{
    int nums[4] = {4, 6, 7, 7};
    int numsSize = 4;
    int *returnSize = (int *)calloc(1, sizeof(int));
    int **returnColumnSize = (int **)malloc(MAX_SIZE * sizeof(int *));
    findSubsequences(nums, numsSize, returnSize, returnColumnSize);
    for (int i = 0; i < (*returnSize); i++)
    {
        printf("%d,", (*(returnColumnSize[i])));
    }
    return 0;
}