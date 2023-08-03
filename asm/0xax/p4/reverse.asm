section .data
	SYS_WRITE equ 1
	STD_OUT   equ 1
	SYS_EXIT  equ 60
	EXIT_CODE equ 0

	NEW_LINE db 0xa
	INPUT db "Hello world!"

section .bss
	OUTPUT	resb	12

section .text
	global _start

_start:
	mov	rsi, INPUT
	xor	rcx, rcx
	cld
	mov	rdi, $ + 15
	call	calculateStrLen
	xor	rax, rax
	xor	rdi, rdi
	jmp	reverseStr

calculateStrLen:
	cmp	byte [rsi], 0
	je	ret

	lodsb
	push	rax
	inc	rcx
	jmp	calculateStrLen

ret:
	push	rdi
	ret

reverseStr:
	cmp	rcx, 0
	je	print
	pop	rax
	mov	[OUTPUT + rdi], rax
	dec	rcx
	inc	rdi
	jmp	reverseStr

print:
	mov	rdx, rdi
	mov	rax, SYS_WRITE
	mov	rdi, STD_OUT
	mov	rsi, OUTPUT
	syscall
	call	newline

	mov	rax, 60
	mov	rdi, 0
	syscall

newline:
	mov	rax, SYS_WRITE
	mov	rdi, STD_OUT
	mov	rsi, NEW_LINE
	mov	rdx, 1
	syscall
	ret
