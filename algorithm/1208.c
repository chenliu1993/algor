#include <stdio.h>
/* *
 * refer to 1004, sliding window.
 * */
int abs(char a, char b)
{
    return (a - b) < 0 ? (b - a) : (a - b);
}
int equalSubstring(char *s, char *t, int maxCost)
{
    int r, l, sum, count;
    r = l = sum = count = 0;
    while (s[r] != NULL)
    {
        if (sum <= maxCost)
        {
            sum += abs(s[r], t[r]);
            r++;
        }
        else
        {
            sum -= abs(s[l], t[l]);
            l++;
        }
        while (sum <= maxCost && r - l > count)
        {
            count = r - l;
        }
    }
    return count;
}
int main(int argc, char *argv[])
{
    char *a = "abcd";
    char *b = "cdef";
    int cost = 3;
    printf("%d\n", equalSubstring(a, b, cost));
}