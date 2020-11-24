#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MAX(x, y) (x < y ? y : x)
int maxScore(int *cardPoints, int cardPointsSize, int k)
{
    int *rc = (int *)calloc(cardPointsSize + 1, sizeof(int));
    int ans = 0;
    for (int i = 0; i < cardPointsSize; i++)
    {
        rc[i + 1] = rc[i] + cardPoints[i];
    }

    for (int i = 0; i <= k; i++)
    {
        int left = i;
        int right = k - i;
        int cur = rc[cardPointsSize] - rc[cardPointsSize - right] + rc[left];
        ans = MAX(ans, cur);
    }
    return ans;
}
int main(int argc, char *argv[])
{
    int cards[] = {2, 2, 2};
    int k = 2;
    printf("%d\n", maxScore(cards, 3, k));
    return 0;
}