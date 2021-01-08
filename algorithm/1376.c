#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
int numOfMinutes(int n, int headID, int *manager, int managerSize, int *informTime, int informTimeSize)
{
    if (n == 1)
        return 0;
    int timeSpent = 0;
    int target, clock;
    for (int i = 0; i < informTimeSize; i++)
    {
        if (informTime[i] == 0)
        {
            target = manager[i];
            clock = 0;
            while (target != headID)
            {
                clock += informTime[target];
                target = manager[target];
            }
            clock += informTime[target];
            timeSpent = MAX(timeSpent, clock);
        }
    }
    return timeSpent;
}

int main(int argc, char *argv[])
{
    int n = 6;
    int headID = 2;
    int manager[6] = {2, 2, -1, 2, 2, 2};
    int informTime[6] = {0, 0, 1, 0, 0, 0};
    printf("%d\n", numOfMinutes(n, headID, manager, 6, informTime, 6));
    return 0;
}