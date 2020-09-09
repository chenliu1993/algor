#include <stdio.h>

#define Len 8
int a[Len] = {4,-3,5,-2,-1,2,6,-2};

void swap(int *x, int *y) {
    int t = *x;
    *x = *y;
    *y = t;
}

void
qicksort(int a[], int left, int right) {
    if(left>=right) return;
    int standard=a[right];
    int start=left, end=right-1;
    while(start<end) {
        while(a[start]<standard&&start<end){
            start++;
        }
        while(a[end]>standard&&start<end) {
            end--;
        }
        swap(&a[start], &a[end]);
    }
    if(a[start]>=a[right]) {
        swap(&a[start], &a[right]);
    }else{
        start++;
    }
    if(start){
        qicksort(a,left,start-1);
    }
    qicksort(a, start+1,right);
}

void
qsort() {
    qicksort(a, 0, Len-1);
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
    qsort();
    printArray(a);
}