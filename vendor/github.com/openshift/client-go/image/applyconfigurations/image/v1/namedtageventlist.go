// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NamedTagEventListApplyConfiguration represents an declarative configuration of the NamedTagEventList type for use
// with apply.
type NamedTagEventListApplyConfiguration struct {
	Tag        *string                               `json:"tag,omitempty"`
	Items      []TagEventApplyConfiguration          `json:"items,omitempty"`
	Conditions []TagEventConditionApplyConfiguration `json:"conditions,omitempty"`
}

// NamedTagEventListApplyConfiguration constructs an declarative configuration of the NamedTagEventList type for use with
// apply.
func NamedTagEventList() *NamedTagEventListApplyConfiguration {
	return &NamedTagEventListApplyConfiguration{}
}

// WithTag sets the Tag field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tag field is set to the value of the last call.
func (b *NamedTagEventListApplyConfiguration) WithTag(value string) *NamedTagEventListApplyConfiguration {
	b.Tag = &value
	return b
}

// WithItems adds the given value to the Items field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Items field.
func (b *NamedTagEventListApplyConfiguration) WithItems(values ...*TagEventApplyConfiguration) *NamedTagEventListApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithItems")
		}
		b.Items = append(b.Items, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *NamedTagEventListApplyConfiguration) WithConditions(values ...*TagEventConditionApplyConfiguration) *NamedTagEventListApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}