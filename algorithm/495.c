#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int findPoisonedDuration(int *timeSeries, int timeSeriesSize, int duration)
{
    if (timeSeries == NULL || timeSeriesSize == 0)
    {
        return 0;
    }
    if (timeSeriesSize == 1)
    {
        return duration;
    }
    int cur = duration, sum = 0;
    for (int i = 0; i < timeSeriesSize - 1; i++)
    {
        if (timeSeries[i + 1] - timeSeries[i] >= cur)
        {
            sum += cur;
        }
        else
        {
            sum += timeSeries[i + 1] - timeSeries[i];
            cur = duration;
        }
    }
    return sum + duration;
}
int main(int argc, char *argv[])
{
    int a[2] = {1, 2};
    int d = 2;
    printf("%d\n", findPoisonedDuration(a, 2, d));
    return 0;
}