section .data
	c1_msg 		db	"Condition one is true", 10
	c1_msg_len 	equ	22
	c2_msg 		db	"Condition two is true", 10
	c2_msg_len 	equ	22

section .text
	global _start
_start:
	mov	rax, 10
	cmp	rax, 11
	jge	.cond1
	jmp	.cond2

.cond1:
	mov	rsi, c1_msg
	mov	rdx, c1_msg_len
	jmp	.main

.cond2:
	mov	rsi, c2_msg
	mov	rdx, c2_msg_len
	jmp	.main

.main:
	mov	rax, 1
	mov	rdi, 1
	syscall

	jmp	.exit

.exit:
	mov	rax, 60
	mov	rdi, 0
	syscall
