#include <stdio.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))

void
InsertSort(int a[]) {
    int tmp, i, j;
    for(i=1;i<arraysize(a);i++) {
        tmp = a[i];
        j=i-1;
        while(j>=0&&a[j]>tmp) {
            a[j+1]=a[j];
            j--;
        }
        a[j+1]=tmp;
    }
}

void
ShellSort(int a[]) {
    int i,j, incre, tmp;
    for(incre=arraysize(a)/2;incre>0;incre/=2){
        for(i=incre;i<arraysize(a);i++){
            tmp=a[i];
            for(j=i;j>=incre;j -= incre){
                if(tmp<a[j-incre]){
                    a[j]=a[j-incre];
                }else{
                    break;
                }
            }
            a[j]=tmp;
        }
    }
}
int
main(int argc, char *argv[]) {
    int a[8] = {4,-3,5,-2,-1,2,6,-2};
    for(int i=0;i<arraysize(a);i++){
        printf("%d,", a[i]);
    }
    printf("\n");
    ShellSort(a);
    for(int i=0;i<arraysize(a);i++){
        printf("%d,", a[i]);
    }
    printf("\n");
}

