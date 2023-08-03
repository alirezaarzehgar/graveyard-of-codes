#include <stdio.h>

extern int addints(int a, int b, int c, int d);
extern int mulby17(int a);
extern int sub5x(int a, int b);

int main()
{
	int a, b, c, d;

	scanf("%d%d%d%d", &a, &b, &c, &d);

	printf("a + b + c + d = %d\n", addints(a, b, c, d));
	printf("a * 17 = %d\n", mulby17(a));
	printf("(a - 5) * b = %d\n", sub5x(a, b));

	return 0;
}
