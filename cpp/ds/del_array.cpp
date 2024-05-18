#include <iostream>

using namespace std;

int bsearch(int *arr, int n, int key)
{
	int low = 0, high = n-1, mid;

	while(low <= high) {
		mid = (low+high)/2;
		if (arr[mid] == key)
			return mid;
		else if (key < arr[mid])
			high = mid-1;
		else
			low = mid+1;
	}
	return -1;
}

void show(int *arr, int n)
{
	for (int i = 0; i < n; i++)
		cout << arr[i] << " ";
	cout << endl;
}

int del(int *arr,int n, int key)
{
	int index = bsearch(arr, n, key);
	if (index == -1)
		return n;

	for (int i = index; i < n; i++)
		arr[i] = arr[i+1];

	return n-1;
}

int main()
{
	int arr[] = {5, 1, 4, 7, 6, 2, 3};
	int out, key = 4, n = sizeof(arr)/sizeof(arr[0]);
	
	show(arr, n);

	out = del(arr, n, key);
	if (out == n) {
		cout << "key not found" << endl;
		return -1;
	} else {
		cout << "new n:" << out << endl;
		show(arr, out);
	}

	return 0;
}
