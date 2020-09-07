#include <stdio.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
static int
max(int a, int b) {
    if (a>b) {
        return a;
    }else{
        return b;
    }
}

// TODO: support negatives.
// static int
// MaxSubSum(const int a[], int left, int right) {
//     int center;
//     int maxleftsum, maxrightsum, maxleftborder, leftborder, maxrightborder,rightborder;
//     if (left==right) {
//         if(a[left]>=0)
//         {return a[left];}
//         else
//         {return 0;}
//     }
//     center = (left+right)/2;
//     maxleftsum = MaxSubSum(a, left, center);
//     maxrightsum = MaxSubSum(a, center+1, right);

//     maxleftborder=maxrightborder=0;
//     leftborder=rightborder=0;
//     for(int i=center;i>=left;i--) {
//         leftborder += a[i];
//         if(leftborder > maxleftborder){
//             maxleftborder=leftborder;
//         }
//     }

//     for(int i=center+1;i<=right;i++) {
//         rightborder += a[i];
//         if (maxrightborder < rightborder){
//             maxrightborder = rightborder;
//         }
//     }
//     printf("%d, %d, %d\n",maxrightborder+maxleftborder, maxleftsum, maxrightsum);
//     return max(max(maxrightborder+maxleftborder, maxleftsum), maxrightsum);
// }

static int
MaxSubSum(const int a[], int len) {
    int current, maxsum;
    current = maxsum = 0;
    for(int j =0;j<len;j++) {
        current += a[j];
        if(current > maxsum){
            maxsum=current;
        }else if(current <0){
            current =0;
        }
    }
    return maxsum;
}

int
main(int argc, char *argv[]) {
    const int a[8] = {4,-3,5,-2,-1,2,6,-2};
    // printf("the max subarray size is %d\n", MaxSubSum(a,0,arraysize(a)-1));
    printf("the max subarray size is %d\n", MaxSubSum(a,arraysize(a)));
    return 0;
}