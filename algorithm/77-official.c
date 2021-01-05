#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

int *temp;
int tempSize;

int **ans;
int ansSize;

void printGrid(int **a, int aSize, int aColSize)
{
    for (int i = 0; i < aSize; i++)
    {

        for (int j = 0; j < aColSize; j++)
        {
            printf("%d,", a[i][j]);
        }
        printf("\n");
    }
    printf("\n");
}

void dfs(int cur, int n, int k)
{
    // 剪枝：temp 长度加上区间 [cur, n] 的长度小于 k，不可能构造出长度为 k 的 temp
    if (tempSize + (n - cur + 1) < k)
    {
        return;
    }
    // 记录合法的答案
    if (tempSize == k)
    {
        int *tmp = malloc(sizeof(int) * k);
        for (int i = 0; i < k; i++)
        {
            tmp[i] = temp[i];
        }
        ans[ansSize++] = tmp;
        return;
    }
    // 考虑选择当前位置
    temp[tempSize++] = cur;
    dfs(cur + 1, n, k);
    tempSize--;
    // 考虑不选择当前位置
    dfs(cur + 1, n, k);
}

int **combine(int n, int k, int *returnSize, int **returnColumnSizes)
{
    temp = malloc(sizeof(int) * k);
    ans = malloc(sizeof(int *) * 10001);
    tempSize = ansSize = 0;
    dfs(1, n, k);
    *returnSize = ansSize;
    *returnColumnSizes = malloc(sizeof(int) * ansSize);
    for (int i = 0; i < ansSize; i++)
    {
        (*returnColumnSizes)[i] = k;
    }
    return ans;
}

int main(int argc, char *argv[])
{
    int n = 4;
    int k = 2;
    int *returnSize = (int *)malloc(sizeof(int));
    memset(returnSize, 0, sizeof(int));
    int **returnColumnSizes = (int **)malloc(n * sizeof(int *));
    int **res = combine(n, k, returnSize, returnColumnSizes);
    printGrid(res, (*returnSize), k);
    return 0;
}