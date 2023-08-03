count	equ	5

section .data
	star	db	"* "
	nl	db	10

section .bss
	dbuf	resb	2

section .text
	global _start
_start:
	mov	ebx, count
loop_top:
	cmp	ebx, 0
	je	main_end
	dec	ebx

	mov	r10, count
loop_inner:
	cmp	r10, 0
	je	loop_end
	dec	r10

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, star
	mov	rdx, 2
	syscall

	jmp	loop_inner
loop_end:
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, nl
	mov	rdx, 1
	syscall

	jmp	loop_top


main_end:
	mov	rax, 60
	mov	rdi, 0
	syscall		
