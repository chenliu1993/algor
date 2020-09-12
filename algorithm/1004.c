#include <stdio.h>
#include <stdlib.h>
/**
 * Below are backtrace to solving 1004, but when ASize is big, then this will cost too much time.
 **/
// void getLongest(int *A, int end, int ASize, int *k, int *count, int *result)
// {
//     int *kTemp = (int *)malloc(sizeof(int));
//     int *countTemp = (int *)malloc(sizeof(int));
//     if (end >= ASize)
//     {
//         *result = *result < *count ? *count : *result;
//         return;
//     }
//     if (A[end])
//     {
//         *count = (*count) + 1;
//         getLongest(A, end + 1, ASize, k, count, result);
//     }
//     else
//     {
//         if (!(*k))
//         {
//             *result = *result < *count ? *count : *result;
//             *count = 0;
//             getLongest(A, end + 1, ASize, k, count, result);
//         }
//         else
//         {
//             // use k.
//             *kTemp = *k;
//             *countTemp = 0;
//             *k = (*k) - 1;
//             A[end] = 1;
//             getLongest(A, end, ASize, k, count, result);
//             A[end] = 0;
//             // don't use k;
//             getLongest(A, end + 1, ASize, kTemp, countTemp, result);
//             free(kTemp);
//             free(countTemp);
//         }
//     }
// }
// int longestOnes(int *A, int ASize, int K)
// {
//     int maxlen;
//     int *k = &K;
//     int *count = (int *)malloc(sizeof(int));
//     int *result = (int *)malloc(sizeof(int));
//     *count = 0;
//     *result = 0;
//     getLongest(A, 0, ASize, k, count, result);
//     maxlen = *result;
//     free(count);
//     free(result);
//     return maxlen;
// }

/* * 
 * sliding window
 * 
 典型的滑动窗口题目：
1.定义两个指针一前r一后l，中间就是最大的子串；
2.当前花费小于等于最大花费时，就让r往前走(包括等于是因为有些位置花费为0)
3.当前花费大于最大花费时，就让l往前走，取r与l的最大差值就是最大长度。

那么如何理解滑动窗口的正确性呢：
对于l在每一个位置：
如果当前子串小于最大花费，r就会挨个右移遍历，r最远可以到达的位置，因此对r为起点的子串遍历是充分的
如果当前子串大于最大花费，那么就证明以l为起点的子串长度小于等于以l-1为起点的最长子串。
那么就已经证明了经过滑动窗口的处理一定可以找到最长的子串。

作者：codefarmer-5
链接：https://leetcode-cn.com/problems/get-equal-substrings-within-budget/solution/cyu-yan-hua-dong-chuang-kou-fa-nei-han-hua-dong-ch/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

作者：codefarmer-5
链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii/solution/cyu-yan-hua-dong-chuang-kou-fa-nei-han-hua-dong-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 * */
int longestOnes(int *A, int ASize, int K)
{
    int r, l, sum, count;
    r = 0;
    l = 0;
    sum = 0;
    count = 0;
    while (r < ASize)
    {
        if (sum <= K)
        {
            sum += (!A[r]);
            r++;
        }
        else
        {
            sum -= (!A[l]);
            l++;
        }
        if (sum <= K && r - l > count)
        {
            count = r - l;
        }
    }
    return count;
}
int main(int argc, char *argv[])
{
    // int a[14] = {0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1};
    // int K = 3;
    // int a[11] = {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0};
    // int K = 2;
    int a[50] = {1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1};
    int K = 10;
    printf("the longesOnes is %d\n", longestOnes(a, 50, K));
    return 0;
}