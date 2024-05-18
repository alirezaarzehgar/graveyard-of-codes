#include <iostream>

using namespace std;

int find(int *arr, int n, int key)
{
	for (int i = 0; i < n; i++)	{
		if (arr[i] == key)
			return i;
	}
	
	return -1;
}

int main()
{
	int arr[] = {5, 1, 4, 7, 6, 2, 3};
	int out, key = 4, n = sizeof(arr)/sizeof(arr[0]);

	out = find(arr, n, key);
	if (out == -1) {
		cout << "key not found" << endl;
		return -1;
	} else {
		cout << "index:" << out << endl;
	}

	return 0;
}
