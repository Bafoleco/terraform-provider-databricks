// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	dataschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ListNestedAttributteBuilder represents a list of complex (non-primitive) types.
type ListNestedAttributeBuilder struct {
	NestedObject       NestedAttributeObject
	Optional           bool
	Required           bool
	Sensitive          bool
	Computed           bool
	DeprecationMessage string
	Validators         []validator.List
	PlanModifiers      []planmodifier.List
}

func (a ListNestedAttributeBuilder) BuildDataSourceAttribute() dataschema.Attribute {
	return dataschema.ListNestedAttribute{
		NestedObject:       a.NestedObject.BuildDataSourceAttribute(),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
	}
}

func (a ListNestedAttributeBuilder) BuildResourceAttribute() schema.Attribute {
	return schema.ListNestedAttribute{
		NestedObject:       a.NestedObject.BuildResourceAttribute(),
		Optional:           a.Optional,
		Required:           a.Required,
		Sensitive:          a.Sensitive,
		DeprecationMessage: a.DeprecationMessage,
		Computed:           a.Computed,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}

func (a ListNestedAttributeBuilder) SetOptional() AttributeBuilder {
	if a.Optional && !a.Required {
		panic("attribute is already optional")
	}
	a.Optional = true
	a.Required = false
	return a
}

func (a ListNestedAttributeBuilder) SetRequired() AttributeBuilder {
	if !a.Optional && a.Required {
		panic("attribute is already required")
	}
	a.Optional = false
	a.Required = true
	return a
}

func (a ListNestedAttributeBuilder) SetSensitive() AttributeBuilder {
	if a.Sensitive {
		panic("attribute is already sensitive")
	}
	a.Sensitive = true
	return a
}

func (a ListNestedAttributeBuilder) SetComputed() AttributeBuilder {
	if a.Computed {
		panic("attribute is already computed")
	}
	a.Computed = true
	return a
}

func (a ListNestedAttributeBuilder) SetReadOnly() AttributeBuilder {
	if a.Computed && !a.Optional && !a.Required {
		panic("attribute is already read only")
	}
	a.Computed = true
	a.Optional = false
	a.Required = false
	a.NestedObject.SetReadOnly()
	return a
}

func (a ListNestedAttributeBuilder) SetDeprecated(msg string) BaseSchemaBuilder {
	a.DeprecationMessage = msg
	return a
}

func (a ListNestedAttributeBuilder) AddValidator(v validator.List) BaseSchemaBuilder {
	a.Validators = append(a.Validators, v)
	return a
}

func (a ListNestedAttributeBuilder) AddPlanModifier(v planmodifier.List) BaseSchemaBuilder {
	a.PlanModifiers = append(a.PlanModifiers, v)
	return a
}

func (a ListNestedAttributeBuilder) ToBlock() BlockBuilder {
	return ListNestedBlockBuilder{
		NestedObject:       convertAttributesToBlocks(a.NestedObject.Attributes, nil),
		DeprecationMessage: a.DeprecationMessage,
		Validators:         a.Validators,
		PlanModifiers:      a.PlanModifiers,
	}
}
