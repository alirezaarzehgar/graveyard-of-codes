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

int main()
{
	int arr[] = {5, 1, 4, 7, 6, 2, 3};
	int out, key = 4, n = sizeof(arr)/sizeof(arr[0]);

	out = bsearch(arr, n, key);
	if (out == -1) {
		cout << "key not found" << endl;
		return -1;
	} else {
		cout << "index:" << out << endl;
	}

	return 0;
}
