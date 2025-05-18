package wit

import (
	"slices"
	"strconv"
)

// ABI is the interface implemented by any type that can report its
// [Canonical ABI] [size], [alignment], and [flat] representation.
//
// [Canonical ABI]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md
// [size]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#element-size
// [alignment]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#alignment
// [flat]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
type ABI interface {
	Size() uintptr
	Align() uintptr
	Flat() []Type
}

// Align aligns ptr with alignment align.
func Align(ptr, align uintptr) uintptr {
	return (ptr + align - 1) &^ (align - 1)
}

// Discriminant returns the smallest WIT integer type that can represent 0...n.
// Used by the [Canonical ABI] for [Variant] types.
//
// [Canonical ABI]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#alignment
func Discriminant(n int) Type {
	switch {
	case n <= 1<<8:
		return U8{}
	case n <= 1<<16:
		return U16{}
	}
	return U32{}
}

// Despecialize [despecializes] k if k can be despecialized. Otherwise, it returns k unmodified.
// See the [canonical ABI documentation] for more information.
//
// [despecializes]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#despecialization
// [canonical ABI documentation]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#despecialization
func Despecialize(k TypeDefKind) TypeDefKind {
	if d, ok := k.(interface{ Despecialize() TypeDefKind }); ok {
		return d.Despecialize()
	}
	return k
}

// HasPointer returns whether or not t contains a [Type] with a pointer, e.g. [String] or [List].
func HasPointer(t TypeDefKind) bool {
	t = Despecialize(t)
	if p, ok := t.(interface{ hasPointer() bool }); ok {
		return p.hasPointer()
	}
	return false
}

// HasResource returns whether or not t contains a resource type, typically an [Own] or [Borrow] handle.
func HasResource(t TypeDefKind) bool {
	t = Despecialize(t)
	if p, ok := t.(interface{ hasResource() bool }); ok {
		return p.hasResource()
	}
	return false
}

// HasBorrow returns whether or not t contains a [Borrow] type.
func HasBorrow(t TypeDefKind) bool {
	t = Despecialize(t)
	if p, ok := t.(interface{ hasBorrow() bool }); ok {
		return p.hasBorrow()
	}
	return false
}

// LowerFunction returns a [Function] signature for lowering [Type] t.
func LowerFunction(t Type) *Function {
	return &Function{
		Name:    "[lower]" + t.TypeName(),
		Kind:    &Freestanding{},
		Params:  []Param{{Name: "v", Type: t}},
		Results: flatParams("f", t.Flat()),
	}
}

// LiftFunction returns a [Function] signature for lifting [Type] t.
func LiftFunction(t Type) *Function {
	return &Function{
		Name:    "[lift]" + t.TypeName(),
		Kind:    &Freestanding{},
		Params:  flatParams("f", t.Flat()),
		Results: []Param{{Name: "v", Type: t}},
	}
}

// Direction represents the direction a type or function is represented within a component,
// whether it is an importer (consumer), or an exporter (producer). When applied to functions,
// this represents the [Canonical ABI] [lift] and [lower] operations, for lowering into or lifting out of linear memory.
//
// [Canonical ABI]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/Explainer.md#canonical-abi
// [lift]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lift
// [lower]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lower
type Direction int

// String implements the Stringer interface.
func (dir Direction) String() string {
	switch dir {
	case Imported:
		return "imported"
	case Exported:
		return "exported"
	default:
		return strconv.Itoa(int(dir))
	}
}

const (
	// Exported represents types and functions imported into a component from the host or another component.
	// This corresponds to the the Canonical ABI [lower] operation, lowering Component Model types into linear memory.
	// Used for calling functions imported using //go:wasmimport.
	//
	// [lower]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lower
	Imported Direction = 0

	// Exported represents types and functions exported from a component to the host or another component.
	// This corresponds to the Canonical ABI [lift] operation, lifting Component Model types out of linear memory.
	// Used for exporting functions using //go:wasmexport.
	//
	// [lift]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lift
	Exported Direction = 1
)

// ResourceNew returns the implied [resource-new] function for t.
// If t is not a [Resource], this returns nil.
//
// [resource-new]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-resourcenew
func (t *TypeDef) ResourceNew() *Function {
	if _, ok := t.Kind.(*Resource); !ok {
		return nil
	}
	return &Function{
		Name:    "[resource-new]" + t.TypeName(),
		Kind:    &Static{Type: t},
		Params:  []Param{{Name: "rep", Type: &TypeDef{Kind: &Borrow{Type: t}}}},
		Results: []Param{{Type: t}},
		Docs:    Docs{Contents: "Creates a new resource handle."},
	}
}

// ResourceRep returns the implied [resource-rep] method for t.
// If t is not a [Resource], this returns nil.
//
// [resource-rep]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-resourcerep
func (t *TypeDef) ResourceRep() *Function {
	if _, ok := t.Kind.(*Resource); !ok {
		return nil
	}
	return &Function{
		Name:    "[resource-rep]" + t.TypeName(),
		Kind:    &Method{Type: t},
		Params:  []Param{{Name: "self", Type: t}},
		Results: []Param{{Type: &TypeDef{Kind: &Borrow{Type: t}}}},
		Docs:    Docs{Contents: "Returns the underlying resource representation."},
	}
}

// ResourceDrop returns the implied [resource-drop] method for t.
// If t is not a [Resource], this returns nil.
//
// [resource-drop]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-resourcedrop
func (t *TypeDef) ResourceDrop() *Function {
	if _, ok := t.Kind.(*Resource); !ok {
		return nil
	}
	return &Function{
		Name:   "[resource-drop]" + t.TypeName(),
		Kind:   &Method{Type: t},
		Params: []Param{{Name: "self", Type: t}},
		Docs:   Docs{Contents: "Drops a resource handle."},
	}
}

// Destructor returns the implied destructor ([dtor]) method for t.
// If t is not a [Resource], this returns nil.
//
// [dtor]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-resourcedrop
func (t *TypeDef) Destructor() *Function {
	if _, ok := t.Kind.(*Resource); !ok {
		return nil
	}
	return &Function{
		Name:   "[dtor]" + t.TypeName(),
		Kind:   &Method{Type: t},
		Params: []Param{{Name: "self", Type: &TypeDef{Kind: &Borrow{Type: t}}}},
		Docs:   Docs{Contents: "Resource destructor."},
	}
}

// PostReturn returns a [post-return] function for f, which is part of the
// Component Model machinery that allows the caller of f to call back into
// the component to clean up results. Returns nil if the Core WebAssembly
// derivative of f has no results, therefore does not require cleanup.
//
// While this accepts a [Direction], this is currently only used for exported functions.
//
// [post-return]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lift
func (f *Function) PostReturn(dir Direction) *Function {
	core := f.CoreFunction(dir)

	if !core.ReturnsPointer() {
		return nil
	}

	params := slices.Clone(core.Results)
	if params[0].Name == "" {
		params[0].Name = "result"
	}

	return &Function{
		Name:   "cabi_post_" + f.Name,
		Kind:   &Freestanding{},
		Params: params,
		Docs:   Docs{Contents: "Post-return cleanup function."},
	}
}

// ReturnsBorrow reports whether [Function] f returns a [Borrow] handle,
// which is not permitted by the Component Model specification.
func (f *Function) ReturnsBorrow() bool {
	for _, r := range f.Results {
		if HasBorrow(r.Type) {
			return true
		}
	}
	return false
}

// ReturnsPointer reports whether [Function] f returns a value containing a pointer,
// which would require a post-return cleanup function if exported.
func (f *Function) ReturnsPointer() bool {
	for _, r := range f.Results {
		if HasPointer(r.Type) {
			return true
		}
	}
	return false
}

const (
	// MaxFlatParams is the maximum number of [flattened parameters] a function can have
	// as defined in the Component Model Canonical ABI.
	//
	// [flattened parameters]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
	MaxFlatParams = 16

	// MaxFlatResults is the maximum number of [flattened results] a function can have
	// as defined in the Component Model Canonical ABI.
	//
	// [flattened results]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
	MaxFlatResults = 1
)

// CoreFunction returns a [Core WebAssembly function] of [Function] f.
// Its params and results may be [flattened] according to the Canonical ABI specification.
// The flattening rules vary based on whether the returned function is imported or exported,
// e.g. using go:wasmimport or go:wasmexport.
//
// [Core WebAssembly function]: https://webassembly.github.io/spec/core/syntax/modules.html#syntax-func
// [flattened]: https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#flattening
func (f *Function) CoreFunction(op Direction) *Function {
	if len(f.Params) == 0 && len(f.Results) == 0 {
		return f
	}

	// Clone the function
	cf := *f

	// Max 16 params
	cf.Params = flattenParams(f.Params)
	if len(cf.Params) > MaxFlatParams {
		cf.Params = []Param{compoundParam("param", "params", f.Params)}
	}

	// Max 1 result
	cf.Results = flattenParams(f.Results)
	if len(cf.Results) > MaxFlatResults {
		p := compoundParam("result", "results", f.Results)
		if op == Exported {
			cf.Results = []Param{p}
		} else {
			cf.Params = append(cf.Params, p)
			cf.Results = nil
		}
	}

	return &cf
}

func flatParams(pfx string, flat []Type) []Param {
	out := make([]Param, len(flat))
	for i, t := range flat {
		out[i] = Param{Name: pfx + strconv.Itoa(i), Type: t}
	}
	return out
}

func flattenParams(params []Param) []Param {
	var out []Param
	for _, p := range params {
		flat := p.Type.Flat()
		if len(flat) == 1 {
			if p.Name == "" {
				p.Name = "result"
			}
			out = append(out, Param{Name: p.Name + "0", Type: flat[0]})
		} else {
			for i, t := range flat {
				out = append(out, Param{Name: p.Name + strconv.Itoa(i), Type: t})
			}
		}
	}
	return out
}

// compoundParam returns a single param that represents
// the combined param(s), using a [Pointer].
func compoundParam(singular, plural string, params []Param) Param {
	if len(params) == 0 {
		panic("BUG: len(params) == 0")
	}

	name := params[0].Name
	var t Type

	if len(params) == 1 {
		if name == "" {
			name = singular
		}
		t = params[0].Type
	} else {
		name = plural
		r := &Record{}
		t = &TypeDef{Kind: r}
		for _, p := range params {
			r.Fields = append(r.Fields,
				Field{
					Name: p.Name,
					Type: p.Type,
				})
		}
	}

	return Param{
		Name: name,
		Type: PointerTo(t),
	}
}
