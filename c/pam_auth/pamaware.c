#include <pwd.h>
#include <err.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <security/pam_appl.h>
#include <security/pam_misc.h>
#include <security/pam_modules.h>
#include <security/pam_modutil.h>

int main()
{
    pam_handle_t *pamh;
    struct pam_conv pamc = {misc_conv, 0};

    if (pam_start("hello", NULL, &pamc, &pamh))
        err(1, "pam_start");

    if (pam_authenticate(pamh, 0) == PAM_SUCCESS)
        printf("authenticated\n");
    else
        printf("unauthorized\n");

    pam_end(pamh, 0);
    return 0;
}
