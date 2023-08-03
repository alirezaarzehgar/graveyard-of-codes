#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "l2.h"

char *createMsg()
{
	char *msg = malloc(14);
	strcat(msg, "Hello, World\n");
	return msg;
}

void printMsg(char *msg)
{
	printf("%s", msg);
}