; sub.s
; tests for subroutines
; 
; Copyright (C) 2018-present Karim Alibhai.

ORG 0
  BSA PRINT
  HLT

; print should be able to print -CTR letters
; to the screen starting at PTR
PRINT: HEX 0
LOOP:  LDA PTR I
       UT
       ISZ PTR
       ISZ CTR
       BUN LOOP
       BUN PRINT I

PTR: DEC 10 ; hard-coded as the next line
HEX 68
HEX 65
HEX 6C
HEX 6C
HEX 6F
HEX 0A
CTR: DEC -6
