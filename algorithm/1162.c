#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include <limits.h>
#include <string.h>
#include <math.h>
#define MAXSIZE 10000
#define MAXGRID 100
typedef struct _node
{
    int x;
    int y;
} node;
int ManhattanDist(node n1, node n2)
{
    return abs(n1.x - n2.x) + abs(n1.y - n2.y);
}

node newPosition(int x, int y)
{
    node *n = (node *)malloc(sizeof(node));
    memset(n, 0, sizeof(node));
    n->x = x;
    n->y = y;
    return *n;
}

int maxDistance(int **grid, int gridSize, int *gridColSize)
{
    int m = gridSize;
    int n = *gridColSize;
    node *queue = (node *)malloc(MAXSIZE * sizeof(node));
    memset(queue, 0, MAXSIZE * sizeof(node));
    int **visited = (int **)malloc(MAXGRID * sizeof(int *));
    for (int i = 0; i < MAXGRID; i++)
    {
        visited[i] = (int *)malloc(MAXGRID * sizeof(int));
        memset(visited[i], 0, MAXGRID * sizeof(int));
    }
    int front = 0, rear = 0;
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (grid[i][j] == 1)
            {
                visited[i][j] = 1;
                queue[rear++] = newPosition(i, j);
            }
        }
    }
    if (rear == 0 || rear == m * n)
    {
        return -1;
    }
    int step = -1;
    while (front < rear)
    {
        int tmp = rear;
        for (int i = front; i < tmp; i++)
        {
            int row = queue[i].x;
            int col = queue[i].y;
            if (row > 0 && grid[row - 1][col] == 0 && visited[row - 1][col] != 1)
            {
                queue[rear++] = newPosition(row - 1, col);
                visited[row - 1][col] = 1;
            }
            if (row < m - 1 && grid[row + 1][col] == 0 && visited[row + 1][col] == 0)
            {
                queue[rear++] = newPosition(row + 1, col);
                visited[row + 1][col] = 1;
            }
            if (col > 0 && grid[row][col - 1] != 1 && visited[row][col - 1] == 0)
            {
                queue[rear++] = newPosition(row, col - 1);
                visited[row][col - 1] = 1;
            }
            if (col < n - 1 && grid[row][col + 1] == 0 && visited[row][col + 1] == 0)
            {
                queue[rear++] = newPosition(row, col + 1);
                visited[row][col + 1] = 1;
            }
        }
        step++;
        front = tmp;
    }
    return step;
}

int main(int argc, char *argv[])
{
    int gridsSize = 3;
    int *gridsColSize = (int *)calloc(1, sizeof(int));
    *gridsColSize = 3;
    int **grids = (int **)malloc(gridsSize * sizeof(int *));
    for (int i = 0; i < gridsSize; i++)
    {
        grids[i] = (int *)malloc(*gridsColSize * sizeof(int));
        memset(grids[i], 0, *gridsColSize * sizeof(int));
    }
    int g1[3] = {1, 0, 1};
    int g2[3] = {0, 0, 0};
    int g3[3] = {1, 0, 1};
    grids[0] = g1;
    grids[1] = g2;
    grids[2] = g3;
    printf("%d\n", maxDistance(grids, gridsSize, gridsColSize));
    return 0;
}