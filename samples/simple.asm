BITS 32
    org 0x7c00
start:
    mov eax, 0x00f1
    mov ebx, 0x0029
    call add_routine
    jmp 0
add_routine:
    mov ebx, eax
    mov eax, 0x1011
