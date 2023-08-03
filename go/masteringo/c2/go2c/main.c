#include "used_by_c.h"
#include <stdio.h>

int main(int argc, char const *argv[])
{
	PrintMessage();
	GoInt num = Multiply(3, 5);
	printf("%d\n", num);
	return 0;
}
