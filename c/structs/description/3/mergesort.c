#include <stdio.h>
#include <stdlib.h>

// #define arraysize(a) \
//     (sizeof(a)/sizeof*(a))
// why this arraysize is not working??

#define Len 8
int a[Len] = {4,-3,5,-2,-1,2,6,-2};

void
mergesort(){
    sort(a, 0, Len-1);
}

void
sort(int a[], int left, int right) {
    if(left==right) {
        return;
    }
    int mid = (left+right)/2;
    sort(a, left, mid);
    sort(a, mid+1,right);
    merge(a, left, mid, right);
}

void
merge(int a[], int left, int mid, int right) {
    int temp[right-left+1];
    int p1=left, p2=mid+1, i=0;;
    while(p1<=mid&&p2<=right){
        temp[i++]=a[p1]<a[p2]?a[p1++]:a[p2++];
    }
    while(p1<=mid) {
        temp[i++]=a[p1++];
    }
    while(p2<=right) {
        temp[i++]=a[p2++];
    }
    for(int j=0;j<right-left+1;j++) {
        a[left+j]=temp[j];
    }
}

void
printArray(int a[]) {
    for(int i=0;i<Len;i++) {
        printf("%d,", a[i]);
    }
    printf("\n");
}

int
main(int argc, char *argv[]) {
    printArray(a);
    mergesort();
    printArray(a);
}

