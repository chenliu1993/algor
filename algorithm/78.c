#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int **ans;
int *ansColSize;
int ansSize;

int *t;
int tSize;

void dfs(int cur, int *nums, int numsSize)
{
    if (cur == numsSize)
    {
        int *tmp = malloc(sizeof(int) * tSize);
        memcpy(tmp, t, sizeof(int) * tSize);
        ansColSize[ansSize] = tSize;
        ans[ansSize++] = tmp;
        return;
    }
    t[tSize++] = nums[cur];
    dfs(cur + 1, nums, numsSize);
    tSize--;
    dfs(cur + 1, nums, numsSize);
}

int **subsets(int *nums, int numsSize, int *returnSize, int **returnColumnSizes)
{
    ans = malloc(sizeof(int *) * (1 << numsSize));
    ansColSize = malloc(sizeof(int) * (1 << numsSize));
    t = malloc(sizeof(int) * numsSize);
    *returnSize = 1 << numsSize;
    ansSize = tSize = 0;
    dfs(0, nums, numsSize);
    *returnColumnSizes = ansColSize;
    return ans;
}
