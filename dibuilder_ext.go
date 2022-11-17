package llvm

/*
#include "IRBindings.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type DIGlobalVariable struct {
	Name        string
	Linkage     string
	File        Metadata
	Line        int
	Type        Metadata
	LocalToUnit bool
	Expr        Metadata
	Decl        Metadata
	AlignInBits uint32
}

func (d *DIBuilder) CreateGlobalVariable(scope Metadata, v DIGlobalVariable) Metadata {
	name := C.CString(v.Name)
	linkage := C.CString(v.Linkage)
	defer C.free(unsafe.Pointer(name))
	result := C.LLVMDIBuilderCreateGlobalVariableExpression(
		d.ref,
		scope.C,
		name, C.size_t(len(v.Name)),
		linkage, C.size_t(len(v.Linkage)),
		v.File.C,
		C.unsigned(v.Line),
		v.Type.C,
		C.LLVMBool(boolToCInt(v.LocalToUnit)),
		v.Expr.C,
		v.Decl.C,
		C.uint32_t(v.AlignInBits))
	return Metadata{C: result}
}
