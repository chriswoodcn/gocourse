// gcc -o main main.c hello.so => main
// exec main
#include<stdio.h>
#include"hello.h"

int main(){
    printf("use hello lib from C\n");
    //C->go
    HelloFromGo();
    return 0;
}
