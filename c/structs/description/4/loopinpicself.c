#include <stdio.h>
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
#define VertexNum  7

int g[VertexNum][VertexNum];
int Visited[VertexNum];


void
printArrary(){
    for(int i=0;i<VertexNum;i++) {
        for(int j=0;j<VertexNum;j++) {
            printf("%d,",g[i][j]);
        }
        printf("\n");
    }
}
void
checkLoop() {
    // init pic
    for(int i=0;i<VertexNum;i++){
        Visited[i]=0;
        for(int j=0;j<VertexNum;j++) {
            g[i][j]=-1;
        }
    }
    for(int i=0;i<VertexNum;i++) {
        for(int j=0;j<VertexNum;j++) {
            int c;
            scanf("%d", &c);
            g[i][j]=c;
        }
    }
    printArrary(g);
    // check
    int single=1, isLoop=1;
    for(int n=0;n<VertexNum;n++) {
        isLoop=1;
        for(int i=0;i<VertexNum;i++) {
            single=1;
            if(!Visited[i]) {
                for(int j=0;j<VertexNum;j++) {
                    if(g[j][i]!=-1) {
                        single = 0;
                    }
                }
                if(single) { // i is single, delete i
                    for(int k=0;k<VertexNum;k++) {
                        g[i][k]=-1;
                    }    
                    Visited[i]=1;
                }
            }
        }
        for(int i=0;i<VertexNum;i++) {
            if(!Visited[i]){
                isLoop=0;
            }
        }
        if(isLoop) {
            break;
        }
    }
    if(isLoop) {
        printf("loop\n");
    }
}

int
main(int argc, char *argv[]){
    checkLoop();
    return 0;
}