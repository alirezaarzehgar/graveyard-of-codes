#include <iostream>

using namespace std;

int merge(int *a, int *b, int n, int m)
{
	int i = m, j = 0, k = 0;

	while (k < m+n) {
		if (((i < m+n) && a[i] < b[j]) || j == m) {
			a[k] = a[i];
			i++;
		} else {
			a[k] = b[j];
			j++;
		}
		k++;
	}

	return n+m;
}

int main()
{
	int arr1[] = {-1, -1, -1, -1, 5, 1, 4}, arr2[] = {6, 2, 7, 8};
	int out, n = 3, m = 4;

	out = merge(arr1, arr2, n, m);
	for (int i = 0; i < out; i++)
		cout << arr1[i] << " ";
	cout << endl;

	return 0;
}
