# Student info : 2019125021 남서아 소프트웨어학과

.data
    num:     .word 0x3
    answer:  .word 0x1

.text
    main:
        # Call the factorial function
        lw  $a0, num
        jal factorial
        sw  $v0, answer

        # Display the results
        li  $v0, 1
        lw  $a0, answer
        syscall

        # end the program
        li  $v0, 10
        syscall
    
    factorial:
        bgt     $a0, 0, recur # if (n <= 0)
        li      $v0, 1
        jr      $ra
    
    recur:
        addi    $sp, $sp, -8 # stack push
        sw      $a0, 4($sp) # n
        sw      $ra, 0($sp) # func return address

        addi    $a0, $a0, -1 # n-1
        jal     factorial
        lw      $t0, 4($sp) # n
        mul     $v0, $v0, $t0 # answer *= n

        lw      $ra, 0($sp) # return address
        addi    $sp, $sp, 8 # stack pop
        jr      $ra
