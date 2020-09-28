#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void insertSort(int *nums, int numsSize)
{
    int i = 1, j, val;
    for (; i < numsSize; i++)
    {
        val = nums[i];
        j = i - 1;
        // from small to big
        while (j >= 0 && nums[j] > val)
        {
            nums[j + 1] = nums[j];
            j--;
        }
        nums[j + 1] = val;
    }
}

int hasGroupsSizeX(int *deck, int deckSize)
{
    if (deckSize < 2 || deck == NULL)
    {
        return 0;
    }
    insertSort(deck, deckSize);
    int n, count;
    for (int x = 2; x <= deckSize; x++)
    {
        if (deckSize % x != 0)
        {
            continue;
        }
        n = deckSize / x;
        count = 0;
        for (int i = 0; i < deckSize; i += x)
        {
            if (deck[i] != deck[i + x - 1])
            {
                break;
            }
            count++;
        }

        if (n == count)
        {
            return 1;
        }
    }
    return 0;
}
int main(int argc, char *argv[])
{
    int a[2] = {1, 1};
    printf("%d\n", hasGroupsSizeX(a, 2));
    return 0;
}
