#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <limits.h>
#define MIN(x, y) (x < y ? x : y)
int minCostClimbingStairs(int *cost, int costSize)
{
    if (costSize < 2 || cost == NULL)
    {
        return 0;
    }
    int *dp = (int *)calloc(costSize + 1, sizeof(int));
    for (int i = 2; i < costSize + 1; i++)
    {
        dp[i] = MIN(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
    }
    return dp[costSize];
}

int main(int argc, char *argv[])
{
    int cost[10] = {1, 100, 1, 1, 1, 100, 1, 1, 100, 1};
    printf("%d\n", minCostClimbingStairs(cost, 10));
    return 0;
}