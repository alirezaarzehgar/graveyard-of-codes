#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <shadow.h>
#include <crypt.h>
#include <err.h>

int main()
{
	char *password, username[30];
	struct spwd *pwd;

	printf("User: ");
	scanf("%s", username);
	password = getpass("Password: ");

	pwd = getspnam(username);
	if (pwd == NULL) {
		printf("unauthorized\n");
		return 1;
	}

	password = crypt(password, pwd->sp_pwdp);
	if (strcmp(pwd->sp_pwdp, password) == 0) {
		printf("authenticated\n");
		return 0;
	}

	printf("unauthorized\n");
	return 1;
}
