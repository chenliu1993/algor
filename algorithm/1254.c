#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

void deleteBorderIsland(int **grid, int gridSize, int *gridColSize)
{
    for (int i = 0; i < gridSize; i++)
    {
        if (grid[i][0] == 0)
        {
            dfs(grid, gridSize, gridColSize, i, 0);
        }
        if (grid[i][*gridColSize - 1] == 0)
        {
            dfs(grid, gridSize, gridColSize, i, *gridColSize - 1);
        }
    }
    for (int i = 0; i < *gridColSize; i++)
    {
        if (grid[0][i] == 0)
        {
            dfs(grid, gridSize, gridColSize, 0, i);
        }
        if (grid[gridSize - 1][i] == 0)
        {
            dfs(grid, gridSize, gridColSize, gridSize - 1, i);
        }
    }
}

void dfs(int **grid, int gridSize, int *gridColSize, int x, int y)
{
    if (x < 0 || x > (gridSize - 1))
    {
        return;
    }
    if (y < 0 || y > ((*gridColSize) - 1))
    {
        return;
    }
    if (grid[x][y] == 1)
    {
        return;
    }
    grid[x][y] = 1;
    dfs(grid, gridSize, gridColSize, x - 1, y);
    dfs(grid, gridSize, gridColSize, x + 1, y);
    dfs(grid, gridSize, gridColSize, x, y - 1);
    dfs(grid, gridSize, gridColSize, x, y + 1);
}

int closedIsland(int **grid, int gridSize, int *gridColSize)
{
    deleteBorderIsland(grid, gridSize, gridColSize);
    int count = 0;
    for (int i = 0; i < gridSize; i++)
    {
        for (int j = 0; j < (*gridColSize); j++)
        {
            if (grid[i][j] == 0)
            {
                dfs(grid, gridSize, gridColSize, i, j);
                count++;
            }
        }
    }
    return count;
}