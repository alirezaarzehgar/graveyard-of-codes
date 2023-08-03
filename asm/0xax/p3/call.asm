global _start

section .data
	msg	db	"Hello, World!", 10

section .text

_start:
		mov rax, 1
		call incRax
		cmp rax, 2
		jne .exit

		mov	rax, 1
		mov	rdi, 1
		mov	rsi, msg
		mov	rdx, 14
		syscall

		jmp	.exit

.exit:
	mov	rax, 60
	mov	rdi, 0
	syscall

incRax:
		inc rax
		ret
