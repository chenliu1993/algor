#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
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

void DFS(int *nums, int numsSize, int *visited, int *ele, int **result, int count)
{
    if (count == numsSize)
    {
        result[rear] = (int *)malloc(numsSize * sizeof(int));
        memcpy(result[rear++], ele, numsSize * sizeof(int));
        return;
    }
    for (int i = 0; i < numsSize; i++)
    {

        if (visited[i])
        {
            continue;
        }
        ele[count] = nums[i];
        visited[i] = 1;
        DFS(nums, numsSize, visited, ele, result, count + 1);
        visited[i] = 0;
    }
}

int **permute(int *nums, int numsSize, int *returnSize, int **returnColumnSizes)
{
    int sum = 1;
    for (int i = 1; i <= numsSize; i++)
    {
        sum *= i;
    }
    int *visited = (int *)calloc(numsSize, sizeof(int));
    int **rc = (int **)malloc(sum * sizeof(int *));

    int *ele = (int *)calloc(numsSize, sizeof(int));
    rear = 0;
    DFS(nums, numsSize, visited, ele, rc, 0);

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
    int a[3] = {1, 2, 3};
    int *returnSize = (int *)calloc(1, sizeof(int));
    int **returnColSize = (int **)calloc(1, sizeof(int *));
    printGrid(permute(a, 3, returnSize, returnColSize), 6, 3);
    return 0;
}