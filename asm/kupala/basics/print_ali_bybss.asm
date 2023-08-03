section .bss
	name	resb	10

section .text
	global _start

_start:
	mov	al, 'a'
	mov	[name], al

        mov     al, 'l'
        mov     [name+1], al

        mov     al, 'i'
        mov     [name+2], al

        mov     al, 10
        mov     [name+3], al

	mov	rax, 1
	mov	rdi, 1
	mov	rsi, name
	mov	rdx, 4
	syscall

	mov	rax, 60
	mov	rdi, 0
	syscall

