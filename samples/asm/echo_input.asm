BITS 32
    org 0x7c00
    mov edx, 0x03f8
    in al, dx
    out dx, al
    jmp 0
