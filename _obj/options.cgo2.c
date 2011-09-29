#line 3 "options.go"
#include <tidy.h>
#include <buffio.h>
#include <errno.h>



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_5047e4bdfdff_Cfunc_tidyOptSetValue(void *v)
{
	struct {
		struct _TidyDoc* p0;
		TidyOptionId p1;
		char __pad12[4];
		tmbchar* p2;
		Bool r;
		char __pad28[4];
	} __attribute__((__packed__)) *a = v;
	a->r = tidyOptSetValue(a->p0, a->p1, a->p2);
}

void
_cgo_5047e4bdfdff_Cfunc_tidyOptSetInt(void *v)
{
	struct {
		struct _TidyDoc* p0;
		TidyOptionId p1;
		char __pad12[4];
		ulong p2;
		Bool r;
		char __pad28[4];
	} __attribute__((__packed__)) *a = v;
	a->r = tidyOptSetInt(a->p0, a->p1, a->p2);
}

