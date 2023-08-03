%include "linux64.inc"

section .data
	nl	db	10
section .bss
	buf	resb	1024
	num	resb	1
section .text
	global _start
_start:
	getline buf, 1024

	xor	r10, r10
loop_push:
	inc	r10
	mov	al, 10
	cmp	[buf+r10], al
	jne	loop_push

	dec	r10
print:
	mov	al, [buf+r10]
	mov	[num], al
	printchar num

	dec	r10
	cmp	r10, 0
	jge	print

	printchar nl
	exit 0
