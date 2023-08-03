#include <stdio.h>

int func()
{
	asm("mov $9090, %rax");
}

extern int fromasm();

int main(int argc, char const *argv[])
{
	printf("%d\n", func());
	printf("%d\n", fromasm());
	return 0;
}
