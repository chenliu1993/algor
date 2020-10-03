#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// typedef struct _position
// {
//     int x;
// } position;

// position newPosition(int x)
// {
//     position *p = (position *)calloc(1, sizeof(position));
//     p->x = x;
//     return *p;
// }

// int *twoSum(int *nums, int numsSize, int target, int *returnSize)
// {
//     if (nums == NULL || numsSize < 2)
//     {
//         *returnSize = 0;
//         return NULL;
//     }
//     int *rc = (int *)calloc(2, sizeof(int));
//     position *stk = (position *)calloc(numsSize, sizeof(position));
//     int front, rear;
//     front = rear = 0;
//     stk[rear++] = newPosition(0);
//     while (front < rear)
//     {
//         int a = stk[front].x;
//         for (int j = a + 1; j < numsSize; j++)
//         {
//             if ((nums[a] + nums[j]) == target)
//             {
//                 printf("%d, %d\n", a, j);
//                 rc[0] = a;
//                 rc[1] = j;
//                 *returnSize = 2;
//                 return rc;
//             }
//             else
//             {
//                 stk[rear++] = newPosition(j);
//             }
//         }
//         front++;
//     }
//     *returnSize = 0;
//     return NULL;
// }

int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    int *rc = (int *)calloc(2, sizeof(int));
    for (int i = 0; i < numsSize; i++)
    {
        for (int j = i + 1; j < numsSize; j++)
        {
            if (nums[i] + nums[j] == target)
            {
                rc[0] = i;
                rc[1] = j;
                *returnSize = 2;
                return rc;
            }
        }
    }
    *returnSize = 0;
    return NULL;
}

void printArray(int *a)
{
    for (int i = 0; i < 2; i++)
    {
        printf("%d,", a[i]);
    }
    printf("\n");
}
int main(int argc, char *argv[])
{
    int a[4] = {2, 7, 11, 5};
    int tgt = 9;
    int *returnSize = (int *)calloc(1, sizeof(int));
    printArray(twoSum(a, 4, tgt, returnSize));
    return 0;
}