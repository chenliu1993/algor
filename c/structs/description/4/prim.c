#include <stdio.h>

/* https://www.cnblogs.com/skywang12345/p/3711506.html */
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

void
prim() {
    int minw = X, minidx=0, vcount=1;
    for(int i=0;i<VertexNum;i++) {
        if(i==0){
            Visited[i]=1;
            Unvisited[i]=0;
        }else{
            Visited[i]=0;
            Unvisited[i]=1;
        }
    }
    while(vcount<VertexNum) {
        for(int i=0;i<VertexNum;i++) {
            if(Visited[i]==1) {
                for(int j=0;j<VertexNum;j++) {
                    if(Unvisited[j]==1) {
                        if(minw>Graph[i][j]){
                            minw= Graph[i][j];
                            minidx=j;
                        }
                    }
                }
            }
        }
        Visited[minidx]=1;
        Unvisited[minidx]=0;
        vcount++;
    }
}