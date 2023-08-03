section .bss
	strp	resb	10
	buf	resb	1

section	.text
	global	_start

_start:
	mov     rax, 0
	mov     rdi, 0
	mov     rsi, strp
	mov     rdx, 10
	syscall

	xor	r10, r10
push2stack:
	mov     al, [strp+r10]
	push    rax

	inc     r10
	mov     al, 10
	cmp     [strp+r10], al
	jne     push2stack

print_stack:
	pop     rax
	call    print_rax
	dec     r10
	jnz     print_stack

	mov     rax, 10
	call    print_rax

	mov     rax, 60
	mov     rdi, 0
	syscall

print_rax:
	mov	[buf], al

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, buf
	mov	rdx, 1
	syscall
	ret

