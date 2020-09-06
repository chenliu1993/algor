#include <stdlib.h>
#include <stdio.h>

int main()
{
  size_t size;
  char *size_str;
  size_str = getenv("TEST");
  if(size_str!=NULL){
    size_t buf_size = strtoul(size_str, NULL, 0);
    printf("%p\n", size_str);
    printf("%s\n", size_str);
    printf("%lu\n", sizeof(buf_size));
    // free(*size_str);
  }
  return 0;
}