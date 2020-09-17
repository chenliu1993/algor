#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
#include <string.h>

typedef struct _Point
{
    int x;
    int y;
} Point;

Point newPoint(int x, int y)
{
    Point p;
    p.x = x;
    p.y = y;
    return p;
}

int compmin(int a, int b)
{
    return a > b ? b : a;
}

int minPathSum(int **grid, int gridSize, int *gridColSize)
{
    // assume each line is the same size.
    int **visited = (int **)malloc(sizeof(int *) * gridSize);
    for (int i = 0; i < gridSize; i++)
    {
        visited[i] = (int *)malloc(sizeof(int) * gridColSize[i]);
        memset(visited[i], 0, sizeof(int) * gridColSize[i]);
    }
    int **sum = (int **)malloc(sizeof(int *) * gridSize);
    for (int i = 0; i < gridSize; i++)
    {
        sum[i] = (int *)malloc(sizeof(int) * gridColSize[i]);
        memset(sum[i], 0, sizeof(int) * gridColSize[i]);
    }
    int **inQ = (int **)malloc(sizeof(int *) * gridSize);
    for (int i = 0; i < gridSize; i++)
    {
        inQ[i] = (int *)malloc(sizeof(int) * gridColSize[i]);
        memset(inQ[i], 0, sizeof(int) * gridColSize[i]);
    }
    Point *queue = (Point *)malloc(sizeof(Point) * (gridColSize[0] * gridSize));
    memset(queue, 0, sizeof(Point) * (gridColSize[0] * gridSize));

    int front, rear;
    front = rear = 0;
    queue[rear++] = newPoint(0, 0);
    sum[0][0] = grid[0][0];
    visited[0][0] = 1;
    while (front != rear)
    {
        Point p = queue[front];
        if ((p.x + 1) < gridSize && visited[p.x + 1][p.y] == 0)
        {
            if (!inQ[p.x + 1][p.y])
            {
                queue[rear++] = newPoint(p.x + 1, p.y);
                inQ[p.x + 1][p.y] = 1;
            }
            if (sum[p.x + 1][p.y] == 0)
            {
                sum[p.x + 1][p.y] = sum[p.x][p.y] + grid[p.x + 1][p.y];
            }
            else
            {
                sum[p.x + 1][p.y] = compmin(sum[p.x + 1][p.y], sum[p.x][p.y] + grid[p.x + 1][p.y]);
            }
        }
        if ((p.y + 1) < gridColSize[0] && visited[p.x][p.y + 1] == 0)
        {
            if (!inQ[p.x][p.y + 1])
            {
                queue[rear++] = newPoint(p.x, p.y + 1);
                inQ[p.x][p.y + 1] = 1;
            }
            if (sum[p.x][p.y + 1] == 0)
            {
                sum[p.x][p.y + 1] = sum[p.x][p.y] + grid[p.x][p.y + 1];
            }
            else
            {
                sum[p.x][p.y + 1] = compmin(sum[p.x][p.y + 1], sum[p.x][p.y] + grid[p.x][p.y + 1]);
            }
        }
        visited[p.x][p.y] = 1;
        front++;
    }
    free(inQ);
    free(queue);
    free(visited);
    return sum[gridSize - 1][gridColSize[0] - 1];
}

// dp
// int minPathSum(int **grid, int gridSize, int *gridColSize)
// {
//     int **dp = grid;
//     dp[0][0] = grid[0][0];
//     for (int i = 1; i < gridSize; i++)
//     {
//         dp[i][0] = dp[i - 1][0] + grid[i][0];
//     }
//     for (int i = 1; i < gridColSize[0]; i++)
//     {
//         dp[0][i] = dp[0][i - 1] + grid[0][i];
//     }
//     for (int i = 1; i < gridSize; i++)
//     {
//         for (int j = 1; j < gridColSize[0]; j++)
//         {
//             dp[i][j] = fmin(dp[i - 1][j] + grid[i][j], dp[i][j - 1] + grid[i][j]);
//         }
//     }
//     return dp[gridSize - 1][gridColSize[0] - 1];
// }

int main(int argc, char *argv[])
{
    int **all = (int **)malloc(sizeof(int *) * 3);
    for (int i = 0; i < 3; i++)
    {
        all[i] = (int *)malloc(sizeof(int) * 3);
        memset(all[i], 0, sizeof(int) * 3);
    }
    int p1[3] = {1, 3, 1};
    int p2[3] = {1, 5, 1};
    int p3[3] = {4, 2, 1};
    all[0] = p1;
    all[1] = p2;
    all[2] = p3;
    int col[3] = {3, 3, 3};
    printf("%d\n", minPathSum(all, 3, col));
    return 0;
}