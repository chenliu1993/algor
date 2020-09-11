// sliding window
#include <stdio.h>

int
max(int a, int b){
    return a>b?a:b;
}
void
printArray(int *a, int start, int end) {
    for(int i=start;i<=end;i++) {
        printf("%d,",a[i]);
    }
    printf("\n");
}

void
swap(int* A, int ASize, int* B, int BSize) {
    int tempS, *temp;
    if(ASize<BSize){
        tempS=ASize;
        ASize=BSize;
        BSize=tempS;
        temp = A;
        A=B;
        B=temp;
    }
}
int
findMax(int* A, int Astart, int Aend, int* B, int Bstart, int Bend) {
    int len,maxlen;
    if((Aend<Astart)||(Bstart>Bend)){return 0;}
    maxlen=len=0;
    for(int i=0;i<=Aend-Astart;i++){
       if(A[Astart+i]==B[Bstart+i]){
           len++;
       }else{
           maxlen=maxlen<len?len:maxlen;
           len=0;
       }
    }
    maxlen=maxlen<len?len:maxlen;
    return maxlen;
}


int
findLength(int* A, int ASize, int* B, int BSize){
    int maxlen=0;
    swap(A,ASize,B,BSize);
    // in
    for(int i=0;i<BSize+1;i++) {
        maxlen=max(maxlen, findMax(A, 0, i-1, B, BSize-i, BSize-1));
    }
    // mid
    for(int i=1;i<ASize-BSize+1;i++) {
        maxlen=max(maxlen, findMax(A, i, i+BSize-1, B, 0, BSize-1));
    }
    // out
    for(int i=ASize-BSize+1,j=1;i<ASize;i++,j++) {
        maxlen=(max(maxlen, findMax(A, i, ASize-1, B, 0, BSize-1-j)));
    }
    return maxlen;
}

int
main(int argc, char *argv[]){
    int a[5]={1,2,3,2,1};
    int b[5]={3,2,1,4,7};
    printf("%d\n",findLength(a,5,b,5));
    return 0;
}