#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int isBoomerang(int **points, int pointsSize, int *pointsColSize)
{
    double l1 = sqrt((points[0][0] - points[1][0]) * (points[0][0] - points[1][0]) + (points[0][1] - points[1][1]) * (points[0][1] - points[1][1]));
    double l2 = sqrt((points[0][0] - points[2][0]) * (points[0][0] - points[2][0]) + (points[0][1] - points[2][1]) * (points[0][1] - points[2][1]));
    double l3 = sqrt((points[2][0] - points[1][0]) * (points[2][0] - points[1][0]) + (points[2][1] - points[1][1]) * (points[2][1] - points[1][1]));
    if (((l1 + l2) > l3) && ((l1 + l3) > l2) && ((l2 + l3) > l1))
    {
        return 1;
    }
    return 0;
}

int main(int argc, char *argv[])
{
    int **points = (int **)malloc(3 * sizeof(int));
    int a1[2] = {1, 1};
    int a2[2] = {2, 3};
    int a3[2] = {3, 2};
    points[0] = a1;
    points[1] = a2;
    points[2] = a3;
    int *pointsColSize = (int *)calloc(3, sizeof(int));
    for (int i = 0; i < 3; i++)
    {
        pointsColSize[i] = 2;
    }
    printf("%d\n", isBoomerang(points, 3, pointsColSize));
    return 0;
}