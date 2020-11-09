#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>
#include <limits.h>
#include <string.h>

#define MAX_PRICE 10001
#define SIZE 10000
#define MIN(a, b) (a) < (b) ? (a) : (b)

typedef struct _Node
{
    int srcCity;
    int price;
    int cnt;
} Node;

int findCheapestPrice(int n, int **flights, int flightsSize, int *flightsColSize, int src, int dst, int K)
{
    Node *queue = (Node *)malloc(SIZE * sizeof(Node));
    memset(queue, 0, SIZE * sizeof(Node));
    int **g = (int **)malloc(n * sizeof(int *));
    for (int i = 0; i < n; i++)
    {
        g[i] = (int *)malloc(n * sizeof(int));
        memset(g[i], 0, n * sizeof(int));
    }
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            if (i == j)
            {
                g[i][j] = 0;
            }
            else
            {
                g[i][j] = MAX_PRICE;
            }
        }
    }
    for (int i = 0; i < flightsSize; i++)
    {
        g[flights[i][0]][flights[i][1]] = flights[i][2];
    }

    int front = 0, rear = 0;
    Node cur = {0}, next = {0};
    queue[rear].srcCity = src;
    queue[rear].price = 0;
    queue[rear++].cnt = -1;
    int rc = INT_MAX;
    while (front < rear)
    {
        cur.srcCity = queue[front].srcCity;
        cur.price = queue[front].price;
        cur.cnt = queue[front].cnt;
        if (cur.cnt > K)
        {
            front++;
            continue;
        }
        if (cur.srcCity == dst)
        {
            rc = MIN(rc, cur.price);
            front++;
            continue;
        }
        for (int i = 0; i < n; i++)
        {
            next.srcCity = i;
            next.price = g[cur.srcCity][i] + cur.price;
            next.cnt = cur.cnt + 1;
            if (cur.srcCity != next.srcCity && g[cur.srcCity][next.srcCity] != MAX_PRICE && next.price < rc)
            {
                queue[rear].srcCity = next.srcCity;
                queue[rear].price = next.price;
                queue[rear++].cnt = next.cnt;
            }
        }
        front++;
    }
    if (rc == INT_MAX)
    {
        return -1;
    }
    return rc;
}

int main(int argc, char *argv[])
{
    int n = 3;
    int src = 0;
    int dst = 2;
    int K = 0;
    int edgesSize = 3;
    int *edgesColSize = (int *)calloc(1, sizeof(int));
    edgesColSize[0] = 3;
    int **edges = (int **)malloc(edgesSize * sizeof(int *));
    for (int i = 0; i < edgesSize; i++)
    {
        edges[i] = (int *)calloc(edgesColSize[0], sizeof(int));
    }
    int e1[3] = {0, 1, 100};
    int e2[3] = {1, 2, 100};
    int e3[3] = {0, 2, 500};
    edges[0] = e1;
    edges[1] = e2;
    edges[2] = e3;
    printf("%d\n", findCheapestPrice(n, edges, edgesSize, edgesColSize, src, dst, K));
    return 0;
}