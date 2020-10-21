#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void qick(int *a, int asize, int start, int end)
{
    if (start >= end)
    {
        return;
    }
    int left = start, right = end - 1, standard = a[end];
    while (left < right)
    {
        while (left < right && a[left] < standard)
        {
            left++;
        }
        while (left < right && a[right] > standard)
        {
            right--;
        }
        swap(&a[left], &a[right]);
    }
    if (a[left] > standard)
    {
        swap(&a[left], &a[end]);
    }
    else
    {
        left++;
    }
    if (left)
    {
        qick(a, asize, 0, left - 1);
    }
    qick(a, asize, left + 1, end);
}
void qicksort(int *a, int asize)
{
    qick(a, asize, 0, asize - 1);
}
int compareFunc(const void *a, const void *b)
{
    return *(int *)a > *(int *)b;
}
int canForm(int a, int b, int c)
{
    if ((a + b) > c && (a + c) > b && (b + c) > a)
    {
        return 1;
    }
    return 0;
}

int largestPerimeter(int *A, int ASize)
{
    if (ASize < 3)
    {
        return 0;
    }
    int maxlen = INT_MIN;
    // qsort(A, ASize, sizeof(int), compareFunc);
    qicksort(A, ASize);
    for (int i = ASize - 1; i >= 2; i--)
    {
        if (canForm(A[i], A[i - 1], A[i - 2]))
        {
            maxlen = maxlen > (A[i] + A[i - 1] + A[i - 2]) ? maxlen : (A[i] + A[i - 1] + A[i - 2]);
        }
    }
    if (maxlen == INT_MIN)
    {
        return 0;
    }
    return maxlen;
}

int main(int argc, char *argv[])
{
    int a[4] = {3, 2, 3, 4};
    printf("%d\n", largestPerimeter(a, 4));
    return 0;
}