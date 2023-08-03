section .text
	global addints
	global mulby17
	global sub5x

addints:
	mov	eax, edi
	add	eax, esi
	add	eax, edx
	add	eax, ecx
	ret

mulby17:
	mov	eax, edi
	imul	eax, 17
	ret

sub5x:
	mov	eax, edi
	sub	eax, 5
	imul	eax, esi
	ret