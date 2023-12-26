#!/usr/bin/env python3
class Fibonacci:
	def __init__(self, n):
		self.n = n
		self.index = 0
		self.current = 0
		self.next = 1

	def __iter__(self):
		return self

	def __next__(self):
		if self.index >= self.n:
			raise StopIteration

		self.current, self.next = self.next, self.current + self.next
		self.index += 1
		return self.current

if __name__ == "__main__":
	for v in iter(Fibonacci(5)):
		print(v)
