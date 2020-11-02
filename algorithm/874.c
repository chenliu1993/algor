#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#define max(a, b) ((a) > (b)) ? (a) : (b)

typedef struct node
{
    int x;
    int y;
} node;

int cmpfunc(const void *a, const void *b)
{
    node *ap = (node *)a;
    node *bp = (node *)b;
    if (ap->x != bp->x)
    {
        return ap->x - bp->x;
    }
    return ap->y - bp->y;
}

int getNextDir(int cur_dir, int command)
{
    if (command == -1)
    {
        cur_dir++;
    }
    else
    {
        cur_dir--;
    }
    int tmp = cur_dir % 4;
    return (tmp >= 0) ? tmp : (tmp + 4);
}

int robotSim(int *commands, int commandsSize, int **obstacles, int obstaclesSize, int *obstaclesColSize)
{
    int max_dir = 0;
    int curX = 0;
    int curY = 0;
    int dir[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}}; //上右下左
    int cur_dir = 0;

    node next;
    node *ob = (node *)malloc(sizeof(node) * obstaclesSize);

    for (int i = 0; i < obstaclesSize; i++)
    {
        ob[i].x = obstacles[i][0];
        ob[i].y = obstacles[i][1];
    }

    qsort(ob, obstaclesSize, sizeof(node), cmpfunc);

    for (int i = 0; i < commandsSize; i++)
    {
        if (commands[i] == -1 || commands[i] == -2)
        {
            cur_dir = getNextDir(cur_dir, commands[i]);
        }
        else if (commands[i] >= 1 && commands[i] <= 9)
        {
            while (commands[i]--)
            {
                next.x = curX + dir[cur_dir][0];
                next.y = curY + dir[cur_dir][1];
                node *item = (node *)bsearch(&next, ob, obstaclesSize, sizeof(node), cmpfunc);
                if (item != NULL)
                {
                    break;
                }
                else
                {
                    curX = next.x;
                    curY = next.y;
                }
            }
        }
        max_dir = max(max_dir, (curX * curX) + (curY * curY));
    }
    return max_dir;
}
