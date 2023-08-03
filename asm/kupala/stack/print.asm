section .data
	msg1	db	"Hello, World!", 10, 0
	msg2	db	"World?", 10, 0

section .text
	global _start

_start:
	mov	rax, msg1
	call	_print

	mov	rax, msg2
	call	_print

	mov	rax, 60
	mov	rdi, 0
	syscall

_print:
	push	rax
	mov	rbx, 0
_print_loop:
	inc	rax
	inc	rbx
	mov	cl, [rax]
	cmp	cl, 0
	jne	_print_loop

	mov	rax, 1
	mov	rdi, 1
	pop	rsi
	mov	rdx, rbx
	syscall
	ret
