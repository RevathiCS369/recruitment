// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recruit/ent/divisionmaster"
	"recruit/ent/regionmaster"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DivisionMasterCreate is the builder for creating a DivisionMaster entity.
type DivisionMasterCreate struct {
	config
	mutation *DivisionMasterMutation
	hooks    []Hook
}

// SetDivisionCode sets the "DivisionCode" field.
func (dmc *DivisionMasterCreate) SetDivisionCode(i int32) *DivisionMasterCreate {
	dmc.mutation.SetDivisionCode(i)
	return dmc
}

// SetOfficeType sets the "OfficeType" field.
func (dmc *DivisionMasterCreate) SetOfficeType(s string) *DivisionMasterCreate {
	dmc.mutation.SetOfficeType(s)
	return dmc
}

// SetDivisionOfficeID sets the "DivisionOfficeID" field.
func (dmc *DivisionMasterCreate) SetDivisionOfficeID(s string) *DivisionMasterCreate {
	dmc.mutation.SetDivisionOfficeID(s)
	return dmc
}

// SetDivisionOfficeName sets the "DivisionOfficeName" field.
func (dmc *DivisionMasterCreate) SetDivisionOfficeName(s string) *DivisionMasterCreate {
	dmc.mutation.SetDivisionOfficeName(s)
	return dmc
}

// SetReportingOfficeType sets the "ReportingOfficeType" field.
func (dmc *DivisionMasterCreate) SetReportingOfficeType(s string) *DivisionMasterCreate {
	dmc.mutation.SetReportingOfficeType(s)
	return dmc
}

// SetNillableReportingOfficeType sets the "ReportingOfficeType" field if the given value is not nil.
func (dmc *DivisionMasterCreate) SetNillableReportingOfficeType(s *string) *DivisionMasterCreate {
	if s != nil {
		dmc.SetReportingOfficeType(*s)
	}
	return dmc
}

// SetReportingOfficeCode sets the "ReportingOfficeCode" field.
func (dmc *DivisionMasterCreate) SetReportingOfficeCode(s string) *DivisionMasterCreate {
	dmc.mutation.SetReportingOfficeCode(s)
	return dmc
}

// SetNillableReportingOfficeCode sets the "ReportingOfficeCode" field if the given value is not nil.
func (dmc *DivisionMasterCreate) SetNillableReportingOfficeCode(s *string) *DivisionMasterCreate {
	if s != nil {
		dmc.SetReportingOfficeCode(*s)
	}
	return dmc
}

// SetEmailID sets the "EmailID" field.
func (dmc *DivisionMasterCreate) SetEmailID(s string) *DivisionMasterCreate {
	dmc.mutation.SetEmailID(s)
	return dmc
}

// SetNillableEmailID sets the "EmailID" field if the given value is not nil.
func (dmc *DivisionMasterCreate) SetNillableEmailID(s *string) *DivisionMasterCreate {
	if s != nil {
		dmc.SetEmailID(*s)
	}
	return dmc
}

// SetMobileNumber sets the "MobileNumber" field.
func (dmc *DivisionMasterCreate) SetMobileNumber(i int32) *DivisionMasterCreate {
	dmc.mutation.SetMobileNumber(i)
	return dmc
}

// SetNillableMobileNumber sets the "MobileNumber" field if the given value is not nil.
func (dmc *DivisionMasterCreate) SetNillableMobileNumber(i *int32) *DivisionMasterCreate {
	if i != nil {
		dmc.SetMobileNumber(*i)
	}
	return dmc
}

// SetRegionCode sets the "RegionCode" field.
func (dmc *DivisionMasterCreate) SetRegionCode(i int32) *DivisionMasterCreate {
	dmc.mutation.SetRegionCode(i)
	return dmc
}

// SetNillableRegionCode sets the "RegionCode" field if the given value is not nil.
func (dmc *DivisionMasterCreate) SetNillableRegionCode(i *int32) *DivisionMasterCreate {
	if i != nil {
		dmc.SetRegionCode(*i)
	}
	return dmc
}

// SetID sets the "id" field.
func (dmc *DivisionMasterCreate) SetID(i int32) *DivisionMasterCreate {
	dmc.mutation.SetID(i)
	return dmc
}

// AddRegionIDs adds the "regions" edge to the RegionMaster entity by IDs.
func (dmc *DivisionMasterCreate) AddRegionIDs(ids ...int32) *DivisionMasterCreate {
	dmc.mutation.AddRegionIDs(ids...)
	return dmc
}

// AddRegions adds the "regions" edges to the RegionMaster entity.
func (dmc *DivisionMasterCreate) AddRegions(r ...*RegionMaster) *DivisionMasterCreate {
	ids := make([]int32, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dmc.AddRegionIDs(ids...)
}

// Mutation returns the DivisionMasterMutation object of the builder.
func (dmc *DivisionMasterCreate) Mutation() *DivisionMasterMutation {
	return dmc.mutation
}

// Save creates the DivisionMaster in the database.
func (dmc *DivisionMasterCreate) Save(ctx context.Context) (*DivisionMaster, error) {
	return withHooks(ctx, dmc.sqlSave, dmc.mutation, dmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DivisionMasterCreate) SaveX(ctx context.Context) *DivisionMaster {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DivisionMasterCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DivisionMasterCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DivisionMasterCreate) check() error {
	if _, ok := dmc.mutation.DivisionCode(); !ok {
		return &ValidationError{Name: "DivisionCode", err: errors.New(`ent: missing required field "DivisionMaster.DivisionCode"`)}
	}
	if _, ok := dmc.mutation.OfficeType(); !ok {
		return &ValidationError{Name: "OfficeType", err: errors.New(`ent: missing required field "DivisionMaster.OfficeType"`)}
	}
	if _, ok := dmc.mutation.DivisionOfficeID(); !ok {
		return &ValidationError{Name: "DivisionOfficeID", err: errors.New(`ent: missing required field "DivisionMaster.DivisionOfficeID"`)}
	}
	if _, ok := dmc.mutation.DivisionOfficeName(); !ok {
		return &ValidationError{Name: "DivisionOfficeName", err: errors.New(`ent: missing required field "DivisionMaster.DivisionOfficeName"`)}
	}
	return nil
}

func (dmc *DivisionMasterCreate) sqlSave(ctx context.Context) (*DivisionMaster, error) {
	if err := dmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	dmc.mutation.id = &_node.ID
	dmc.mutation.done = true
	return _node, nil
}

func (dmc *DivisionMasterCreate) createSpec() (*DivisionMaster, *sqlgraph.CreateSpec) {
	var (
		_node = &DivisionMaster{config: dmc.config}
		_spec = sqlgraph.NewCreateSpec(divisionmaster.Table, sqlgraph.NewFieldSpec(divisionmaster.FieldID, field.TypeInt32))
	)
	if id, ok := dmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dmc.mutation.DivisionCode(); ok {
		_spec.SetField(divisionmaster.FieldDivisionCode, field.TypeInt32, value)
		_node.DivisionCode = value
	}
	if value, ok := dmc.mutation.OfficeType(); ok {
		_spec.SetField(divisionmaster.FieldOfficeType, field.TypeString, value)
		_node.OfficeType = value
	}
	if value, ok := dmc.mutation.DivisionOfficeID(); ok {
		_spec.SetField(divisionmaster.FieldDivisionOfficeID, field.TypeString, value)
		_node.DivisionOfficeID = value
	}
	if value, ok := dmc.mutation.DivisionOfficeName(); ok {
		_spec.SetField(divisionmaster.FieldDivisionOfficeName, field.TypeString, value)
		_node.DivisionOfficeName = value
	}
	if value, ok := dmc.mutation.ReportingOfficeType(); ok {
		_spec.SetField(divisionmaster.FieldReportingOfficeType, field.TypeString, value)
		_node.ReportingOfficeType = value
	}
	if value, ok := dmc.mutation.ReportingOfficeCode(); ok {
		_spec.SetField(divisionmaster.FieldReportingOfficeCode, field.TypeString, value)
		_node.ReportingOfficeCode = value
	}
	if value, ok := dmc.mutation.EmailID(); ok {
		_spec.SetField(divisionmaster.FieldEmailID, field.TypeString, value)
		_node.EmailID = value
	}
	if value, ok := dmc.mutation.MobileNumber(); ok {
		_spec.SetField(divisionmaster.FieldMobileNumber, field.TypeInt32, value)
		_node.MobileNumber = value
	}
	if value, ok := dmc.mutation.RegionCode(); ok {
		_spec.SetField(divisionmaster.FieldRegionCode, field.TypeInt32, value)
		_node.RegionCode = value
	}
	if nodes := dmc.mutation.RegionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   divisionmaster.RegionsTable,
			Columns: []string{divisionmaster.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(regionmaster.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DivisionMasterCreateBulk is the builder for creating many DivisionMaster entities in bulk.
type DivisionMasterCreateBulk struct {
	config
	builders []*DivisionMasterCreate
}

// Save creates the DivisionMaster entities in the database.
func (dmcb *DivisionMasterCreateBulk) Save(ctx context.Context) ([]*DivisionMaster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DivisionMaster, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DivisionMasterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DivisionMasterCreateBulk) SaveX(ctx context.Context) []*DivisionMaster {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DivisionMasterCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DivisionMasterCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
