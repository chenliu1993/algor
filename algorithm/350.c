#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int compareFunc(const void *a, const void *b)
{
    return *(int *)a > *(int *)b;
}
int *intersection(int *nums1, int nums1Size, int *nums2, int nums2Size, int *returnSize)
{
    int *rc = (int *)calloc(nums2Size < nums1Size ? nums2Size : nums1Size, sizeof(int));
    int rear = 0;
    qsort(nums1, nums1Size, sizeof(int), compareFunc);
    qsort(nums2, nums2Size, sizeof(int), compareFunc);
    for (int i = 0, j = 0; i < nums1Size && j < nums2Size;)
    {
        if (nums1[i] < nums2[j])
        {
            i++;
        }
        else if (nums1[i] > nums2[j])
        {
            j++;
        }
        else
        {
            rc[rear++] = nums1[i];
            i++;
            j++;
            // if (rear > 1 && rc[rear - 1] == rc[rear - 2])
            // {
            //     rear--;
            // }
        }
    }
    rc = (int *)realloc(rc, rear * sizeof(int));
    *returnSize = rear;
    return rc;
}
void printArray(int *a, int asize)
{
    for (int i = 0; i < asize; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int a[4] = {1, 2, 2, 1};
    int b[2] = {2, 2};
    int *returnSize = (int *)calloc(1, sizeof(int));
    int *c = intersection(a, 4, b, 2, returnSize);
    printArray(c, *returnSize);
    return 0;
}