BITS 32
    org 0x7c00
    mov eax, 0x1
    cmp eax, 0x2
    jnz not_equal

equal:
    jmp 0

not_equal:
    mov eax, 0x2
    cmp eax, 0x2
    jz equal
