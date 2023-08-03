section .data
	path db	"filename.txt"

section .text
	global _start
_start:
	mov	rax, 85
	mov	rdi, path
	mov	rsi, 420
	syscall
	mov	rax, 60
	mov	rdi, 0
	syscall
