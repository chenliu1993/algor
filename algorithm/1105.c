#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
#define MAX(x, y) (x < y ? y : x)
int minHeightShelves(int **books, int booksSize, int *booksColSize, int shelf_width)
{
    int temp_width, curheight;
    int *dp = (int *)calloc(booksSize + 1, sizeof(int));
    for (int i = 1; i < booksSize + 1; i++)
    {
        dp[i] = INT_MAX;
    }
    for (int i = 1; i < booksSize + 1; i++)
    {
        temp_width = 0;
        curheight = 0;
        for (int j = i; j > 0; j--)
        {
            temp_width += books[j - 1][0];
            curheight = MAX(curheight, books[j - 1][1]);
            if (temp_width <= shelf_width)
            {
                dp[i] = MIN(dp[i], dp[j - 1] + curheight);
            }
        }
    }
    return dp[booksSize];
}
int main(int argc, char *argv[])
{
    return 0;
}