#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int orangesRotting(int **grid, int gridSize, int *gridColSize)
{
    int m = gridSize, n = gridColSize[0], count = 0, end;
    while (1)
    {
        end = 1;
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                if (grid[i][j] == 3)
                {
                    grid[i][j] = 2;
                }
                else if (grid[i][j] == 1)
                {
                    if ((i < m - 1 && grid[i + 1][j] == 2) || (j < n - 1 && grid[i][j + 1] == 2))
                    {
                        grid[i][j] = 2;
                        end = 0;
                    }
                }
                else if (grid[i][j] == 2)
                {
                    if (i < m - 1 && grid[i + 1][j] == 1)
                    {
                        grid[i + 1][j] = 3;
                        end = 0;
                    }
                    if (j < n - 1 && grid[i][j + 1] == 1)
                    {
                        grid[i][j + 1] = 3;
                        end = 0;
                    }
                }
            }
        }
        if (end)
        {
            break;
        }
        else
        {
            count++;
        }
    }
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {

            if (grid[i][j] == 1)
            {
                return -1;
            }
        }
    }
    return count;
}