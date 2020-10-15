#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void reverseString(char *s, int sSize)
{
    int left, right;
    left = 0;
    right = sSize - 1;
    while (left < right)
    {
        char temp = s[left];
        s[left] = s[right];
        s[right] = temp;
        left++;
        right--;
    }
}
int main(int argc, char *argv[])
{
    char s[7] = "Hannah\0";
    printf("%s\n", s);
    reverseString(s, 6);
    printf("%s\n", s);
    return 0;
}