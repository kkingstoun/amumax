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

// CUDA handle for madd2 kernel
var madd2_code cu.Function

// Stores the arguments for madd2 kernel invocation
type madd2_args_t struct{
	 arg_dst unsafe.Pointer
	 arg_src1 unsafe.Pointer
	 arg_fac1 float32
	 arg_src2 unsafe.Pointer
	 arg_fac2 float32
	 arg_N int
	 argptr [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for madd2 kernel invocation
var madd2_args madd2_args_t

func init(){
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	 madd2_args.argptr[0] = unsafe.Pointer(&madd2_args.arg_dst)
	 madd2_args.argptr[1] = unsafe.Pointer(&madd2_args.arg_src1)
	 madd2_args.argptr[2] = unsafe.Pointer(&madd2_args.arg_fac1)
	 madd2_args.argptr[3] = unsafe.Pointer(&madd2_args.arg_src2)
	 madd2_args.argptr[4] = unsafe.Pointer(&madd2_args.arg_fac2)
	 madd2_args.argptr[5] = unsafe.Pointer(&madd2_args.arg_N)
	 }

// Wrapper for madd2 CUDA kernel, asynchronous.
func k_madd2_async ( dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, N int,  cfg *config) {
	if Synchronous{ // debug
		Sync()
		timer.Start("madd2")
	}

	madd2_args.Lock()
	defer madd2_args.Unlock()

	if madd2_code == 0{
		madd2_code = fatbinLoad(madd2_map, "madd2")
	}

	 madd2_args.arg_dst = dst
	 madd2_args.arg_src1 = src1
	 madd2_args.arg_fac1 = fac1
	 madd2_args.arg_src2 = src2
	 madd2_args.arg_fac2 = fac2
	 madd2_args.arg_N = N
	

	args := madd2_args.argptr[:]
	cu.LaunchKernel(madd2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous{ // debug
		Sync()
		timer.Stop("madd2")
	}
}

// maps compute capability on PTX code for madd2 kernel.
var madd2_map = map[int]string{ 0: "" ,
52: madd2_ptx_52  }

// madd2 PTX code for various compute capabilities.
const(
  madd2_ptx_52 = `
.version 7.0
.target sm_52
.address_size 64

	// .globl	madd2

.visible .entry madd2(
	.param .u64 madd2_param_0,
	.param .u64 madd2_param_1,
	.param .f32 madd2_param_2,
	.param .u64 madd2_param_3,
	.param .f32 madd2_param_4,
	.param .u32 madd2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<7>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [madd2_param_0];
	ld.param.u64 	%rd2, [madd2_param_1];
	ld.param.f32 	%f1, [madd2_param_2];
	ld.param.u64 	%rd3, [madd2_param_3];
	ld.param.f32 	%f2, [madd2_param_4];
	ld.param.u32 	%r2, [madd2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f4, [%rd8];
	mul.f32 	%f5, %f4, %f2;
	fma.rn.f32 	%f6, %f3, %f1, %f5;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f6;

BB0_2:
	ret;
}


`
 )
