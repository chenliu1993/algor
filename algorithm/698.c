#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
int dfs(int *nums, int numsSize, int target, int needfind, int start, int cur, int *used)
{
    if (needfind == 0)
    {
        return 1;
    }
    if (start >= numsSize)
    {
        return 0;
    }
    int rc = 0;
    for (int i = start; i < numsSize; i++)
    {
        if (used[i] != 0)
        {
            continue;
        }
        used[i] = 1;
        if (cur + nums[i] == target)
        {
            rc = dfs(nums, numsSize, target, needfind - 1, 0, 0, used);
        }
        else if (cur + nums[i] < target)
        {
            rc = dfs(nums, numsSize, target, needfind, i + 1, cur + nums[i], used);
        }
        if (rc != 1)
        {
            used[i] = 0;
        }
        else
        {
            return 1;
        }
    }
    return 0;
}
int canPartitionKSubsets(int *nums, int numsSize, int k)
{
    int sum = 0, maxnum = INT_MIN;
    for (int i = 0; i < numsSize; i++)
    {
        sum += nums[i];
        maxnum = maxnum < nums[i] ? nums[i] : maxnum;
    }
    if (sum % k != 0)
    {
        return 0;
    }
    int target = sum / k;
    if (maxnum > target)
    {
        return 0;
    }
    int *used = (int *)calloc(17, sizeof(int));
    int rc = dfs(nums, numsSize, target, k, 0, 0, used);
    return rc;
}

int main(int argc, char *argv[])
{
    int a[] = {4, 3, 2, 3, 5, 2, 1};
    int k = 4;
    printf("%d\n", canPartitionKSubsets(a, 5, k));
    return 0;
}