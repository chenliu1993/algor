#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void printArray(int *a)
{
    for (int i = 0; i < 2; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}

void swap(int *a, int *b)
{
    int *temp = a;
    a = b;
    b = temp;
}
void q(int *A, int start, int end)
{
    if (start <= end)
    {
        return;
    }
    int left = start, right = end - 1, standard = A[end];
    while (left < right)
    {
        while (left < right && A[left] < standard)
        {
            left++;
        }
        while (left < right && A[right] > standard)
        {
            right--;
        }
        swap(&A[left], &A[right]);
    }
    if (A[left] > A[end])
    {
        swap(&A[left], &A[end]);
    }
    else
    {
        left++;
    }
    if (left)
    {
        q(A, 0, left - 1);
    }
    q(A, left + 1, right);
}
void quicksort(int *A, int ASize)
{
    q(A, 0, ASize - 1);
}
int smallestRangeI(int *A, int ASize, int K)
{
    int i, max = A[0], min = A[0];
    for (i = 0; i < ASize; i++)
    {
        if (A[i] > max)
            max = A[i];
        if (A[i] < min)
            min = A[i];
    }
    return max - min > 2 * K ? max - min - 2 * K : 0;
}
