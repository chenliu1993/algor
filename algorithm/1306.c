#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* *
 * BFS
 * */

int canReach(int *arr, int arrSize, int start)
{
    int front = 0, rear = 0, can = 0;
    int *queue = (int *)malloc((arrSize + 1) * sizeof(int));
    int *visited = (int *)malloc((arrSize + 1) * sizeof(int));
    memset(queue, 0, (arrSize + 1) * sizeof(int));
    memset(visited, 0, (arrSize + 1) * sizeof(int));
    if (arr == NULL || arrSize == 0)
    {
        goto RETURN;
    }
    queue[rear++] = start;
    while (front != rear)
    {
        int end = rear;
        for (int i = front; i < end; i++)
        {
            if (arr[queue[i]] == 0)
            {
                can = 1;
                goto RETURN;
            }
            int prev = queue[i] - arr[queue[i]];
            int afte = queue[i] + arr[queue[i]];
            if ((prev >= 0) && visited[prev] == 0)
            {
                queue[rear++] = prev;
            }
            if ((afte < arrSize) && visited[afte] == 0)
            {
                queue[rear++] = afte;
            }
            visited[queue[i]] = 1;
            front++;
        }
    }
RETURN:
    free(visited);
    free(queue);
    return can;
}
int main(int argc, char *argv[])
{
    int a[5] = {3, 0, 2, 1, 2};
    int start = 2;
    printf("%d\n", canReach(a, 5, start));
    return 0;
}