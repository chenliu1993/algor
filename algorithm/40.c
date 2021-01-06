#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int *ansColumnSize;
int ansSize;
int curSize;

void dfs(int *candidates, int candidatesSize, int **ans, int *cur, int target, int start, int sum, int *visited)
{
    if (sum >= target)
    {
        if (sum == target)
        {
            int *temp = (int *)malloc(curSize * sizeof(int));
            memcpy(temp, cur, curSize * sizeof(int));
            ans[ansSize] = temp;
            ansColumnSize[ansSize++] = curSize;
        }
        return;
    }
    for (int i = start; i < candidatesSize; i++)
    {
        if (i - 1 >= start && candidates[i] == candidates[i - 1])
        {
            continue;
        }
        if (visited[i] != 0)
        {
            continue;
        }
        visited[i] = 1;
        cur[curSize++] = candidates[i];
        dfs(candidates, candidatesSize, ans, cur, target, i + 1, sum + candidates[i], visited);
        visited[i] = 0;
        curSize--;
    }
}

int compareFunc(const void *a, const void *b)
{
    return (*(int *)a) < (*(int *)b);
}

int **combinationSum2(int *candidates, int candidatesSize, int target, int *returnSize, int **returnColumnSizes)
{
    qsort(candidates, candidatesSize, sizeof(int), compareFunc);
    int **ans = (int **)malloc(1001 * sizeof(int *));
    int *visited = (int *)calloc(candidatesSize, sizeof(int));
    ansColumnSize = (int *)malloc(1001 * sizeof(int));
    ansSize = curSize = 0;
    int cur[2001];
    dfs(candidates, candidatesSize, ans, cur, target, 0, 0, visited);
    *returnSize = ansSize;
    *returnColumnSizes = ansColumnSize;
    return ans;
}