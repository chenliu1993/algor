#include <stdio.h>

long int
Pow(long int a, unsigned int n) {
    if (n==0){
        return 1;
    }
    if(n==1){
        return a;
    }
    if(n%2==0){
        return Pow(a*a, n/2);
    }else{
        return Pow(a*a, n/2)*a;
    }
}

int
main(int argc, char *argv[]) {
    printf("the res is %ld\n", Pow(2, 10));
    return 0;
}