%include "linux64.inc"

%macro getrand 1
	mov	rax, sys_getrandom
	mov	rdi, %1
	mov	rsi, 1
	mov	rdx, 0
	syscall
%endmacro

section .data
	nl	db	10
section .bss
	rand	resb	1
section .text
	global _start
_start:
	getrand rand

	mov	eax, [rand]
	printnum eax, rand
	printchar nl

	exit 0

