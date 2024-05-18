#include <iostream>

using namespace std;

int del(int *arr, int low, int high, int key)
{
	int mid;

	while(low <= high) {
		mid = (low+high)/2;
		if (arr[mid] == key)
			return mid;
		else if (key < arr[mid])
			return del(arr, low, mid-1, key);
		else
			return del(arr, mid+1, high, key);
	}
	return -1;
}

int main()
{
	int arr[] = {1, 4, 7, 9, 12, 15, 17, 19};
	int out, key = 4, n = sizeof(arr)/sizeof(arr[0]);

	out = del(arr, 0, n, key);
	if (out == -1) {
		cout << "key not found" << endl;
		return -1;
	} else {
		cout << "index:" << out << endl;
	}

	return 0;
}
