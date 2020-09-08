#include <stdio.h>
#include <stdlib.h>
/* https://blog.csdn.net/weixin_43268636/article/details/87274679 */
#define arraysize(a) \
    (sizeof(a)/sizeof*(a))
typedef struct  Edge				//定义边集数组元素，v1,v2存顶点，weight存权重。
{
    int v1;
    int v2;
    int weight;
}Edge;

typedef struct ALGraph		//定义图的结构，peak存顶点的数量，edge存边的数量
{											//指针p作为边集数组，指针m为作为顶点数组
    int peak;
    int edge;
    Edge *p;
    int *m;
}ALGraph;

void CreatALGraph(ALGraph *G)
{
    int i,j;
    printf("输入图的顶点数量和边的数量：");
    scanf("%d %d",&G->peak,&G->edge);
    G->p=(Edge *)malloc(sizeof(Edge)*(G->edge+1));
    G->m=(int *)malloc(sizeof(int)*G->peak);
    for(i=0;i<G->peak;i++)
    {
          printf("请输入输入顶点：");
          scanf("%d",&G->m[i]);
    }
    for(i=0;i<G->edge;i++)
    {
        printf("请输入(vi-vj)和权重：");
        scanf("%d %d %d",&G->p[i].v1,&G->p[i].v2,&G->p[i].weight);
    }
    for(i=0 ;i<G->edge;i++)				//冒泡排序法，权重从小到大存在边集数组中
    {
        for(j=G->edge-1;j>i;j--)
        {
            if(G->p[i].weight>G->p[j].weight)
            {
                G->p[G->edge]=G->p[i];
                G->p[i]=G->p[j];
                G->p[j]=G->p[G->edge];
            }
        }
    }
}
int Find(int *parent,int g)				//通过parent[]找到可连接的边
{
    while(parent[g]!=0)
    {
        g=parent[g];
    }
    return g;
}
int Finish(ALGraph *G,int *parent)		//判断生成树是否完成，完成的标志是生成树的边等于顶点的数量减1
{
    int i,n=0;

    for(i=0;i<G->peak;i++)
    {
        if(parent[i])
        {
            n++;
        }
    }
    if(n==G->peak-1)
    {
        return 1;
    }
    return 0;
}
int FindPeak(ALGraph *G,int g)		//找到顶点的下标
{
    int i;
    for(i=0;i<G->peak;i++)
    {
        if(G->m[i]==g)
            return i;
    }
    return -1;
}
void MinTree_Kruskal(ALGraph *G)
{
    int i,a,b;

    int parent[G->peak];

    for(i=0;i<G->peak;i++)		//初始化parent[]
    {
        parent[i]=0;
    }
    for(i=0;i<G->edge;i++)
    {
        a=Find(parent,FindPeak(G,G->p[i].v1));
        b=Find(parent,FindPeak(G,G->p[i].v2));

        if(a!=b)				//如果a==b则表是a和b在同一颗生成树上，如果a和b连接则为生成环，不符合生成树
        {
            parent[a]=b;
            printf("%d->%d   %d\n",G->p[i].v1,G->p[i].v2,G->p[i].weight);
        }
        if(Finish(G,parent))		//完成后返回
        {
            return;
        }
    }
}

int  main()
{
    ALGraph *G=(ALGraph *)malloc(sizeof(ALGraph));
    CreatALGraph(G);
    MinTree_Kruskal(G);
    return 0;
}

