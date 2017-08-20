; Hello world!
; rely on the following functions:
;   strprint
;   exit
%include    'strprint.asm'
%include    'exit.asm'

section .data
    msg  db 'Hello world!', 0x0     ; the message followed by a terminating null byte

section .text

    global _start

_start:
    mov     rax, msg                ; mov the message to actually write
    call    strprint                ; print the message
    call    exit                    ; exit the program
