#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int countSum(int *a, int left, int right)
{
    int sum;
    for (int i = left; i <= right; i++)
    {
        sum += a[i];
    }
    return sum;
}
int sumSubarrayMins(int *A, int ASize)
{
    int rightPtr, count;
    rightPtr = count = 0;
    while (rightPtr < ASize)
    {
        for (int i = 0; i <= rightPtr; i++)
        {
            count += countSum(A, i, rightPtr);
        }
        rightPtr++;
    }
    for (int i = 0; i <= rightPtr; i++)
    {
        count += countSum(A, i, rightPtr);
    }
    return count % (10 ^ 9 + 7);
}
int main(int argc, char *argv[])
{
    int a[4] = {3, 1, 2, 4};
    printf("%d\n", sumSubarrayMins(a, 4));
    return 0;
}

// class Solution {
// private:
//     const int mod = 1e9 + 7;
// public:
//     int sumSubarrayMins(vector<int>& A) {
//         stack<int> inc_stack; //单调递增栈, 存储的是数组A的索引
//         A.push_back(-1); //保证栈里所有元素都弹出
//         int len = A.size();
//         // {__n__    A[i]    ___m____ };    如 8, 4, 6, 5, 7, 9, 3, 0 -> 对于A[i] = 5, n = 1 ([6]), m = 2 ([7, 9])
//         // A[i]左侧有n个连续的大于A[i]的数, 右侧有连续m个大于A[i]的数, 则A[i]作为最小值的数组个数为(m+1)*(n+1)
//         // 关键即为找到每个A[i]对应的区间, 满足A[i]为该区间内的最小值
//         int ans = 0;
//         for(int i = 0; i < len; i++){
//             while(!inc_stack.empty() && A[inc_stack.top()] >= A[i]){ // A[i]比栈顶元素小, 说明栈顶元素的区间截止, 该处理栈顶元素
//                 int index = inc_stack.top(); //取栈顶索引
//                 inc_stack.pop();
//                 int pre = index - (inc_stack.empty()? -1 : inc_stack.top()); //对应n+1
//                 int after = i - index;   //对应m+1
//                 ans += A[index]*pre*after;   //A[i]作为最小值的次数为(m+1)*(n+1), 即ans加上A[i]*(m+1)*(n+1)
//                 ans %= mod;
//             }
//             inc_stack.push(i);
//         }
//         return ans;
//     }
// };
