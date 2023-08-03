section .data
	nl	db	10
	time	dq	0, 100000000

section .bss
	buf	resb	255
	char	resb	1

section .text
	global _start
_start:
	mov	rax, 0		; sys_read
	mov	rdi, 0		; stdin
	mov	rsi, buf
	mov	rdx, 255
	syscall

	mov	rax, buf
loop_str:
	mov	cl, 0
	cmp	[rax], cl
	je	end_main

	push	rax
	mov	rax, 35		; sys_nanosleep
	mov	rdi, time
	mov	rsi, 0
	syscall
	pop	rax

	mov	cl, [rax]
	mov	[char], cl

	push	rax
	mov	rax, 1		; sys_write
	mov	rdi, 1		; stdout
	mov	rsi, char
	mov	rdx, 1
	syscall
	pop	rax

	inc	rax
	loop	loop_str

end_main:
	mov	rax, 60		; sys_exit
	mov	rdi, 0
	syscall
