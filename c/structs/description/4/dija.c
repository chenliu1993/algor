#include <stdio.h>
#include <stdlib.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
#define X 10000
#define VertexNum  7  //实际上共有六个顶点（1---6）
#define EdgeNum  9

int Graph[VertexNum][VertexNum] =
//0  1  2  3  4  5  6
{ X, X, X, X, X, X, X,  //0
  X, X, 6, 3, X, X, X,  //1
  X, X, X, X, 5, X, X,  //2
  X, X, 2, X, 3, 4, X,  //3
  X, X, X, X, X, X, 3,  //4
  X, X, X, X, 2, X, 5,  //5
  X, X, X, X, X, X, X   //6
};

int Visited[VertexNum];
int Unvisited[VertexNum];
int Distance[VertexNum];

void
PrintArray(int a[]) {
    for(int i=0;i<VertexNum;i++) {
        printf("%d,",a[i]);
    }
    printf("\n");
}

void
Dija(int src) {
    int vcount=1, minlen=X,minidx=-1,i,j;
    i=j=0;
    for(i=0;i<VertexNum;i++) {
        Distance[i] = Graph[src][i];
    }
    Distance[src]=0;
    for(i=0;i<VertexNum;i++) {
        if(i==src){
            Unvisited[i]=0;
            Visited[i]=1;
        }else{
            Unvisited[i]=1;
            Visited[i]=0;
        }
    }
    while(vcount<VertexNum) {
        // iterate through u
        for(i=0;i<VertexNum;i++){
            if(Unvisited[i]==1) {
                PrintArray(Visited);
                // iterate through v
                for(j=0;j<VertexNum;j++) {
                    if(Visited[j]==1) {
                        if(Distance[j]+Graph[j][i]<minlen){
                            minlen=Distance[j]+Graph[j][i];
                            minidx=i;
                        }
                    }
                }
                Visited[minidx]=1;
                Unvisited[minidx]=0;
                Distance[minidx]=minlen;
                vcount++;
            }
        }
    }
}

int
main(int argc, char *argv[]){
    Dija(2);
    return 0;
}
