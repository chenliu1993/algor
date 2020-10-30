#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int direction[4][2] = {{1, 0}, {0, 1}, {-1, 0}, {0, -1}};
int islandPerimeter(int **grid, int gridSize, int *gridColSize)
{
    int m = gridSize;
    int n = gridColSize[0];
    int circle = 0, newi, newj, curlen;
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (grid[i][j])
            {
                curlen = 0;
                for (int k = 0; k < 4; k++)
                {
                    newi = i + direction[k][0];
                    newj = j + direction[k][1];
                    // here grid[newi][newj] must be put at the end of condition, otherwise position newi,newj will be over grid's border since this will be checked first.
                    if (newi < 0 || newj >= n || newi >= m || newj < 0 || !grid[newi][newj])
                    {
                        curlen++;
                    }
                }
                circle += curlen;
            }
        }
    }
    return circle;
}
int main(int argc, char *argv[])
{
    int **a = (int **)malloc(4 * sizeof(int *));
    for (int i = 0; i < 4; i++)
    {
        a[i] = (int *)calloc(4, sizeof(int));
    }
    int a1[4] = {0, 1, 0, 0};
    int a2[4] = {1, 1, 1, 0};
    int a3[4] = {0, 1, 0, 0};
    int a4[4] = {1, 1, 0, 0};
    a[0] = a1;
    a[1] = a2;
    a[2] = a3;
    a[3] = a4;
    int *colSize = (int *)calloc(4, sizeof(int));
    for (int i = 0; i > 4; i++)
    {
        colSize[i] = 4;
    }
    printf("%d\n", islandPerimeter(a, 4, colSize));
    return 0;
}