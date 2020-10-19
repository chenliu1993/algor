#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int rear;
int compareFunc(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}
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
void DFS(int *nums, int numsSize, int **rc, int *ele, int *visited, int count)
{
    if (count == numsSize)
    {
        rc[rear] = (int *)malloc(sizeof(int) * numsSize);
        memcpy(rc[rear++], ele, numsSize * sizeof(int));
        return;
    }
    for (int i = 0; i < numsSize; i++)
    {
        if (visited[i] || (i > 0 && nums[i] == nums[i - 1] && !visited[i - 1]))
        {
            continue;
        }
        ele[count] = nums[i];
        visited[i] = 1;
        DFS(nums, numsSize, rc, ele, visited, count + 1);
        visited[i] = 0;
    }
}
int **permuteUnique(int *nums, int numsSize, int *returnSize, int **returnColumnSizes)
{
    int sum = 1;
    for (int i = 1; i <= numsSize; i++)
    {
        sum *= i;
    }
    int **rc = (int **)malloc(sum * sizeof(int *));
    int *ele = (int *)calloc(numsSize, sizeof(int));
    int *visited = (int *)calloc(numsSize, sizeof(int));
    rear = 0;
    qsort(nums, numsSize, sizeof(int), compareFunc);
    DFS(nums, numsSize, rc, ele, visited, 0);
    *returnSize = sum;
    *returnColumnSizes = (int *)malloc(sum * sizeof(int));

    for (int i = 0; i < sum; i++)
    {
        (*returnColumnSizes)[i] = numsSize;
    }
    return rc;
}
int main(int argc, char *argv[])
{
    int a[3] = {1, 1, 2};
    int *returnSize = (int *)calloc(1, sizeof(int));
    int **returnColSize = (int **)calloc(1, sizeof(int *));
    printGrid(permuteUnique(a, 3, returnSize, returnColSize), 3, 3);
    return 0;
}