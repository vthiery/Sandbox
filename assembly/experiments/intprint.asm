; void intprint(integer n)
; print an integer
; rely on strprint
%include 'strprint.asm'

intprint:
    call    internal_intprint       ; call internal_intprint
    push    rax                     ; preserve value
    mov     rax, 0xa
    push    rax
    mov     rax, cr_lf
    call    internal_strprint       ; call internal_strprint from strprint.asm
    pop     rax
    pop     rax
    ret

; ************************************

internal_intprint:
    push    rax
    push    rcx
    push    rdx
    push    rsi
    xor     rcx, rcx                ; initialize counter to 0

divide:
; count bytes to print
    inc     rcx                     ; increment counter
    xor     rdx, rdx                ; empty rdx
    mov     rsi, 0xa                ; copy 10 into rsi
    idiv    rsi                     ; divide rax / rsi
    add     rdx, 0x30               ; convert rdx to its ascii representation
    push    rdx                     ; preserve rdx value
    cmp     rax, 0x0                ; compare rax to 0
    jnz     divide                  ; loop

print:
; print characters from the stack
    dec     rcx                     ; decrement counter
    mov     rax, rsp                ; mov stack ptr into rax
    call    strprint                ; now we can call our string version of print
    pop     rax                     ; remove last char from stack
    cmp     rcx, 0x0                ; are we done yet?
    jnz     print                   ; loop
; restore values and return
    pop     rsi
    pop     rdx
    pop     rcx
    pop     rax
    ret

