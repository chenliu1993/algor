#include <stdio.h>
#include <stdlib.h>

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
int path[VertexNum];
int Distance[VertexNum];

void Dijkstra(int Begin)
{
  int MinEdge, Vertex, i,j, Edges;
  Edges = 1;
  Visited[Begin] = 1;
  for (i = 1; i<VertexNum; i++) Distance[i] = Graph[Begin][i];

  Distance[Begin] = 0;
  printf("     1  2  3  4  5  6\\n");
  printf("-----------------------------------\\n");
  printf("s:%d", Edges);
  for( i=1; i<VertexNum; i++)
  if (Distance[i] == X) printf("  *"); else printf("%3d",Distance[i]);
  printf("\\n");
  while( Edges<VertexNum-1)
  {
    Edges++; MinEdge = X;
    for(j=1; j<VertexNum; j++)
    if (Visited[j]==0 && MinEdge > Distance[j] )
    {
        Vertex = j; MinEdge = Distance[j];
    }
    Visited[Vertex] = 1;
    printf("s:%d",Edges);
    for(j=1; j<VertexNum; j++)
    {
      if (Visited[j] == 0 && Distance[Vertex] + Graph[Vertex][j] <Distance[j])
      {   Distance[j] = Distance[Vertex] + Graph[Vertex][j];
        path[j] = Vertex;
      }
      //printf("%6d",Distance[j]);
       if (Distance[j] == X) printf("  *"); else printf("%3d",Distance[j]);
    }
    printf("\\n");
  }
}

void main()
{
  
  int i;
  int k;
 // clrscr();
  for(i=0; i<VertexNum; i++) { Visited[i] = 0;  path[i] = 1;}
  Dijkstra(1);
  printf("\\n\\nAll Path-------------------------\\n");


  for(i=2; i<VertexNum; i++) //printf("%5d",Visited[i]);
  {
     printf("[%d] ",Distance[i]);
     k = i;
     do
     {
       printf("%d<--",k);
       k  = path[k];
     } while (k!=1);
     printf("1 \\n");
  }
}