; void strprint(string msg)
; print a string
; rely on strlen
%include 'strlen.asm'

strprint:
    call    internal_strprint   ; call internal_strprint
; Print a linefeed via internal_strprint
    push    rax                 ; push rax on the stack
    mov     rax, 0xa            ; copy a linefeed to rax
    push    rax                 ; push rax again to get an address and not the value
    mov     rax, cr_lf          ; now get back the address of the 'string' linefeed
    call    internal_strprint   ; call internal_strprint to print the linefeed
; Now, we need to remove the linefeed from the stack and restore the original
; value of eax
    pop     rax                 ; remove linefeed from stack
    pop     rax                 ; restore original value
    ret                         ; return

; ************************************

; Internal strprint doing the actual job
internal_strprint:
; Print the string
    push    rdx                 ; we need to push these
    push    rcx                 ;   four registers' values
    push    rbx                 ;   on the stack as they
    push    rax                 ;   will be used here
    call    strlen              ; compute the length of the string
    mov     rdx, rax            ; copy result of strlen (in rax) into rdx
    pop     rax                 ; restore value of rax, i.e. the address of the msg
    mov     rcx, rax            ; copy into rcx
    mov     rbx, 1              ; file descriptor: stdout
    mov     rax, 4              ; sys_write
    int     0x80                ; kernel call
; Restore values and return
    pop     rbx                 ; we need to pop from
    pop     rcx                 ;   the stack in order
    pop     rdx                 ;   to restore their values
    ret                         ; return

cr_lf:
    db 0xd, 0xa, '$'            ; carriage return and linefeed
