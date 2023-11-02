#include <iostream>
#include <vector>
#include <stack>

using namespace std;

#define MAX_STACK_SIZE 25

template <typename T>
class VectorStack
{
private:
	vector<T> stack_storage;

public:
	void push(T v)
	{
		stack_storage.push_back(v);
	}

	void pop()
	{
		stack_storage.pop_back();
	}

	T top()
	{
		return stack_storage.back();
	}
};

template <typename T>
class ArrayStack
{
private:
	int current = 0;
	T stack_storage[MAX_STACK_SIZE];

public:
	void push(T v)
	{
		stack_storage[current++] = v;
	}

	void pop()
	{
		stack_storage[--current] = (T)NULL;
	}

	T top()
	{
		return stack_storage[current-1];
	}
};

int main()
{
	// Using std::stack
	stack<string> name_stack;
	string names[] = {"ali", "mohammad", "reza", "hamid"};

	for (auto &&name : names)
		name_stack.push(name);
	while (!name_stack.empty()) {
		cout << name_stack.top() << " ";
		name_stack.pop();
	}
	cout << endl;

	// Our stack implementation with vector
	VectorStack<int> vstack;
	for (int i = 0; i < 12; i++)
		vstack.push(i*i);
	for (int i = 0; i < 12; i++) {
		cout << vstack.top() << " ";
		vstack.pop();
	}
	cout << endl;

	// Our stack implementation with array
	ArrayStack<int> astack;
	for (int i = 0; i < 12; i++)
		astack.push(i*2);
	for (int i = 0; i < 12; i++) {
		cout << astack.top() << " ";
		astack.pop();
	}
	cout << endl;

	return (EXIT_SUCCESS);
}
