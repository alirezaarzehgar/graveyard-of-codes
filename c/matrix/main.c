/**
 * @file main.c
 * @author Alireza Arzehgar (alirezaarzehgar@gmail.com)
 * @brief Product two matrixes
 * @version 0.1
 * @date 2024-06-13
 *
 * @copyright Copyright (c) 2024
 *
 */
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

int main()
{
	bool debug = false;

#ifdef DEBUG
	debug = true;
#endif

	int an, am, bn, bm;

	printf("Enter n,m of matrix A: ");
	scanf("%d,%d", &an, &am);

	printf("Enter n,m of matrix B: ");
	scanf("%d,%d", &bn, &bm);

	if (am != bn) {
		fprintf(stderr, "Count of A rows should equal to B colums");
		exit(EXIT_FAILURE);
	}

	int A[an][am], B[bn][bm], OUT[an][bm];

	printf("Enter matrix A:\n");
	for (size_t i = 0; i < an; i++) {
		for (size_t j = 0; j < am; j++)
			scanf("%d", &A[i][j]);
	}

	printf("Enter matrix B:\n");
	for (size_t i = 0; i < bn; i++) {
		for (size_t j = 0; j < bm; j++)
			scanf("%d", &B[i][j]);
	}

	for (size_t i = 0; i < an; i++) {
		for (size_t k = 0; k < bm; k++) {
			int sum = 0;
			for (size_t j = 0; j < bn; j++) {
				sum += A[i][j] * B[j][k];
				if (debug)
					printf("(%d*%d) %c ", A[i][j], B[j][k], (j == bn-1)? '=' : '+');
			}
			OUT[i][k] = sum;
			if (debug)
				printf("%d\n", sum);
		}
	}

	printf("A*B:\n");
	for (size_t i = 0; i < an; i++) {
		for (size_t j = 0; j < bm; j++)
			printf("%d\t", OUT[i][j]);
		printf("\n");
	}

	return (EXIT_SUCCESS);
}
