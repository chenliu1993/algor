#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? x : y)
int maxValue(int **grid, int gridSize, int *gridColSize)
{
    int m = gridSize;
    int n = gridColSize[0];
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (i == 0 && j == 0)
            {
                continue;
            }
            if (i == 0)
            {
                grid[i][j] += grid[i][j - 1];
            }
            else if (j == 0)
            {
                grid[i][j] += grid[i - 1][j];
            }
            else
            {
                grid[i][j] += MAX(grid[i - 1][j], grid[i][j - 1]);
            }
        }
    }
    return grid[m - 1][n - 1];
}

int main(int argc, char *argv[])
{
    int m = 3, n = 3;
    int *colSize = (int *)calloc(1, sizeof(int));
    colSize[0] = n;
    int **a = (int **)malloc(m * sizeof(int *));
    int a1[3] = {1, 3, 1};
    int a2[3] = {1, 5, 1};
    int a3[3] = {4, 2, 1};
    a[0] = a1;
    a[1] = a2;
    a[2] = a3;
    printf("%d\n", maxValue(a, m, colSize));
    return 0;
}