STDIN		equ	0
STDOUT		equ	1
STDERR		equ	2

SYS_READ	equ	0
SYS_WRITE	equ	1
SYS_EXIT	equ	60

%macro exit 1
	mov	rax, SYS_EXIT
	mov	rdi, %1
	syscall
%endmacro

%macro print 1
	mov	rax, %1
	push	rax
	xor	rbx, rbx

%%loop:
	inc	rax
	inc	rbx

	mov	cl, [rax]
	cmp	cl, 0
	jne	%%loop

	mov	rax, SYS_WRITE
	mov	rdi, STDOUT
	mov	rsi, %1
	mov	rdx, rbx
	syscall
%endmacro
