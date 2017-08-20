; void exit()
; Exit program
exit:
    mov     ebx, 0          ; no errors status
    mov     eax, 1          ; sys_exit
    int     0x80            ; kernel call
    ret                     ; return

