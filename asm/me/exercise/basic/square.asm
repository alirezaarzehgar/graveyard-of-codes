%include "linux64.inc"

section .data
	nl	db	10, 0
	x	db	"X", 0
	square	db	" square: ", 0
	start	db	"* ", 0
	
section .bss
	num	resb	10
	buf	resb	2

section .text
	global _start
_start:
	getline num, 10
	atoi num, r10

	printnum r10d, buf
	printchar x
	printnum r10d, buf
	printstr square
	printchar nl

	mov	r14d, r10d
loop_top:
	cmp	r14d, 0
	je	end_main
	dec	r14d

	mov	r15d, r10d
loop_print:
	cmp	r15d, 0
	je	print
	dec	r15d

	printstr start
	loop	loop_print

print:
	printchar nl
	loop	loop_top

end_main:
	exit 0
