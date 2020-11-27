#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>
// https : //leetcode-cn.com/problems/jian-sheng-zi-lcof/solution/ji-suan-fa-by-jason-2/
int cuttingRope(int n)
{
    if (n <= 3)
        return n - 1;
    int a = n / 3 - 1, b = n % 3, p = 1000000007;
    long long int sum = 1;
    while (a > 0)
    {
        sum = (sum * 3) % p;
        a--;
    }
    if (b == 0)
        sum = (sum * 3) % p;
    else if (b == 1)
        sum = (sum * 4) % p;
    else
        sum = (sum * 6) % p;
    return sum;
}

// 作者：enjoy-13
// 链接：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof/solution/cyu-yan-shuang-bai-by-enjoy-13-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
int main(int argc, char *argv[])
{
    printf("%d\n", cuttingRope(120));
    return 0;
}