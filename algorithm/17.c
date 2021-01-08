#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <ctype.h>
#include <math.h>

void printGrid(char **res, int size)
{
	for (int i = 0; i < size; i++)
	{
		printf("%s\n", res[i]);
	}
}

char **ans;
int ansSize;
int rear;

char dict[8][4] = {{"abc"}, {"def"}, {"ghi"}, {"jkl"}, {"mno"}, {"pqrs"}, {"tuv"}, {"wxyz"}};

void dfs(char *digits, int k, char *ele)
{
	if (k == strlen(digits))
	{
		ansSize++;
		ans = (char **)realloc(ans, ansSize * sizeof(char *));
		ans[ansSize - 1] = malloc(strlen(digits) * sizeof(char));
		memcpy(ans[ansSize - 1], ele, strlen(digits) * sizeof(char));
		return;
	}

	char *target = dict[digits[k] - '2'];
	for (int i = 0; i < strlen(target); i++)
	{
		ele[rear++] = target[i];
		dfs(digits, k + 1, ele);
		rear--;
	}
}

char **letterCombinations(char *digits, int *returnSize)
{
	ans = malloc(0);
	rear = 0;
	ansSize = 0;
	char *ele = (char *)malloc(strlen(digits) * sizeof(char));
	dfs(digits, 0, ele);
	(*returnSize) = ansSize;
	return ans;
}

int main(int argc, char *argv[])
{
	int *returnSize = (int *)calloc(1, sizeof(int));
	char digits[3] = {"23\0"};
	char **res = letterCombinations(digits, returnSize);
	printGrid(res, *returnSize);
	return 0;
}