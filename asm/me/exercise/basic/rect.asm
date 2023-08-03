%include "linux64.inc"

section .data
	nl	db	10
	star	db	"* "
	space	db	"  "

section .bss
	num	resb	10

section .text
	global _start
_start:
	getline num, 10
	atoi num, r10


	xor	r14, r14
loop_col:
	cmp	r14, r10
	jge	end_main
	inc	r14

	xor	r15, r15
loop_row:
	cmp	r15, r14
	jge	end_loop_row
	inc	r15

	printstr star

	loop	loop_row

end_loop_row:
	printchar nl

	loop	loop_col

end_main:	
	exit 0
