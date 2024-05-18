#include <iostream>

using namespace std;

void show(int *arr, int n)
{
	for (int i = 0; i < n; i++)
		cout << arr[i] << " ";
	cout << endl;
}

int insert(int *arr, int n, int c, int index, int val)
{
	if (n >= c)
		return n;

	for (int i = n; i >= index; i--)
		arr[i+1] = arr[i];
	arr[index] = val;

	return n+1;
}

int main()
{
	int arr[7] = {5, 1, 4, 7, 6, 2};
	int out, n = 6, c = 7;
	
	show(arr, n);

	out = insert(arr, n, c, 2, 12);
	if (out == n) {
		cout << "array is full" << endl;
		return -1;
	} else {
		cout << "new n:" << out << endl;
		show(arr, out);
	}

	return 0;
}
