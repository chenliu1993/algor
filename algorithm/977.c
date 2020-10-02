#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void q(int *arr, int start, int end)
{
    if (start <= end)
        return;
    int left = start, right = end - 1, standard = arr[end];
    while (left < right)
    {
        while (left < right && arr[left] < standard)
        {
            left++;
        }
        while (left < right && arr[right] > standard)
        {
            right--;
        }
        swap(&arr[left], &arr[right]);
    }
    if (arr[left] > arr[end])
    {
        swap(&arr[left], &arr[end]);
    }
    else
    {
        left++;
    }
    if (left)
    {
        q(arr, 0, left - 1);
    }
    q(arr, left + 1, end);
}
void qickSort(int *arr, int arrSize)
{
    q(arr, 0, arrSize - 1);
}
void insertSort(int *arr, int arrSize)
{
    int m, j;
    for (int i = 1; i < arrSize; i++)
    {
        m = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > m)
        {
            arr[j + 1] = arr[j];
            j--;
        }
        arr[j + 1] = m;
    }
}
int *sortedSquares(int *A, int ASize, int *returnSize)
{
    for (int i = 0; i < ASize; i++)
    {
        A[i] = A[i] * A[i];
    }
    insertSort(A, ASize);
    *returnSize = ASize;
    return A;
}
void printArray(int *arr, int arrSize)
{
    for (int j = 0; j < arrSize; j++)
    {
        printf("%d,", arr[j]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int a[5] = {-4, -1, 0, 3, 10};
    int *returnSize = (int *)calloc(1, sizeof(int));
    int *b = sortedSquares(a, 5, returnSize);
    printArray(b, 5);
    return 0;
}