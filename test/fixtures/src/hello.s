; hello.s
; simple hello world program
;
; Copyright (C) 2018-present Karim Alibhai.

ORG 0
  LDA H
  UT
  LDA E
  UT
  LDA L
  UT
  UT
  LDA O
  UT
  LDA LF
  UT
  HLT

H: HEX 68
E: HEX 65
L: HEX 6C
O: HEX 6F
LF: HEX 0A
