section .data
	count	equ	10

section .bss
	digit	resb	2

section .text
	global _start
_start:
	mov	al, 10
	mov	[digit+1], al

	xor	r10, r10
loop:
	mov	rax, r10
	call	printNum

	inc	r10
	cmp	r10, count
	jne	loop

	mov	rax, 60
	mov	rdi, 0
	syscall

printNum:
	add	rax, 48
	mov	[digit], al
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, digit
	mov	rdx, 2
	syscall
	ret
