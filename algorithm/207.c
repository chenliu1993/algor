#include <stdio.h>
#include <stdlib.h>

void
printArray(int **a, int size) {
    for(int i=0;i<size;i++) {
        printf("%d,%d\n",a[i][0], a[i][1]);
    }
}

int
canFinish(int numCourses, int** prerequisites, int prerequisitesSize){
    int x,y,sum;
    // create.
    int** value = (int **)malloc(numCourses*sizeof(int*));
    for(int i=0;i<numCourses;i++){
        value[i]=(int*)malloc(sizeof(int)*numCourses);
        memset(value[i], 0 ,numCourses*sizeof(int));
    }
    int* distance = (int*) malloc(numCourses*sizeof(int));
    int* visited = (int*) malloc(numCourses*sizeof(int));
    memset(distance, 0, sizeof(int)*numCourses);
    memset(visited, 0, sizeof(int)*numCourses);
    for(int i=0;i<prerequisitesSize;i++){
        x=prerequisites[i][0];
        y=prerequisites[i][1];
        value[x][y]=1;
    }
    for(int i=0;i<numCourses;i++){
        sum=0;
        for(int j=0;j<numCourses;j++) {
            sum+=value[j][i];
        }
        distance[i]=sum;
    }
    // checkloop
    for(int i=0;i<numCourses;i++) {
        for(int j=0;j<numCourses;j++){
            if(!visited[j]){
                if(distance[j]==0){
                    visited[j]=1;
                    for(int k=0;k<numCourses;k++){
                        if(!visited[k]){
                            if(value[j][k]){
                                distance[k]-=1;
                                value[j][k]=0;
                            }
                        }
                    }
                }
            }
        }
    }
    int isloop=0;
    for(int i=0;i<numCourses;i++) {
        if(distance[i]!=0){
            isloop=1;
        }
    }
    return isloop;
}

int
main(int argc, char *argv[]){
    int num=2;
    int size=1;
    int **pre=(int**)malloc(size*sizeof(int*));
    int *p1=(int*)malloc(2*sizeof(int));
    p1[0]=1;
    p1[1]=0;
    // int *p2=(int*)malloc(2*sizeof(int));
    // p2[0]=0;
    // p2[1]=1;
    pre[0]=p1;
    // pre[1]=p2;

    printArray(pre, size);
    printf("%d\n", canFinish(num, pre, size));
}