#include <iostream>

using namespace std;

#define QUEUE_CAP	6

int rear = -1, front = -1, queue[QUEUE_CAP+1];

// Often named enqueue
void insert(int order)
{
	if (front == (rear+1)%QUEUE_CAP) {
		cout << "queue is full" << endl;
		return;
	}
	queue[rear] = order;
	rear++;
}

// Often named dequeue
int del()
{
	if (front == (front+1)%QUEUE_CAP) {
		cout << "queue is empty" << endl;
		return 0;
	}
	
	return queue[front++];
}

int main()
{
	rear = front = 0;

	for (int i = 0; i < 6; i++)
		insert(i);

	for (int i = 0; i < QUEUE_CAP; i++) {
		cout << queue[i] << " ";
	}
	cout << endl;
	

	for (int i = 0; i < 6; i++)
		cout << "enqueued order:" << del() << endl;
	
	return 0;
}
