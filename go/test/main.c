#include <stdio.h>

int main()
{
	int sum = 0;
	int list[] = {856, 345, 189, 375, 902, 197,
	              528, 847, 186, 912, 986, 387,
	              192, 559, 951, 441, 259, 100,
	              334, 449, 214, 904, 120, 205
	             };

	for (int i = 0; i < sizeof(list) / sizeof(int); i++) {
		int t = (list[i] / 10) % 10;
		if (t % 2 == 1)
			sum++;
	}

  printf("%d\n", sum);

	return 0;
}
