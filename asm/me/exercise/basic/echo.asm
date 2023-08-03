%include "linux64.inc"

section .data
	nl	db	10
	space	db	" "
section .bss
	num	resb	1
section .text
	global _start
_start:
	pop	r10

pop	r11
loop_echo:
	dec	r10
	jz	end_main

	pop	r11
	printstr r11
	printchar space

	loop	loop_echo

end_main:
	printchar nl
	exit 0
