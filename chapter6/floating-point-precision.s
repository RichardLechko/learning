	.file	"floating-point-precision.c"
	.text
	.globl	main
	.type	main, @function
main:
.LFB6:
	.cfi_startproc
	endbr64
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset 6, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register 6
	movl	%edi, -4(%rbp)
	movq	%rsi, -16(%rbp)
	movq	%rdx, -24(%rbp)
	movl	$0, j.2(%rip)
	movsd	i.1(%rip), %xmm1
	movl	j.2(%rip), %eax
	movsd	.LC0(%rip), %xmm0
	addsd	%xmm1, %xmm0
	cltq
	leaq	0(,%rax,8), %rdx
	leaq	a.0(%rip), %rax
	movsd	%xmm0, (%rdx,%rax)
	movl	$0, %eax
	popq	%rbp
	.cfi_def_cfa 7, 8
	ret
	.cfi_endproc
.LFE6:
	.size	main, .-main
	.local	j.2
	.comm	j.2,4,4
	.data
	.align 8
	.type	i.1, @object
	.size	i.1, 8
i.1:
	.long	0
	.long	1072693248
	.align 32
	.type	a.0, @object
	.size	a.0, 64
a.0:
	.long	0
	.long	1072693248
	.long	0
	.long	1073741824
	.long	0
	.long	1074266112
	.long	0
	.long	1074790400
	.long	0
	.long	1075052544
	.long	0
	.long	1075314688
	.long	0
	.long	1075576832
	.zero	8
	.section	.rodata
	.align 8
.LC0:
	.long	0
	.long	1072693248
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
