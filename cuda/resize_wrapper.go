package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import(
	"unsafe"
	"github.com/MathieuMoalic/amumax/cuda/cu"
	"github.com/MathieuMoalic/amumax/timer"
	"sync"
)

// CUDA handle for resize kernel
var resize_code cu.Function

// Stores the arguments for resize kernel invocation
type resize_args_t struct{
	 arg_dst unsafe.Pointer
	 arg_Dx int
	 arg_Dy int
	 arg_Dz int
	 arg_src unsafe.Pointer
	 arg_Sx int
	 arg_Sy int
	 arg_Sz int
	 arg_layer int
	 arg_scalex int
	 arg_scaley int
	 argptr [11]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for resize kernel invocation
var resize_args resize_args_t

func init(){
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	 resize_args.argptr[0] = unsafe.Pointer(&resize_args.arg_dst)
	 resize_args.argptr[1] = unsafe.Pointer(&resize_args.arg_Dx)
	 resize_args.argptr[2] = unsafe.Pointer(&resize_args.arg_Dy)
	 resize_args.argptr[3] = unsafe.Pointer(&resize_args.arg_Dz)
	 resize_args.argptr[4] = unsafe.Pointer(&resize_args.arg_src)
	 resize_args.argptr[5] = unsafe.Pointer(&resize_args.arg_Sx)
	 resize_args.argptr[6] = unsafe.Pointer(&resize_args.arg_Sy)
	 resize_args.argptr[7] = unsafe.Pointer(&resize_args.arg_Sz)
	 resize_args.argptr[8] = unsafe.Pointer(&resize_args.arg_layer)
	 resize_args.argptr[9] = unsafe.Pointer(&resize_args.arg_scalex)
	 resize_args.argptr[10] = unsafe.Pointer(&resize_args.arg_scaley)
	 }

// Wrapper for resize CUDA kernel, asynchronous.
func k_resize_async ( dst unsafe.Pointer, Dx int, Dy int, Dz int, src unsafe.Pointer, Sx int, Sy int, Sz int, layer int, scalex int, scaley int,  cfg *config) {
	if Synchronous{ // debug
		Sync()
		timer.Start("resize")
	}

	resize_args.Lock()
	defer resize_args.Unlock()

	if resize_code == 0{
		resize_code = fatbinLoad(resize_map, "resize")
	}

	 resize_args.arg_dst = dst
	 resize_args.arg_Dx = Dx
	 resize_args.arg_Dy = Dy
	 resize_args.arg_Dz = Dz
	 resize_args.arg_src = src
	 resize_args.arg_Sx = Sx
	 resize_args.arg_Sy = Sy
	 resize_args.arg_Sz = Sz
	 resize_args.arg_layer = layer
	 resize_args.arg_scalex = scalex
	 resize_args.arg_scaley = scaley
	

	args := resize_args.argptr[:]
	cu.LaunchKernel(resize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous{ // debug
		Sync()
		timer.Stop("resize")
	}
}

// maps compute capability on PTX code for resize kernel.
var resize_map = map[int]string{ 0: "" ,
52: resize_ptx_52  }

// resize PTX code for various compute capabilities.
const(
  resize_ptx_52 = `
.version 7.0
.target sm_52
.address_size 64

	// .globl	resize

.visible .entry resize(
	.param .u64 resize_param_0,
	.param .u32 resize_param_1,
	.param .u32 resize_param_2,
	.param .u32 resize_param_3,
	.param .u64 resize_param_4,
	.param .u32 resize_param_5,
	.param .u32 resize_param_6,
	.param .u32 resize_param_7,
	.param .u32 resize_param_8,
	.param .u32 resize_param_9,
	.param .u32 resize_param_10
)
{
	.reg .pred 	%p<33>;
	.reg .f32 	%f<85>;
	.reg .b32 	%r<58>;
	.reg .b64 	%rd<18>;


	ld.param.u64 	%rd5, [resize_param_0];
	ld.param.u32 	%r23, [resize_param_1];
	ld.param.u32 	%r29, [resize_param_2];
	ld.param.u64 	%rd6, [resize_param_4];
	ld.param.u32 	%r24, [resize_param_5];
	ld.param.u32 	%r25, [resize_param_6];
	ld.param.u32 	%r26, [resize_param_8];
	ld.param.u32 	%r27, [resize_param_9];
	ld.param.u32 	%r28, [resize_param_10];
	cvta.to.global.u64 	%rd1, %rd6;
	mov.u32 	%r30, %ntid.x;
	mov.u32 	%r31, %ctaid.x;
	mov.u32 	%r32, %tid.x;
	mad.lo.s32 	%r1, %r30, %r31, %r32;
	mov.u32 	%r33, %ntid.y;
	mov.u32 	%r34, %ctaid.y;
	mov.u32 	%r35, %tid.y;
	mad.lo.s32 	%r2, %r33, %r34, %r35;
	setp.lt.s32	%p1, %r1, %r23;
	setp.lt.s32	%p2, %r2, %r29;
	and.pred  	%p3, %p1, %p2;
	@!%p3 bra 	BB0_29;
	bra.uni 	BB0_1;

BB0_1:
	mov.f32 	%f74, 0f00000000;
	setp.lt.s32	%p4, %r28, 1;
	mov.f32 	%f73, %f74;
	@%p4 bra 	BB0_28;

	mul.lo.s32 	%r3, %r2, %r28;
	mul.lo.s32 	%r4, %r1, %r27;
	mul.lo.s32 	%r5, %r26, %r25;
	and.b32  	%r6, %r27, 3;
	mov.f32 	%f74, 0f00000000;
	mov.u32 	%r52, 0;
	mov.f32 	%f73, %f74;

BB0_3:
	add.s32 	%r8, %r52, %r3;
	setp.lt.s32	%p5, %r27, 1;
	@%p5 bra 	BB0_27;

	add.s32 	%r40, %r8, %r5;
	mul.lo.s32 	%r9, %r40, %r24;
	mov.u32 	%r53, 0;
	mov.f32 	%f48, 0f00000000;
	setp.eq.s32	%p6, %r6, 0;
	@%p6 bra 	BB0_5;

	setp.eq.s32	%p7, %r6, 1;
	@%p7 bra 	BB0_13;

	setp.eq.s32	%p8, %r6, 2;
	@%p8 bra 	BB0_10;

	setp.lt.s32	%p9, %r8, %r25;
	setp.lt.s32	%p10, %r4, %r24;
	and.pred  	%p11, %p9, %p10;
	mov.u32 	%r53, 1;
	@!%p11 bra 	BB0_10;
	bra.uni 	BB0_9;

BB0_9:
	add.s32 	%r43, %r4, %r9;
	mul.wide.s32 	%rd7, %r43, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.nc.f32 	%f49, [%rd8];
	add.f32 	%f73, %f73, %f49;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_10:
	add.s32 	%r11, %r53, %r4;
	setp.lt.s32	%p12, %r11, %r24;
	setp.lt.s32	%p13, %r8, %r25;
	and.pred  	%p14, %p13, %p12;
	@!%p14 bra 	BB0_12;
	bra.uni 	BB0_11;

BB0_11:
	add.s32 	%r44, %r11, %r9;
	mul.wide.s32 	%rd9, %r44, 4;
	add.s64 	%rd10, %rd1, %rd9;
	ld.global.nc.f32 	%f50, [%rd10];
	add.f32 	%f73, %f73, %f50;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_12:
	add.s32 	%r53, %r53, 1;

BB0_13:
	add.s32 	%r14, %r53, %r4;
	setp.lt.s32	%p15, %r14, %r24;
	setp.lt.s32	%p16, %r8, %r25;
	and.pred  	%p17, %p16, %p15;
	@!%p17 bra 	BB0_15;
	bra.uni 	BB0_14;

BB0_14:
	add.s32 	%r45, %r14, %r9;
	mul.wide.s32 	%rd11, %r45, 4;
	add.s64 	%rd12, %rd1, %rd11;
	ld.global.nc.f32 	%f51, [%rd12];
	add.f32 	%f73, %f73, %f51;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_15:
	add.s32 	%r53, %r53, 1;
	mov.f32 	%f67, %f74;
	mov.f32 	%f68, %f73;
	bra.uni 	BB0_16;

BB0_5:
	mov.f32 	%f67, %f74;
	mov.f32 	%f68, %f73;
	mov.f32 	%f74, %f48;
	mov.f32 	%f73, %f48;

BB0_16:
	setp.lt.u32	%p18, %r27, 4;
	@%p18 bra 	BB0_27;

	add.s32 	%r56, %r4, %r53;
	mad.lo.s32 	%r47, %r24, %r40, %r56;
	mul.wide.s32 	%rd13, %r47, 4;
	add.s64 	%rd17, %rd1, %rd13;
	mov.f32 	%f74, %f67;
	mov.f32 	%f73, %f68;

BB0_18:
	setp.lt.s32	%p19, %r56, %r24;
	setp.lt.s32	%p20, %r8, %r25;
	and.pred  	%p21, %p20, %p19;
	@!%p21 bra 	BB0_20;
	bra.uni 	BB0_19;

BB0_19:
	ld.global.nc.f32 	%f52, [%rd17];
	add.f32 	%f73, %f73, %f52;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_20:
	add.s32 	%r48, %r56, 1;
	setp.lt.s32	%p22, %r48, %r24;
	and.pred  	%p24, %p20, %p22;
	@!%p24 bra 	BB0_22;
	bra.uni 	BB0_21;

BB0_21:
	ld.global.nc.f32 	%f53, [%rd17+4];
	add.f32 	%f73, %f73, %f53;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_22:
	add.s32 	%r49, %r56, 2;
	setp.lt.s32	%p25, %r49, %r24;
	and.pred  	%p27, %p20, %p25;
	@!%p27 bra 	BB0_24;
	bra.uni 	BB0_23;

BB0_23:
	ld.global.nc.f32 	%f54, [%rd17+8];
	add.f32 	%f73, %f73, %f54;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_24:
	add.s32 	%r50, %r56, 3;
	setp.lt.s32	%p28, %r50, %r24;
	and.pred  	%p30, %p20, %p28;
	@!%p30 bra 	BB0_26;
	bra.uni 	BB0_25;

BB0_25:
	ld.global.nc.f32 	%f55, [%rd17+12];
	add.f32 	%f73, %f73, %f55;
	add.f32 	%f74, %f74, 0f3F800000;

BB0_26:
	add.s32 	%r56, %r56, 4;
	add.s32 	%r53, %r53, 4;
	setp.lt.s32	%p31, %r53, %r27;
	add.s64 	%rd17, %rd17, 16;
	@%p31 bra 	BB0_18;

BB0_27:
	add.s32 	%r52, %r52, 1;
	setp.lt.s32	%p32, %r52, %r28;
	@%p32 bra 	BB0_3;

BB0_28:
	cvta.to.global.u64 	%rd14, %rd5;
	mad.lo.s32 	%r51, %r2, %r23, %r1;
	mul.wide.s32 	%rd15, %r51, 4;
	add.s64 	%rd16, %rd14, %rd15;
	div.rn.f32 	%f56, %f73, %f74;
	st.global.f32 	[%rd16], %f56;

BB0_29:
	ret;
}


`
 )
