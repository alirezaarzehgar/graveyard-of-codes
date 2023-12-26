CFLAGS := -lpam -lpam_misc -lcrypt -lcrypto

all: pam_hello.so pamaware

pam_hello.so: hello.c
	cc -FPIC -c hello.c
	ld --whole-archive -shared $(CFLAGS) -o pam_hello.so hello.o

pamaware:
	cc -FPIC $(CFLAGS) pamaware.c -o pamaware

install: pam_hello.so
	cp hello /etc/pam.d/
	install -s -groot -oroot -t/lib/security pam_hello.so


clean:
	rm -rf *.o *.so pamaware /lib/security/pam_hello.so
