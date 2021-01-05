#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int rear;

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

void DFS(int n, int k, int *ele, int **res, int *visited, int count, int *returnSize, int **returnColSizes)
{
    if (count == k)
    {
        (*returnSize)++;
        res = realloc(res, (*returnSize) * sizeof(int *));
        res[(*returnSize) - 1] = (int *)malloc(k * sizeof(int));
        memcpy(res[(*returnSize) - 1], ele, k * sizeof(int));
        returnColSizes[(*returnSize) - 1] = &k;
        return;
    }
    for (int i = 1; i <= n; i++)
    {
        if (visited[i] != 0)
        {
            continue;
        }
        visited[i] = 1;
        ele[count] = i;
        DFS(n, k, ele, res, visited, count + 1, returnSize, returnColSizes);
        visited[i] = 0;
    }
}
int **combine(int n, int k, int *returnSize, int **returnColumnSizes)
{
    int *visited = (int *)calloc(n + 1, sizeof(int));
    int *ele = (int *)malloc(k * sizeof(int));
    int **res = malloc(0);
    rear = 0;
    DFS(n, k, ele, res, visited, 0, returnSize, returnColumnSizes);
    return res;
}
int main(int argc, char *argv[])
{
    int n = 4;
    int k = 2;
    int *returnSize = (int *)malloc(sizeof(int));
    memset(returnSize, 0, sizeof(int));
    int **returnColumnSizes = (int **)malloc(50000 * sizeof(int *));
    int **res = combine(n, k, returnSize, returnColumnSizes);
    printGrid(res, (*returnSize), k);
    return 0;
}