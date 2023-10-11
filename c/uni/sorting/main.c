#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <errno.h>
#include <getopt.h>

#define DEFAULT_ARR_SIZE 40

void bubble_sort(int *arr, size_t n)
{
	for (size_t i = 0; i < n - 1; i++) {
		bool swapped = false;

		for (size_t j = 0; j < n - i - 1; j++) {
			if (arr[j] > arr[j+1]) {
				int tmp = arr[j];
				arr[j] = arr[j+1];
				arr[j+1] = tmp;

				swapped = true;

			}
		}

		if (!swapped)
			break;		
	}
}

void insertion_sort(int *arr, size_t n)
{
	for (size_t j, i = 1; i < n; i++) {
		int key = arr[i];

		for (j = i - 1; arr[j] > key; j--)
			arr[j + 1] = arr[j];
		arr[j + 1] = key;
	}
}

int main(int argc, char *const *argv)
{
	int ch, *arr;
	size_t n = 0;
	void(*sorting_algorithm)(int *, size_t) = bubble_sort;

	while ((ch = getopt(argc, argv, "bi")) != -1) {
		switch (ch)
		{
		case 'b':
			printf("Sort with bubble sort algorithm.\n");
			sorting_algorithm = bubble_sort;
			break;
		case 'i':
			printf("Sort with insertion sort algorithm.\n");
			sorting_algorithm = insertion_sort;
			break;
		default:
			fprintf(stderr, "usage: sorting [bi]");
			return EXIT_FAILURE;
		}
	}
	argc -= optind;
	argv += optind;	

	arr = malloc(sizeof(int) * DEFAULT_ARR_SIZE);
	if (!arr && errno) {
		perror("malloc");
		exit(EXIT_FAILURE);
	}

	while (scanf("%d", &arr[n++]) > 0) {
		if (n % DEFAULT_ARR_SIZE == 0) {
			arr = realloc(arr, sizeof(int) * n + DEFAULT_ARR_SIZE);
			if (!arr && errno) {
				perror("realloc");
				exit(EXIT_FAILURE);
			}
		}
	}
	n--;

	sorting_algorithm(arr, n);

	for (size_t i = 0; i < n; i++)
		printf("%d ", arr[i]);
	putchar('\n');

	return EXIT_SUCCESS;
}
