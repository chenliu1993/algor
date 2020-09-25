#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>

int canChange(char *a, char *b, int len)
{
    int diffCount = 0;
    for (int i = 0; i < len; i++)
    {
        if (a[i] != b[i])
        {
            diffCount++;
        }
    }
    if (diffCount <= 1)
    {
        return 1;
    }
    return 0;
}

/** BFS
 * while(queue empty?){
 *  for(){
 *      queue out
 *      queue in
 *  }
 *  update conditions.
 * }
 **/
// int ladderLengthBFS(char *beginWord, char *endWord, char **wordList, int wordListSize)
// {
//     int inList = 0;
//     for (int i = 0; i < wordListSize; i++)
//     {
//         if (!strcmp(wordList[i], endWord))
//         {
//             inList = 1;
//             break;
//         }
//     }
//     if (!inList)
//     {
//         return 0;
//     }

//     int *visited = (int *)malloc(sizeof(int) * wordListSize);
//     memset(visited, 0, sizeof(int) * wordListSize);
//     int slen = strlen(beginWord);
//     int **queue = (int **)malloc(sizeof(int *) * wordListSize);
//     for (int i = 0; i < wordListSize; i++)
//     {
//         queue[i] = (int *)malloc(sizeof(int) * slen);
//         memset(queue[i], 0, sizeof(int) * slen);
//     }
//     int front, rear, rangeCount, minlen;
//     front = rear = rangeCount = 0;
//     minlen = 1;
//     queue[rear++] = beginWord;
//     int levelCount = rear - front;
//     while (front != rear)
//     {
//         for (int i = 0; i < levelCount; i++)
//         {
//             if (strcmp(queue[front], endWord) == 0)
//             {
//                 return minlen;
//             }
//             else
//             {
//                 for (int k = 0; k < wordListSize; k++)
//                 {
//                     if (!visited[k] && canChange(queue[front], wordList[k], slen))
//                     {
//                         visited[k] = 1;
//                         queue[rear++] = wordList[k];
//                         rangeCount++;
//                     }
//                 }
//             }
//             front++;
//         }
//         levelCount = rangeCount;
//         rangeCount = 0;
//         minlen++;
//     }
//     return minlen;
// }

/** DFS
 * if(condistion meet){
 *  return
 * }
 * for(){
 *  backtrace()
 * }
 * retrun
 **/
int ladderLength(char *beginWord, char *endWord, char **wordList, int wordListSize)
{
    // if wordList already contains the tgt word.
    int inList = 0;
    for (int i = 0; i < wordListSize; i++)
    {
        if (!strcmp(endWord, wordList[i]))
        {
            inList = 1;
            break;
        }
    }
    if (!inList)
    {
        return 0;
    }
    int slen = strlen(beginWord);
    int len = 1, minlen = INT_MAX;
    int *visited = (int *)malloc(sizeof(int) * wordListSize);
    memset(visited, 0, sizeof(int) * wordListSize);
    char **stack = (char **)malloc(sizeof(char *) * wordListSize);
    for (int i = 0; i < wordListSize; i++)
    {
        stack[i] = (char *)malloc(sizeof(char) * slen);
        memset(stack[i], 0, sizeof(char) * slen);
    }
    stack[0] = beginWord;
    ladderLengthDFS(endWord, visited, stack, 1, &len, &minlen, wordList, wordListSize, slen);
    if (minlen == INT_MAX)
    {
        minlen = 0;
    }
    return minlen;
}

void ladderLengthDFS(char *endWord, int *visited, char **stack, int stackSize, int *len, int *minlen, char **wordList, int wordListSize, int slen)
{
    int oldLen, oldSize;
    if (!strcmp(endWord, stack[stackSize - 1]))
    {
        printArray(stack, stackSize);
        *minlen = *minlen > *len ? *len : *minlen;
        return;
    }
    for (int i = 0; i < wordListSize; i++)
    {
        if (canChange(stack[stackSize - 1], wordList[i], slen) && !visited[i])
        {
            visited[i] = 1;
            oldSize = stackSize;
            oldLen = *len;
            stack[stackSize++] = wordList[i];
            *len = *len + 1;
            ladderLengthDFS(endWord, visited, stack, stackSize, len, minlen, wordList, wordListSize, slen);
            stackSize = oldSize;
            *len = oldLen;
            visited[i] = 0;
        }
    }
}