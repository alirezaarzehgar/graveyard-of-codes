section .data
	num1	equ	100
	num2	equ	50
	msg 	db	"Correct", 10
	msglen	equ	8

section .text
	global _start

_start:
	mov	rax, num1
	mov	rbx, num2
	add	rax, rbx

	cmp	rax, 150
	jne	.exit
	jmp	.correct

.correct:
	mov	rax, 1
	mov	rdi, 1
	mov	rsi, msg
	mov	rdx, msglen
	syscall

	jmp	.exit

.exit:
	mov	rax, 60
	mov	rdi, 0
	syscall
