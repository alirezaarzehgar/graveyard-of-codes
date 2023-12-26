#include <iostream>

using namespace std;

int main()
{
	for (size_t i = 0; i < 4; i++) {
		for (size_t j = 0; j <= i; j++)
			printf("* ");
		putchar('\n');
		
	}
	
	return 0;
}
