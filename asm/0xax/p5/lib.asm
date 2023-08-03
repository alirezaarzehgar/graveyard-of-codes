print:
	mov	rax, 1
	mov	rdi, 1
	; rsi selected
	; rdx selected
	syscall
	ret

exit:
	mov	rax, 60
	mov	rdi, 0
	syscall
	ret
