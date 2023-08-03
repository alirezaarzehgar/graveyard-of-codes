%include "linux64.inc"

section .data
	data	times 10 dw 0
	nl	db	10
	name	db 	"ali", 10
section .bss
	num	resb	10
section .text
	global _start
_start:
	printstr name
;	mov	[name], dword "Mohammad", 10

	printstr name
	exit 0
