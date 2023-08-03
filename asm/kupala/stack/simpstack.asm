section .data
	digit	db	0, 10

section .text
	global _start

_start:
	push	1
	push	2
	push	3
	push	4

	pop	rax
	call	print_rax_num
        pop     rax
        call    print_rax_num
        pop     rax
        call    print_rax_num
        pop     rax
        call    print_rax_num

	mov	rax, 60
	mov	rdi, 0
	syscall

print_rax_num:
	add	rax, 48
	mov	[digit], al

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, digit
	mov	rdx, 2
	syscall
	ret

