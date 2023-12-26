#include <string.h>
#include <shadow.h>
#include <crypt.h>
#include <security/pam_ext.h>
#include <security/pam_modules.h>

PAM_EXTERN int pam_sm_authenticate(pam_handle_t *pamh,int flags,
                                   int argc,const char **argv)
{
	const char *password, *username;
	struct spwd *pwd;

	pam_get_user(pamh, &username, "User:");
	pam_get_authtok(pamh, PAM_AUTHTOK, &password, "Password:");
	if ((pwd = getspnam(username)) == NULL)
		return PAM_AUTHINFO_UNAVAIL;

	password = crypt(password, pwd->sp_pwdp);
	if (strcmp(pwd->sp_pwdp, password) == 0)
		return PAM_SUCCESS;
	return PAM_AUTH_ERR;
}

PAM_EXTERN int pam_sm_setcred(pam_handle_t *pamh,int flags,
                              int argc,const char **argv)
{
	return PAM_SUCCESS;
}
