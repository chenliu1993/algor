#include <stdio.h>
#include <stdlib.h>
#include <string.h>
/**
 * Question: Why local run will get a result of 2?
 * */
int isSpecial(int x, int y, int **mat, int matSize, int matColSize)
{
    for (int i = 0; i < matSize; i++)
    {
        if (i == x)
        {
            continue;
        }
        printf("%d\n", mat[i][y]);
        if (mat[i][y])
        {
            return 0;
        }
    }
    for (int i = 0; i < matColSize; i++)
    {
        if (i == y)
        {
            continue;
        }
        if (mat[x][i])
        {
            return 0;
        }
    }
    return 1;
}
int numSpecial(int **mat, int matSize, int *matColSize)
{
    int m = matColSize[0], count = 0;
    for (int i = 0; i < matSize; i++)
    {
        for (int j = 0; j < m; j++)
        {
            if (mat[i][j] && isSpecial(i, j, mat, matSize, m))
            {
                printf("%d,%d\n", i, j);
                count++;
            }
        }
    }
    return count;
}

int main(int argc, char *argv[])
{
    int a1[3] = {1, 0, 0};
    int a2[3] = {0, 0, 1};
    int a3[3] = {1, 0, 0};
    int **a = (int **)calloc(3, sizeof(int *));
    for (int i = 0; i < 3; i++)
    {
        a[i] = (int *)calloc(1, sizeof(int));
    }
    a[0] = a1;
    a[1] = a2;
    a[3] = a3;
    int matColSize[1] = {3};
    printf("%d\n", numSpecial(a, 3, matColSize));
    return 0;
}