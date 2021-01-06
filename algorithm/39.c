#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int ansSize;

int curSize;

int *ansColumnSize;
void dfs(int *candidates, int candidatesSize, int **ans, int *cur, int target, int idx)
{
	if (idx == candidatesSize)
	{
		return;
	}
	if (target == 0)
	{
		int *temp = (int *)malloc(curSize * sizeof(int));
		memcpy(temp, cur, curSize * sizeof(int));
		ans[ansSize] = temp;
		ansColumnSize[ansSize++] = curSize;
		return;
	}
	dfs(candidates, candidatesSize, ans, cur, target, idx + 1);
	if (target - candidates[idx] >= 0)
	{
		cur[curSize++] = candidates[idx];
		dfs(candidates, candidatesSize, ans, cur, target - candidates[idx], idx);
		curSize--;
	}
}

int **combinationSum(int *candidates, int candidatesSize, int target, int *returnSize, int **returnColumnSizes)
{
	ansSize = curSize = 0;
	int **ans = malloc(1001 * sizeof(int *));
	ansColumnSize = malloc(1001 * sizeof(int));
	int cur[2001];
	dfs(candidates, candidatesSize, ans, cur, target, 0);
	*returnSize = ansSize;
	*returnColumnSizes = ansColumnSize;
	return ans;
}