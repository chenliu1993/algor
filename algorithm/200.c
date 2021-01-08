#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

void dfs(char **grid, int gridSize, int gridColSize, int **visited, int x, int y)
{
    if (x < 0 || x >= gridSize)
    {
        return;
    }
    if (y < 0 || y >= gridColSize)
    {
        return;
    }
    if (grid[x][y] == '0' || visited[x][y] == 1)
    {
        return;
    }
    visited[x][y] = 1;
    dfs(grid, gridSize, gridColSize, visited, x - 1, y);
    dfs(grid, gridSize, gridColSize, visited, x + 1, y);
    dfs(grid, gridSize, gridColSize, visited, x, y - 1);
    dfs(grid, gridSize, gridColSize, visited, x, y + 1);
}

int numIslands(char **grid, int gridSize, int *gridColSize)
{
    int m = gridSize;
    int n = gridColSize[0];
    int **visited = (int **)malloc(m * sizeof(int *));
    for (int i = 0; i < m; i++)
    {
        visited[i] = (int *)calloc(n, sizeof(int));
    }
    int count = 0;
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (grid[i][j] != '0' && visited[i][j] != 1)
            {
                dfs(grid, m, n, visited, i, j);
                count++;
            }
        }
    }
    return count;
}