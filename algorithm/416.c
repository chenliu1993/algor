#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int compareFunc(const void *a, const void *b)
{

    return (*(int *)a) > (*(int *)b);
}
void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}
void q(int *nums, int start, int end)
{
    if (start >= end)
    {
        return;
    }
    int left = start, right = end - 1, standard = nums[end];
    while (left < right)
    {
        while (left < right && nums[left] < nums[standard])
        {
            left++;
        }
        while (left < right && nums[right] > nums[standard])
        {
            right--;
        }
        swap(&nums[left], &nums[right]);
    }
    if (nums[left] > nums[end])
    {
        swap(&nums[left], &nums[end]);
    }
    else
    {
        left++;
    }
    if (left)
    {
        q(nums, 0, left - 1);
    }
    q(nums, left + 1, end);
}
void qicksort(int *nums, int numsSize)
{
    q(nums, 0, numsSize - 1);
}
// see from official answer.
int canPartition(int *nums, int numsSize)
{
    if (numsSize < 2)
    {
        return 0;
    }
    // qsort(nums, numsSize, sizeof(int), compareFunc);
    qicksort(nums, numsSize);
    int sum = 0, target;
    for (int i = 0; i < numsSize; i++)
    {
        sum += nums[i];
    }
    if (sum % 2 != 0)
    {
        return 0;
    }
    target = sum / 2;
    if (nums[numsSize - 1] > target)
    {
        return 0;
    }
    int **dp = (int **)malloc(numsSize * sizeof(int *));
    for (int i = 0; i < numsSize; i++)
    {
        dp[i] = (int *)malloc((target + 1) * sizeof(int));
        memset(dp[i], 0, (target + 1) * sizeof(int));
    }
    // Border condition
    for (int i = 0; i < numsSize; i++)
    {
        dp[i][0] = 1;
    }
    dp[0][nums[0]] = 1;
    // Dynamic Planning part
    for (int i = 1; i < numsSize; i++)
    {
        for (int j = 1; j < target + 1; j++)
        {
            if (nums[i] > j)
            {
                // nums[i] must not be selected
                dp[i][j] = dp[i - 1][j];
            }
            else
            {
                // dp[i][j] == nums[i] not selected || nums[i] selected
                dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i]];
            }
        }
    }
    return dp[numsSize - 1][target];
}

int main(int argc, char *argv[])
{
    int a[4] = {1, 5, 5, 11};
    printf("%d\n", canPartition(a, 4));
    return 0;
}