section .data
	SYS_WRITE	equ	1
	SYS_EXIT	equ	60
	STDOUT		equ	1
	EXIT_CODE	equ	0

	NEWLINE		db	10
	WRONG_ARGS	db	"Must be two command line argument", 10
	WRONG_ARGS_LEN	equ	34

section .text
	global _start

_start:
	pop	rcx
	cmp	rcx, 3
	jne	argcError

	add	rsp, 8
	pop	rsi
	call	str2int
	mov	r10, rax

	pop	rsi
	call	str2int
	mov	r11, rax

	add	r10, r11
	mov	rax, r10
	xor	r12, r12
	jmp	int2str

	call exit

argcError:
	mov	rax, SYS_WRITE
	mov	rdi, STDOUT
	mov	rsi, WRONG_ARGS
	mov	rdx, WRONG_ARGS_LEN
	syscall
	call exit

str2int:
	mov	rax, rax
	mov	rcx, 10

next:
	cmp	[rsi], byte 0
	je	return_str
	mov	bl, [rsi]
	sub	bl, 48
	mul	rcx
	add	rax, rbx
	inc	rsi
	jmp	next

return_str:
	ret

int2str:
	mov	rdx, 0
	mov	rbx, 10
	div	rbx
	add	rdx, 48
	add	rdx, 0x0
	push	rdx
	inc	r12
	cmp	rax, 0x0
	jne	int2str
	jmp	print

print:
	mov	rax, 1
	mul	r12
	mov	r12, 8
	mul	r12
	mov	rdx, rax

	mov	rax, SYS_WRITE
	mov	rdi, STDOUT
	mov	rsi, rsp
	syscall

	jmp exit

exit:
	mov	rax, SYS_EXIT
	mov	rdi, EXIT_CODE
	syscall
