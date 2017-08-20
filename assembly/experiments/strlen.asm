; int strlen(string msg)
; Compute the length of a given string
strlen:
    push    rbx                 ; push value on the stack to preserve
    mov     rbx, rax

.nextchar:
    cmp     byte [rax], 0       ; compare the byte pointed by eax against 0
    jz      .finished           ; if ZF is set, go to finished
    inc     rax                 ; increment eax, our counter
    jmp     .nextchar           ; loop

.finished:
    sub     rax, rbx            ; both initially pointed to the same address
                                ; eax has been incremented in nextchar
                                ; therefore, sub eax, ebx gives the count
    pop     rbx                 ; pop back from the stack
    ret                         ; return
