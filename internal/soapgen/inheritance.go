package soapgen

import "github.com/enuan/soap-go/xsd"

// flattenSchemaInheritance rewrites complexContent/extension complex types
// in-place so that every derived type carries the fields of its ancestors as a
// plain content model.
//
// encoding/xml (and this generator) operate on flat Go structs, so inherited
// fields must be merged into the derived struct. Without this step, fields
// declared in a base type (e.g. a session token in BaseRequest or the result
// state in ResultBase) are silently dropped from the generated code.
//
// Base resolution uses local names, matching the rest of the generator; cycles
// (invalid XSD, but seen in the wild) are broken instead of recursing forever.
func flattenSchemaInheritance(schema *xsd.Schema) {
	if schema == nil {
		return
	}

	byName := make(map[string]*xsd.ComplexType, len(schema.ComplexTypes))
	for i := range schema.ComplexTypes {
		ct := &schema.ComplexTypes[i]
		if ct.Name != "" {
			byName[ct.Name] = ct
		}
	}

	resolving := make(map[string]bool)

	var flatten func(ct *xsd.ComplexType)
	flatten = func(ct *xsd.ComplexType) {
		if ct == nil || ct.ComplexContent == nil || ct.ComplexContent.Extension == nil {
			return
		}
		ext := ct.ComplexContent.Extension

		// Collect inherited content from the base type, if resolvable.
		var inheritedElems []xsd.Element
		var inheritedAnys []xsd.Any
		var inheritedAttrs []xsd.Attribute
		if baseName := extractLocalName(ext.Base); baseName != "" {
			if base := byName[baseName]; base != nil && base != ct && !resolving[baseName] {
				resolving[baseName] = true
				flatten(base)
				delete(resolving, baseName)

				baseElems, baseAnys := contentModelChildren(base)
				inheritedElems = append(inheritedElems, baseElems...)
				inheritedAnys = append(inheritedAnys, baseAnys...)
				inheritedAttrs = append(inheritedAttrs, base.Attributes...)
			}
		}

		ownElems, ownAnys := extensionContentModelChildren(ext)

		ct.Sequence = &xsd.Sequence{
			Elements: append(inheritedElems, ownElems...),
			Any:      append(inheritedAnys, ownAnys...),
		}
		ct.Attributes = append(inheritedAttrs, ext.Attributes...)
		ct.ComplexContent = nil
	}

	for i := range schema.ComplexTypes {
		flatten(&schema.ComplexTypes[i])
		flattenInlineComplexTypes(&schema.ComplexTypes[i], flatten)
	}
	for i := range schema.Elements {
		flattenInlineComplexTypes(schema.Elements[i].ComplexType, flatten)
	}
}

// flattenInlineComplexTypes applies flatten to anonymous complex types nested
// in the content model of ct (elements can carry inline anonymous types that
// themselves use complexContent/extension).
func flattenInlineComplexTypes(ct *xsd.ComplexType, flatten func(*xsd.ComplexType)) {
	if ct == nil {
		return
	}
	var walk func(elements []xsd.Element)
	walk = func(elements []xsd.Element) {
		for i := range elements {
			inline := elements[i].ComplexType
			if inline == nil {
				continue
			}
			flatten(inline)
			flattenInlineComplexTypes(inline, flatten)
		}
	}
	if ct.Sequence != nil {
		walk(ct.Sequence.Elements)
	}
	if ct.Choice != nil {
		walk(ct.Choice.Elements)
	}
	if ct.All != nil {
		walk(ct.All.Elements)
	}
}
