#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
int canPlaceFlowers(int *flowerbed, int flowerbedSize, int n)
{
    int count = 0, i; //计数器
    if (flowerbedSize == 1 && flowerbed[0] == 0)
        count++;
    else
    {
        for (i = 0; i < flowerbedSize; i++)
        {
            if (flowerbed[i] == 0)
            {
                if (i == 0) //首元素,处理左边界
                {
                    if (flowerbed[1] == 0)
                    {
                        count++;
                        flowerbed[i] = 1;
                    }
                }
                else if (i == flowerbedSize - 1) //处理右边界
                {
                    if (flowerbed[i - 1] == 0)
                        count++;
                    flowerbed[i] = 1;
                }
                else if (flowerbed[i - 1] == 0 && flowerbed[i + 1] == 0) //处理中间元素
                {
                    count++;
                    flowerbed[i] = 1;
                }
            }
        }
    }
    if (count >= n)
        return 1;
    return 0;
}