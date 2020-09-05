#include <stdio.h>
#include "./lib.h"

typedef struct Open_return Open_return;

int main() {
    GoString gs = {.p = "test", .n=4};
    char * res = TestGoFunction(gs);
    printf("result = %s\n",res);
//    Open_return openResult = Open("file:memory:?mode=memory",0);
//    if (!openResult.r0 || !openResult.r1) {
//        printf("failed to open database connection\n");
//    }
    return 0;
}