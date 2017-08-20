; int factorial(int)

section .text

    global factorial

factorial:
    cmp     rdi, 0x1            ; if the number N is less than or equal to 1,
    jnbe    recursive           ; then jump to recursive
    mov     rax, 0x1            ; else return 1
    ret

recursive:
    push    rdi                 ; store the integer N on the stack
    dec     rdi                 ; decrement N
    call    factorial           ; here is the whole recursivity, calling factorial(N-1)
    pop     rdi                 ; restore original value
    imul    rax, rdi            ; perform the multiplication N * factorial(N-1)
    ret

