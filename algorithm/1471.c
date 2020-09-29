#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int midval;
void printArray(int *arr, int arrSize)
{
    for (int i = 0; i < arrSize; i++)
    {
        printf("%d,", arr[i]);
    }
    printf("\n");
}
int compareStrongest(int a, int b, int m)
{
    int first = abs(a - m) - abs(b - m);
    if (first > 0 || (first == 0 && a > b))
    {
        return 1;
    }
    return 0;
}

// void strongestSort(int *arr, int arrSize, int mid)
// {
//     int val, j;
//     for (int i = 1; i < arrSize; i++)
//     {
//         val = arr[i];
//         j = i - 1;
//         while (j >= 0 && compareStrongest(val, arr[j], mid))
//         {
//             arr[j + 1] = arr[j];
//             j--;
//         }
//         arr[j + 1] = val;
//     }
// }
void swap(int *x, int *y)
{
    int t = *x;
    *x = *y;
    *y = t;
}
// quick sort
void strongestQSort(int *arr, int left, int right, int mid)
{
    if (left >= right)
    {
        return;
    }
    int start = left, end = right - 1;
    int stand = arr[right];
    while (start < end)
    {
        while (compareStrongest(arr[start], stand, mid) && start < end)
        {
            start++;
        }
        while (!compareStrongest(arr[end], stand, mid) && start < end)
        {
            end--;
        }
        swap(&arr[start], &arr[end]);
    }
    if (compareStrongest(arr[right], arr[start], mid))
    {
        swap(&arr[start], &arr[right]);
    }
    else
    {
        start++;
    }
    if (start)
    {
        strongestQSort(arr, left, start - 1, mid);
    }
    strongestQSort(arr, start + 1, right, mid);
}
void strongestSort(int *arr, int arrSize, int mid)
{
    strongestQSort(arr, 0, arrSize - 1, mid);
}

void q(int *arr, int left, int right)
{
    if (left >= right)
    {
        return;
    }
    int start = left, end = right - 1, standard = arr[right];
    while (start < end)
    {
        while (arr[start] < standard && start < end)
        {
            start++;
        }
        while (arr[end] > standard && start < end)
        {
            end--;
        }
        swap(&arr[start], &arr[end]);
    }
    if (arr[start] >= arr[right])
    {
        swap(&arr[start], &arr[right]);
    }
    else
    {
        start++;
    }
    if (start)
    {
        q(arr, left, start - 1);
    }
    q(arr, start + 1, right);
}
void quickSort(int *arr, int arrSize)
{
    q(arr, 0, arrSize - 1);
}
void insertSort(int *arr, int arrSize)
{
    int val, j;
    for (int i = 1; i < arrSize; i++)
    {
        val = arr[i];
        j = i - 1;
        while (j >= 0 && val < arr[j])
        {
            arr[j + 1] = arr[j];

            j--;
        }
        arr[j + 1] = val;
    }
}
/**
 * |arr[i] - m| > |arr[j] - m|
 * |arr[i] - m| == |arr[j] - m|，且 arr[i] > arr[j]
 */

int compareFunc1(const void *a, const void *b)
{
    return (*(int *)a - *(int *)b);
}
int compareFunc2(const void *a, const void *b)
{
    int *p = (int *)a;
    int *q = (int *)b;
    if (abs(*p - midval) > abs(*q - midval))
    {
        return 0;
    }
    else if (abs(*p - midval) == abs(*q - midval) && *p > *q)
    {
        return 0;
    }
    return 1;
}
int *getStrongest(int *arr, int arrSize, int k, int *returnSize)
{
    if (arr == NULL)
    {
        return arr;
    }
    // from stdlib
    qsort(arr, arrSize, sizeof(int), compareFunc1);
    printArray(arr, arrSize);
    midval = arr[(arrSize - 1) / 2];
    printf("mid val is %d\n", midval);
    int *rc = (int *)calloc(k, sizeof(int));
    // strongestSort(arr, arrSize, midval);
    qsort(arr, arrSize, sizeof(int), compareFunc2);
    printArray(arr, arrSize);

    for (int i = 0; i < k; i++)
    {
        rc[i] = arr[i];
    }
    *returnSize = k;
    return rc;
}

int main(int argc, char *argv[])
{
    int *returnSize = (int *)calloc(1, sizeof(int));

    // int *b = (int *)calloc(100000, sizeof(int));
    int a[9] = {499, 456, 938, 734, 734, -688, -521, -985, 601};
    int k = 6;
    int *rc = getStrongest(a, 9, k, returnSize);
    if (rc == NULL)
    {
        printf("NULL\n");
    }
    printArray(rc, k);
    return 0;
}