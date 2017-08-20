; void count(int n)
; print numbers from 1 to n
; rely on intprint and exit
%include 'intprint.asm'
%include 'exit.asm'

section .data
    msg db 'Enter an integer (max 255 bytes): ', 0x0            ; msg asking for an integer

section .bss
    input:  resb    255                                         ; reserve a 255 bytes space in memory

section .text

    global _start

_start:
; Print the introduction message
    mov     rax, msg
    call    strprint
; Read input from user
    mov     rdx, 255                                            ; store number of bytes to be read into rdx
    mov     rcx, input                                          ; buffer
    mov     rbx, 0                                              ; file descriptor: stdin
    mov     rax, 3                                              ; sys_read
    int     0x80                                                ; kernel call
; start the counter
    xor     rcx, rcx                                            ; start counter to 0
next:
    inc     rcx                                                 ; increment counter
    mov     rax, rcx                                            ; move counter into rax to print
    call    intprint                                            ; call print for integers
    cmp     rcx, input
    jne     next                                                ; loop
    call    exit                                                ; call exit

