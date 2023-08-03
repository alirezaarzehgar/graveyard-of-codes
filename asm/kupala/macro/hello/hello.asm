%include "helper.inc"

section .data
	text	db	"Hello, World!", 10

section .text
	global _start

_start:
	print text
	exit 0
