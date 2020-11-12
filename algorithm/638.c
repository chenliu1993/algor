#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
#define MIN(x, y) (x < y ? x : y)
int getSum(int *price, int priceSize, int *needs, int needsSize)
{
    int sum = 0;
    for (int i = 0; i < priceSize; i++)
    {
        sum += price[i] * needs[i];
    }
    return sum;
}
int shopping(int *price, int priceSize, int **special, int specialSize, int *specialColSize, int *needs, int needsSize)
{
    int j = 0, rc = getSum(price, priceSize, needs, needsSize);
    for (int i = 0; i < specialSize; i++)
    {
        int *needsClone = (int *)malloc(needsSize * sizeof(int));
        memcpy(needsClone, needs, needsSize);
        for (j = 0; j < needsSize; j++)
        {
            int productDiff = needs[j] - special[i][j];
            if (productDiff < 0)
            {
                break;
            }
            needsClone[j] = productDiff;
        }
        if (j == needsSize)
        {
            rc = MIN(rc, special[i][j] + shopping(price, priceSize, special, specialSize, specialColSize, needsClone, needsSize));
        }
    }
    return rc;
}

int shoppingOffers(int *price, int priceSize, int **special, int specialSize, int *specialColSize, int *needs, int needsSize)
{
    return shopping(price, priceSize, special, specialSize, specialColSize, needs, needsSize);
}

int main(int argc, char *argv[])
{
    int price[2] = {2, 5};
    int needs[2] = {3, 2};
    int s1[3] = {3, 0, 5};
    int s2[3] = {1, 2, 10};
    int ssize = 2;
    int *scolsize = (int *)calloc(1, sizeof(int));
    *scolsize = 3;
    int **s = (int **)malloc(ssize * sizeof(int *));
    for (int i = 0; i < ssize; i++)
    {
        s[i] = (int *)calloc(*scolsize, sizeof(int));
    }
    s[0] = s1;
    s[1] = s2;
    printf("%d\n", shoppingOffers(price, 2, s, ssize, scolsize, needs, 2));
    return 0;
}