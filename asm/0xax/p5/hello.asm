; bad

struc person
	name	resb	10
	age	resb	1
endstruc

section	.data
	msg	db	"hello, world!", 10
	msglen	equ	14

	p:
		istruc person
			at name, db "ali", 10
			at age, db 25
		iend

section .text
	global _start

_start:
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, [person.name]
	mov	rdx, 4
	syscall

	mov	rax, 60
	mov	rdi, 0
	syscall
