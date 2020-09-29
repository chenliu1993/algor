#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define MAX 26
typedef struct _node
{
    char key;
    int count;
} node;

int compareStruct(const void *a, const void *b)
{
    const node *node1 = (const node *)a;
    const node *node2 = (const node *)b;
    if (node1->count == node2->count)
    {
        return node1->key - node2->key;
    }
    else
    {
        return node2->count - node1->count;
    }
}

int leastInterval(char *tasks, int tasksSize, int n)
{
    node chunkList[MAX] = {0};
    int result = 0;
    int taskSizeNow = tasksSize;

    if (tasks == NULL)
    {
        return 0;
    }

    for (int i = 0; i < MAX; i++)
    {
        chunkList[i].key = 'A' + i;
    }
    for (int i = 0; i < tasksSize; i++)
    {
        chunkList[tasks[i] - 'A'].count++;
    }
    while (1)
    {
        qsort(chunkList, MAX, sizeof(node), compareStruct);
        if (chunkList[0].count == 0)
        {
            break;
        }

        if (n > MAX)
        {
            for (int i = 0; i < MAX; i++)
            {
                if (chunkList[i].count != 0)
                {
                    chunkList[i].count--;
                    result++;
                    taskSizeNow--;

                    if (taskSizeNow == 0)
                    {
                        break;
                    }
                }
                else
                {
                    result += (n + 1 - MAX);
                    break;
                }
            }
        }
        else
        {
            for (int i = 0; i < n + 1; i++)
            {
                if (chunkList[i].count != 0)
                {
                    chunkList[i].count--;
                    result++;
                    taskSizeNow--;
                    if (taskSizeNow == 0)
                    {
                        break;
                    }
                }
                else
                {
                    result += (n + 1 - i);
                    break;
                }
            }
        }
    }
    return result;
}

//
// int leastInterval(char* tasks, int tasksSize, int n)
// {
//     /* 依次执行，直接返回 */
//     if (n == 0) {
//         return tasksSize;
//     }

//     int flag[26] = {0};
//     int max = 0;
//     /* 计算每个字母出现次数 */
//     for (int i = 0; i < tasksSize; i++) {
//         flag[tasks[i] - 'A']++;
//         if (flag[tasks[i] - 'A'] > max) {
//             /* 记录最大次数 */
//             max = flag[tasks[i] - 'A'];
//         }
//     }

//     /* 计算出现最大次数的字母个数 */
//     int cnt = 0;
//     for (int i = 0; i < 26; i++) {
//         if (flag[i] == max) {
//             cnt++;
//         }
//     }

//     /* 出现最大次数字母执行完的最短时间 */
//     int t = (max - 1) * (n + 1) + cnt;
//     /* 时间不能小于字母个数 */
//     if (t < tasksSize) {
//         t = tasksSize;
//     }
//     return t;
// }

int main(int argc, char *argv[])
{
    char a[6] = {'A', 'A', 'A', 'B', 'B', 'B'};
    int n = 2;
    printf("%d\n", leastInterval(a, 6, n));
    return 0;
}