# basic

Basic computer from CEG2136. Created by me in an attempt to study for my exam by coding out the content.

It is capable of executing the simple instructions on a 16-bit architecture consisting of an ALU + CU and a data bus. It is similar to our lab 4 assignment but it is written in go rather than with the Altera Quartus platform.

## Usage

Grab it via `go get github.com/karimsa/basic`.

### Compiling & running

`hello.s`:

```asm
ORG 0
  LDA H
  UT
  LDA I
  UT
  LDA LF
  UT
  HLT

H: HEX 68
I: HEX 65
LF: HEX 0A
```

```shell
$ go run ./cmd/asm/asm.go hello.s hello.out
$ go run ./cmd/boot/boot.go hello.out
hi
Halting
```

## Architecture

### Registers

The CPU has 9 registers and a 4096x16 memory unit. The breakdown is from the lecture slides:

![](.github/registers.png)

### Instruction set

Supported instructions (in machine code):

![](.github/instructions.png)

## License

Licensed under MIT license.

Copyright &copy; 2018-present Karim Alibhai.
