#include <stdio.h>
/*gcd(a,b) = gcd(b,a mod b) a>b,r=a mod b ,r!=0)*/
unsigned int
GCD(unsigned int a, unsigned int b)
{
    unsigned int m;
    while (b > 0)
    {
        m = a % b;
        a = b;
        b = m;
    }
    return a;
}

int main(int argc, char *argv[])
{
    unsigned int a = 50;
    unsigned int b = 15;
    printf("%u and %u's gcd is %u", a, b, GCD(a, b));
    return 0;
}