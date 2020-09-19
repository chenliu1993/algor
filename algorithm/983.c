#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/**
 * dp[366], dp[0]=0; dp[1]=cost[0];
 * dp[i] means minimum cost for days[0] to days[i].
 * if i in days[i], then:
 * dp[i]=min(dp[i-1]+cost[0], dp[i-7]+cost[1], dp[i-30]+cost[2]);
 * if [...<0], then cost is cost[...]
 * else:
 * dp[i]=dp[i-1];
 **/
#define MIN(a, b) ((a < b) ? a : b)
int mincostTickets(int *days, int daysSize, int *costs, int costsSize)
{
    int *calendar = (int *)malloc(sizeof(int) * 366);
    memset(calendar, 0, sizeof(int) * 366);
    int *dp = (int *)malloc(sizeof(int) * 366);
    memset(dp, 0, sizeof(int) * 366);
    for (int i = 0; i < daysSize; i++)
    {
        calendar[days[i]] = 1;
    }
    dp[0] = 0;
    for (int i = 1; i < 366; i++)
    {
        if (calendar[i])
        {
            int days_30_cost = (i - 30) < 0 ? costs[2] : dp[i - 30] + costs[2];
            int days_7_cost = (i - 7) < 0 ? costs[1] : dp[i - 7] + costs[1];
            int days_1_cost = (i - 1) < 0 ? costs[0] : dp[i - 1] + costs[0];
            dp[i] = MIN(MIN(days_1_cost, days_7_cost), days_30_cost);
        }
        else
        {
            dp[i] = dp[i - 1];
        }
    }
    return dp[days[daysSize - 1]];
}

int main(int argc, char *argv[])
{
    int days[6] = {1, 4, 6, 7, 8, 20};
    int costs[3] = {2, 7, 15};
    printf("%d\n", mincostTickets(days, 6, costs, 6));
    return 0;
}