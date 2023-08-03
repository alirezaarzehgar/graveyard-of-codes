section .bss
  digit resb  2

section .text
  global  _start
_start:
  mov rax, 2
  call printRAXNumber

  mov rdi, 1
  mov rsi, 1
  mov rdx, 1
  mov rcx, 1
  mov r8, 1
  mov r9, 1

  call sumall

  mov rax, 60
  mov rdi, 0
  syscall

sumall:
  mov rax, rdi
  add rax, rsi
  add rax, rdx
  add rax, rcx
  add rax, r8
  add rax, r9

  call printRAXNumber
  ret

printRAXNumber:
  add rax, 48
  mov [digit], al
  mov  al, 10
  mov [digit+1], al

  mov rax, 1
  mov rdi, 1
  mov rsi, digit
  mov rdx, 2
  syscall
  ret
