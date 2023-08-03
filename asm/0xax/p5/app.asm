%include "lib.asm"

section .data
	msg	db	"Hello, World!", 10
	msglen	equ	14

section .text
	global _start

_start:
	mov	rsi, msg
	mov	rdx, msglen
	call	print
	call	exit
