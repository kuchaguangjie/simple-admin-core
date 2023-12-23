// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/rpc/ent/dictionarydetail"
)

// DictionaryDetailCreate is the builder for creating a DictionaryDetail entity.
type DictionaryDetailCreate struct {
	config
	mutation *DictionaryDetailMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ddc *DictionaryDetailCreate) SetCreatedAt(t time.Time) *DictionaryDetailCreate {
	ddc.mutation.SetCreatedAt(t)
	return ddc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableCreatedAt(t *time.Time) *DictionaryDetailCreate {
	if t != nil {
		ddc.SetCreatedAt(*t)
	}
	return ddc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddc *DictionaryDetailCreate) SetUpdatedAt(t time.Time) *DictionaryDetailCreate {
	ddc.mutation.SetUpdatedAt(t)
	return ddc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableUpdatedAt(t *time.Time) *DictionaryDetailCreate {
	if t != nil {
		ddc.SetUpdatedAt(*t)
	}
	return ddc
}

// SetStatus sets the "status" field.
func (ddc *DictionaryDetailCreate) SetStatus(u uint8) *DictionaryDetailCreate {
	ddc.mutation.SetStatus(u)
	return ddc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableStatus(u *uint8) *DictionaryDetailCreate {
	if u != nil {
		ddc.SetStatus(*u)
	}
	return ddc
}

// SetSort sets the "sort" field.
func (ddc *DictionaryDetailCreate) SetSort(u uint32) *DictionaryDetailCreate {
	ddc.mutation.SetSort(u)
	return ddc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableSort(u *uint32) *DictionaryDetailCreate {
	if u != nil {
		ddc.SetSort(*u)
	}
	return ddc
}

// SetTitle sets the "title" field.
func (ddc *DictionaryDetailCreate) SetTitle(s string) *DictionaryDetailCreate {
	ddc.mutation.SetTitle(s)
	return ddc
}

// SetKey sets the "key" field.
func (ddc *DictionaryDetailCreate) SetKey(s string) *DictionaryDetailCreate {
	ddc.mutation.SetKey(s)
	return ddc
}

// SetValue sets the "value" field.
func (ddc *DictionaryDetailCreate) SetValue(s string) *DictionaryDetailCreate {
	ddc.mutation.SetValue(s)
	return ddc
}

// SetDictionaryID sets the "dictionary_id" field.
func (ddc *DictionaryDetailCreate) SetDictionaryID(u uint64) *DictionaryDetailCreate {
	ddc.mutation.SetDictionaryID(u)
	return ddc
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableDictionaryID(u *uint64) *DictionaryDetailCreate {
	if u != nil {
		ddc.SetDictionaryID(*u)
	}
	return ddc
}

// SetID sets the "id" field.
func (ddc *DictionaryDetailCreate) SetID(u uint64) *DictionaryDetailCreate {
	ddc.mutation.SetID(u)
	return ddc
}

// SetDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID.
func (ddc *DictionaryDetailCreate) SetDictionariesID(id uint64) *DictionaryDetailCreate {
	ddc.mutation.SetDictionariesID(id)
	return ddc
}

// SetNillableDictionariesID sets the "dictionaries" edge to the Dictionary entity by ID if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableDictionariesID(id *uint64) *DictionaryDetailCreate {
	if id != nil {
		ddc = ddc.SetDictionariesID(*id)
	}
	return ddc
}

// SetDictionaries sets the "dictionaries" edge to the Dictionary entity.
func (ddc *DictionaryDetailCreate) SetDictionaries(d *Dictionary) *DictionaryDetailCreate {
	return ddc.SetDictionariesID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (ddc *DictionaryDetailCreate) Mutation() *DictionaryDetailMutation {
	return ddc.mutation
}

// Save creates the DictionaryDetail in the database.
func (ddc *DictionaryDetailCreate) Save(ctx context.Context) (*DictionaryDetail, error) {
	ddc.defaults()
	return withHooks(ctx, ddc.sqlSave, ddc.mutation, ddc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddc *DictionaryDetailCreate) SaveX(ctx context.Context) *DictionaryDetail {
	v, err := ddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddc *DictionaryDetailCreate) Exec(ctx context.Context) error {
	_, err := ddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddc *DictionaryDetailCreate) ExecX(ctx context.Context) {
	if err := ddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddc *DictionaryDetailCreate) defaults() {
	if _, ok := ddc.mutation.CreatedAt(); !ok {
		v := dictionarydetail.DefaultCreatedAt()
		ddc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddc.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.DefaultUpdatedAt()
		ddc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ddc.mutation.Status(); !ok {
		v := dictionarydetail.DefaultStatus
		ddc.mutation.SetStatus(v)
	}
	if _, ok := ddc.mutation.Sort(); !ok {
		v := dictionarydetail.DefaultSort
		ddc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddc *DictionaryDetailCreate) check() error {
	if _, ok := ddc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DictionaryDetail.created_at"`)}
	}
	if _, ok := ddc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DictionaryDetail.updated_at"`)}
	}
	if _, ok := ddc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "DictionaryDetail.sort"`)}
	}
	if _, ok := ddc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "DictionaryDetail.title"`)}
	}
	if _, ok := ddc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "DictionaryDetail.key"`)}
	}
	if _, ok := ddc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "DictionaryDetail.value"`)}
	}
	return nil
}

func (ddc *DictionaryDetailCreate) sqlSave(ctx context.Context) (*DictionaryDetail, error) {
	if err := ddc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ddc.mutation.id = &_node.ID
	ddc.mutation.done = true
	return _node, nil
}

func (ddc *DictionaryDetailCreate) createSpec() (*DictionaryDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &DictionaryDetail{config: ddc.config}
		_spec = sqlgraph.NewCreateSpec(dictionarydetail.Table, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64))
	)
	if id, ok := ddc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ddc.mutation.CreatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ddc.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ddc.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := ddc.mutation.Sort(); ok {
		_spec.SetField(dictionarydetail.FieldSort, field.TypeUint32, value)
		_node.Sort = value
	}
	if value, ok := ddc.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ddc.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := ddc.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := ddc.mutation.DictionariesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionariesTable,
			Columns: []string{dictionarydetail.DictionariesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DictionaryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DictionaryDetailCreateBulk is the builder for creating many DictionaryDetail entities in bulk.
type DictionaryDetailCreateBulk struct {
	config
	err      error
	builders []*DictionaryDetailCreate
}

// Save creates the DictionaryDetail entities in the database.
func (ddcb *DictionaryDetailCreateBulk) Save(ctx context.Context) ([]*DictionaryDetail, error) {
	if ddcb.err != nil {
		return nil, ddcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ddcb.builders))
	nodes := make([]*DictionaryDetail, len(ddcb.builders))
	mutators := make([]Mutator, len(ddcb.builders))
	for i := range ddcb.builders {
		func(i int, root context.Context) {
			builder := ddcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictionaryDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, ddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddcb.driver, spec); err != nil {
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
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, ddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddcb *DictionaryDetailCreateBulk) SaveX(ctx context.Context) []*DictionaryDetail {
	v, err := ddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddcb *DictionaryDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := ddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddcb *DictionaryDetailCreateBulk) ExecX(ctx context.Context) {
	if err := ddcb.Exec(ctx); err != nil {
		panic(err)
	}
}
