#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int singleNumber(int *nums, int numsSize)
{
    int tgt = 0;
    for (int i = 0; i < numsSize; i++)
    {
        tgt ^= nums[i];
    }
    return tgt;
}
int main(int argc, char *argv[])
{
    int a[3] = {2, 2, 1};
    printf("%d\n", singleNumber(a, 3));
    return 0;
}