#include <stdio.h>
#include <stdlib.h>

int canCompleteCircuit(int *gas, int gasSize, int *cost, int costSize)
{
    int totalGas = 0, can = -1, j;
    for (int i = 0; i < gasSize; i++)
    {
        if ((gas[i] - cost[i]) < 0)
        {
            continue;
        }
        j = i;
        totalGas = gas[j] - cost[j];
        j++;
        j = j % gasSize;
        while (j != i && totalGas >= 0)
        {
            totalGas = totalGas + gas[j] - cost[j];
            j++;
            j = j % gasSize;
        }
        if ((totalGas >= 0) && (j == i))
        {
            can = i;
            break;
        }
    }
    return can;
}

int main(int argc, char *argv[])
{
    int gas[3] = {3, 3, 4};
    int cost[3] = {3, 4, 4};
    printf("%d\n", canCompleteCircuit(gas, 3, cost, 3));
    return 0;
}