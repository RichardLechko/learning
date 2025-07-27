	.file	"constants-dos-and-donts.c"
# GNU C17 (Ubuntu 13.3.0-6ubuntu2~24.04) version 13.3.0 (x86_64-linux-gnu)
#	compiled by GNU C version 13.3.0, GMP version 6.3.0, MPFR version 4.2.1, MPC version 1.3.1, isl version isl-0.26-GMP

# GGC heuristics: --param ggc-min-expand=100 --param ggc-min-heapsize=131072
# options passed: -mtune=generic -march=x86-64 -O0 -fasynchronous-unwind-tables -fstack-protector-strong -fstack-clash-protection -fcf-protection
	.text
	.globl	main
	.type	main, @function
main:
.LFB0:
	.cfi_startproc
	endbr64	
	pushq	%rbp	#
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp	#,
	.cfi_def_cfa_register 6
	subq	$32, %rsp	#,
# constants-dos-and-donts.c:3: int main() {
	movq	%fs:40, %rax	# MEM[(<address-space-1> long unsigned int *)40B], tmp124
	movq	%rax, -8(%rbp)	# tmp124, D.3190
	xorl	%eax, %eax	# tmp124
	movq	%rsp, %rax	#, tmp98
	movq	%rax, %rsi	# tmp98, saved_stack.1_16
# constants-dos-and-donts.c:15:     const int arraySize = arraySizes[0]; // this is also OK
	movl	arraySizes.0(%rip), %eax	# arraySizes[0], tmp99
	movl	%eax, -28(%rbp)	# tmp99, arraySize
# constants-dos-and-donts.c:17:     int array[ arraySize ]; 
	movl	-28(%rbp), %eax	# arraySize, tmp100
	cltq
	subq	$1, %rax	#, _2
	movq	%rax, -24(%rbp)	# _3, D.3180
	movl	-28(%rbp), %eax	# arraySize, tmp104
	cltq
	leaq	0(,%rax,4), %rdx	#, _13
	movl	$16, %eax	#, tmp122
	subq	$1, %rax	#, tmp105
	addq	%rdx, %rax	# _13, tmp106
	movl	$16, %edi	#, tmp123
	movl	$0, %edx	#, tmp109
	divq	%rdi	# tmp123
	imulq	$16, %rax, %rax	#, tmp108, tmp110
	movq	%rax, %rcx	# tmp110, tmp112
	andq	$-4096, %rcx	#, tmp112
	movq	%rsp, %rdx	#, tmp113
	subq	%rcx, %rdx	# tmp112, tmp113
.L2:
	cmpq	%rdx, %rsp	# tmp113,
	je	.L3	#,
	subq	$4096, %rsp	#,
	orq	$0, 4088(%rsp)	#,
	jmp	.L2	#
.L3:
	movq	%rax, %rdx	# tmp110, tmp114
	andl	$4095, %edx	#, tmp114
	subq	%rdx, %rsp	# tmp114,
	movq	%rax, %rdx	# tmp110, tmp115
	andl	$4095, %edx	#, tmp115
	testq	%rdx, %rdx	# tmp115
	je	.L4	#,
	andl	$4095, %eax	#, tmp116
	subq	$8, %rax	#, tmp116
	addq	%rsp, %rax	#, tmp117
	orq	$0, (%rax)	#,
.L4:
	movq	%rsp, %rax	#, tmp111
	addq	$3, %rax	#, tmp118
	shrq	$2, %rax	#, tmp119
	salq	$2, %rax	#, tmp120
	movq	%rax, -16(%rbp)	# tmp120, array.0
	movq	%rsi, %rsp	# saved_stack.1_16,
	movl	$0, %eax	#, _26
# constants-dos-and-donts.c:20: }
	movq	-8(%rbp), %rdx	# D.3190, tmp125
	subq	%fs:40, %rdx	# MEM[(<address-space-1> long unsigned int *)40B], tmp125
	je	.L6	#,
	call	__stack_chk_fail@PLT	#
.L6:
	leave	
	.cfi_def_cfa 7, 8
	ret	
	.cfi_endproc
.LFE0:
	.size	main, .-main
	.section	.rodata
	.align 8
	.type	arraySizes.0, @object
	.size	arraySizes.0, 8
arraySizes.0:
	.long	123
	.long	256
	.ident	"GCC: (Ubuntu 13.3.0-6ubuntu2~24.04) 13.3.0"
	.section	.note.GNU-stack,"",@progbits
	.section	.note.gnu.property,"a"
	.align 8
	.long	1f - 0f
	.long	4f - 1f
	.long	5
0:
	.string	"GNU"
1:
	.align 8
	.long	0xc0000002
	.long	3f - 2f
2:
	.long	0x3
3:
	.align 8
4:
