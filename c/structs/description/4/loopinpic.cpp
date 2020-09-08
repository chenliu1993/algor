/* https://www.cnblogs.com/cmai/p/7517729.html */
/**
 * 1.统计各个图中各个点的入度数（能够到达这个点的点的数量）。

　　2.然后找出入度数为0的点（无向图找入度数为1的点）。

　　3.删除入度数为0的点，将其边也删除。

　　4.重复2，直到所有点入度都为0，则为无环图，如果找不到入度为0的点，则为有环图。**/

#include<stdio.h>
using namespace std;
int graph[100][100];//用来存储图的数组
bool isVisited[100];//判断这个点是否已经删除
int main()
{
    int n,e;
    while (scanf("%d",&n)!=EOF&&n!=0)//获取点数
    {
        for(int i = 0;i<100;i++)
        {
            isVisited[i] = 0;
            for(int j = 0 ;j<100;j++)
            {
                graph[i][j] = -1;//初始化数据，所有的边都为-1，代表这两个点之间不能联通
            }
        }
        scanf("%d",&e);//获取边数
        for(int i = 0 ;i<e;i++)//构建图
        {
            int a,b,c;
            scanf("%d %d %d",&a,&b,&c);
            graph[a-1][b-1] = c;
        }
        int isResult = 1;
        for(int i = 0 ;i<n;i++)//进行n次循环，每次循环删除一个入度为0的点，所以进行n次循环
        {
            for(int j = 0;j<n;j++)//遍历所有的点，找入度为0的点
            {
                if(!isVisited[j])//判断该点是否删除
                {
                    bool isCanVisited = true;//辅助变量，判断这个点是否入度为0
                    for(int k = 0;k<n ;k++)
                    {
                        if(graph[k][j]!=-1)
                        {
                            isCanVisited = 0;//如果存在能够访问这个点的边，则该点入度不为0
                        }
                    }
                    if(isCanVisited)//如果该点入度为0，则下边是删除该点和删除其相邻边
                    {
                        for(int k = 0 ;k<n;k++)
                        {
                            graph[j][k] = -1;//删除相邻边，即将值变为-1
                        }
                        isVisited[j] = true;//删除该点
                    }
                }
            }
            isResult = true;
            for(int j = 0 ;j<n;j++)//进行循环判断当前所有点是否已经全部删除，如果全部删除，如果全部删除则跳出，否则继续循环
            {
                if(!isVisited[j])
                {
                    isResult = 0;
                }
            }
            if(isResult)
                break;
        }
        isResult = true;
        for(int i = 0 ;i<n;i++)//在所有点遍历后，则通过这个循环来判断是否所有点都已经删除，如果全部删除，则为无环图，否则为有环图
        {
            if(!isVisited[i])
                isResult = 0;
        }
        if(isResult)
            printf("无环");
        if(!isResult)
            printf("有环");
    }
    return 0;
 
}