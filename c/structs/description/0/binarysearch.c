#include <stdio.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
int
BinarySearch(const int a[], int len, int tgt) {
    int left, right, mid;
    left = 0;
    right = len-1;
    while(left<=right){
        mid = (left+right)/2;
        if(tgt<a[mid]){
            right=mid-1;
        }else if(tgt>a[mid]){
            left  = mid+1;
        }else{
            return mid;
        }
    }
    return -1;
}

int
main(int argc, char *argv[]) {
    const int a[8] = {4,-3,5,-2,-1,2,6,-2};
    int X=-2;
    printf("the target index is %d\n", BinarySearch(a, arraysize(a), X));
    return 0;
}