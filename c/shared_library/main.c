#include <dlfcn.h>

int main(int argc, char **argv)
{
	void *handle, (*func)();
	handle = dlopen(argv[1], RTLD_LAZY);
	func = dlsym(handle, "func");

	func();

	dlclose(handle);
	return 0;
}
