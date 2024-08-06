// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/contrib/entgql/internal/todo/ent/onetomany"
	"entgo.io/contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneToManyQuery is the builder for querying OneToMany entities.
type OneToManyQuery struct {
	config
	ctx               *QueryContext
	order             []onetomany.OrderOption
	inters            []Interceptor
	predicates        []predicate.OneToMany
	withParent        *OneToManyQuery
	withChildren      *OneToManyQuery
	loadTotal         []func(context.Context, []*OneToMany) error
	modifiers         []func(*sql.Selector)
	withNamedChildren map[string]*OneToManyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OneToManyQuery builder.
func (otmq *OneToManyQuery) Where(ps ...predicate.OneToMany) *OneToManyQuery {
	otmq.predicates = append(otmq.predicates, ps...)
	return otmq
}

// Limit the number of records to be returned by this query.
func (otmq *OneToManyQuery) Limit(limit int) *OneToManyQuery {
	otmq.ctx.Limit = &limit
	return otmq
}

// Offset to start from.
func (otmq *OneToManyQuery) Offset(offset int) *OneToManyQuery {
	otmq.ctx.Offset = &offset
	return otmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otmq *OneToManyQuery) Unique(unique bool) *OneToManyQuery {
	otmq.ctx.Unique = &unique
	return otmq
}

// Order specifies how the records should be ordered.
func (otmq *OneToManyQuery) Order(o ...onetomany.OrderOption) *OneToManyQuery {
	otmq.order = append(otmq.order, o...)
	return otmq
}

// QueryParent chains the current query on the "parent" edge.
func (otmq *OneToManyQuery) QueryParent() *OneToManyQuery {
	query := (&OneToManyClient{config: otmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := otmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := otmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(onetomany.Table, onetomany.FieldID, selector),
			sqlgraph.To(onetomany.Table, onetomany.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, onetomany.ParentTable, onetomany.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(otmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (otmq *OneToManyQuery) QueryChildren() *OneToManyQuery {
	query := (&OneToManyClient{config: otmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := otmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := otmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(onetomany.Table, onetomany.FieldID, selector),
			sqlgraph.To(onetomany.Table, onetomany.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, onetomany.ChildrenTable, onetomany.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(otmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OneToMany entity from the query.
// Returns a *NotFoundError when no OneToMany was found.
func (otmq *OneToManyQuery) First(ctx context.Context) (*OneToMany, error) {
	nodes, err := otmq.Limit(1).All(setContextOp(ctx, otmq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{onetomany.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otmq *OneToManyQuery) FirstX(ctx context.Context) *OneToMany {
	node, err := otmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OneToMany ID from the query.
// Returns a *NotFoundError when no OneToMany ID was found.
func (otmq *OneToManyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otmq.Limit(1).IDs(setContextOp(ctx, otmq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{onetomany.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otmq *OneToManyQuery) FirstIDX(ctx context.Context) int {
	id, err := otmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OneToMany entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OneToMany entity is found.
// Returns a *NotFoundError when no OneToMany entities are found.
func (otmq *OneToManyQuery) Only(ctx context.Context) (*OneToMany, error) {
	nodes, err := otmq.Limit(2).All(setContextOp(ctx, otmq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{onetomany.Label}
	default:
		return nil, &NotSingularError{onetomany.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otmq *OneToManyQuery) OnlyX(ctx context.Context) *OneToMany {
	node, err := otmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OneToMany ID in the query.
// Returns a *NotSingularError when more than one OneToMany ID is found.
// Returns a *NotFoundError when no entities are found.
func (otmq *OneToManyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otmq.Limit(2).IDs(setContextOp(ctx, otmq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{onetomany.Label}
	default:
		err = &NotSingularError{onetomany.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otmq *OneToManyQuery) OnlyIDX(ctx context.Context) int {
	id, err := otmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OneToManies.
func (otmq *OneToManyQuery) All(ctx context.Context) ([]*OneToMany, error) {
	ctx = setContextOp(ctx, otmq.ctx, ent.OpQueryAll)
	if err := otmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OneToMany, *OneToManyQuery]()
	return withInterceptors[[]*OneToMany](ctx, otmq, qr, otmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (otmq *OneToManyQuery) AllX(ctx context.Context) []*OneToMany {
	nodes, err := otmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OneToMany IDs.
func (otmq *OneToManyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if otmq.ctx.Unique == nil && otmq.path != nil {
		otmq.Unique(true)
	}
	ctx = setContextOp(ctx, otmq.ctx, ent.OpQueryIDs)
	if err = otmq.Select(onetomany.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otmq *OneToManyQuery) IDsX(ctx context.Context) []int {
	ids, err := otmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otmq *OneToManyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, otmq.ctx, ent.OpQueryCount)
	if err := otmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, otmq, querierCount[*OneToManyQuery](), otmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (otmq *OneToManyQuery) CountX(ctx context.Context) int {
	count, err := otmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otmq *OneToManyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, otmq.ctx, ent.OpQueryExist)
	switch _, err := otmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (otmq *OneToManyQuery) ExistX(ctx context.Context) bool {
	exist, err := otmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OneToManyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otmq *OneToManyQuery) Clone() *OneToManyQuery {
	if otmq == nil {
		return nil
	}
	return &OneToManyQuery{
		config:       otmq.config,
		ctx:          otmq.ctx.Clone(),
		order:        append([]onetomany.OrderOption{}, otmq.order...),
		inters:       append([]Interceptor{}, otmq.inters...),
		predicates:   append([]predicate.OneToMany{}, otmq.predicates...),
		withParent:   otmq.withParent.Clone(),
		withChildren: otmq.withChildren.Clone(),
		// clone intermediate query.
		sql:  otmq.sql.Clone(),
		path: otmq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (otmq *OneToManyQuery) WithParent(opts ...func(*OneToManyQuery)) *OneToManyQuery {
	query := (&OneToManyClient{config: otmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	otmq.withParent = query
	return otmq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (otmq *OneToManyQuery) WithChildren(opts ...func(*OneToManyQuery)) *OneToManyQuery {
	query := (&OneToManyClient{config: otmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	otmq.withChildren = query
	return otmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OneToMany.Query().
//		GroupBy(onetomany.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (otmq *OneToManyQuery) GroupBy(field string, fields ...string) *OneToManyGroupBy {
	otmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OneToManyGroupBy{build: otmq}
	grbuild.flds = &otmq.ctx.Fields
	grbuild.label = onetomany.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.OneToMany.Query().
//		Select(onetomany.FieldName).
//		Scan(ctx, &v)
func (otmq *OneToManyQuery) Select(fields ...string) *OneToManySelect {
	otmq.ctx.Fields = append(otmq.ctx.Fields, fields...)
	sbuild := &OneToManySelect{OneToManyQuery: otmq}
	sbuild.label = onetomany.Label
	sbuild.flds, sbuild.scan = &otmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OneToManySelect configured with the given aggregations.
func (otmq *OneToManyQuery) Aggregate(fns ...AggregateFunc) *OneToManySelect {
	return otmq.Select().Aggregate(fns...)
}

func (otmq *OneToManyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range otmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, otmq); err != nil {
				return err
			}
		}
	}
	for _, f := range otmq.ctx.Fields {
		if !onetomany.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if otmq.path != nil {
		prev, err := otmq.path(ctx)
		if err != nil {
			return err
		}
		otmq.sql = prev
	}
	return nil
}

func (otmq *OneToManyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OneToMany, error) {
	var (
		nodes       = []*OneToMany{}
		_spec       = otmq.querySpec()
		loadedTypes = [2]bool{
			otmq.withParent != nil,
			otmq.withChildren != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OneToMany).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OneToMany{config: otmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(otmq.modifiers) > 0 {
		_spec.Modifiers = otmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := otmq.withParent; query != nil {
		if err := otmq.loadParent(ctx, query, nodes, nil,
			func(n *OneToMany, e *OneToMany) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := otmq.withChildren; query != nil {
		if err := otmq.loadChildren(ctx, query, nodes,
			func(n *OneToMany) { n.Edges.Children = []*OneToMany{} },
			func(n *OneToMany, e *OneToMany) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range otmq.withNamedChildren {
		if err := otmq.loadChildren(ctx, query, nodes,
			func(n *OneToMany) { n.appendNamedChildren(name) },
			func(n *OneToMany, e *OneToMany) { n.appendNamedChildren(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range otmq.loadTotal {
		if err := otmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (otmq *OneToManyQuery) loadParent(ctx context.Context, query *OneToManyQuery, nodes []*OneToMany, init func(*OneToMany), assign func(*OneToMany, *OneToMany)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OneToMany)
	for i := range nodes {
		fk := nodes[i].ParentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(onetomany.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (otmq *OneToManyQuery) loadChildren(ctx context.Context, query *OneToManyQuery, nodes []*OneToMany, init func(*OneToMany), assign func(*OneToMany, *OneToMany)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OneToMany)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(onetomany.FieldParentID)
	}
	query.Where(predicate.OneToMany(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(onetomany.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (otmq *OneToManyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otmq.querySpec()
	if len(otmq.modifiers) > 0 {
		_spec.Modifiers = otmq.modifiers
	}
	_spec.Node.Columns = otmq.ctx.Fields
	if len(otmq.ctx.Fields) > 0 {
		_spec.Unique = otmq.ctx.Unique != nil && *otmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, otmq.driver, _spec)
}

func (otmq *OneToManyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(onetomany.Table, onetomany.Columns, sqlgraph.NewFieldSpec(onetomany.FieldID, field.TypeInt))
	_spec.From = otmq.sql
	if unique := otmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if otmq.path != nil {
		_spec.Unique = true
	}
	if fields := otmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, onetomany.FieldID)
		for i := range fields {
			if fields[i] != onetomany.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if otmq.withParent != nil {
			_spec.Node.AddColumnOnce(onetomany.FieldParentID)
		}
	}
	if ps := otmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otmq *OneToManyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otmq.driver.Dialect())
	t1 := builder.Table(onetomany.Table)
	columns := otmq.ctx.Fields
	if len(columns) == 0 {
		columns = onetomany.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otmq.sql != nil {
		selector = otmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otmq.ctx.Unique != nil && *otmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range otmq.modifiers {
		m(selector)
	}
	for _, p := range otmq.predicates {
		p(selector)
	}
	for _, p := range otmq.order {
		p(selector)
	}
	if offset := otmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (otmq *OneToManyQuery) Modify(modifiers ...func(s *sql.Selector)) *OneToManySelect {
	otmq.modifiers = append(otmq.modifiers, modifiers...)
	return otmq.Select()
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (otmq *OneToManyQuery) WithNamedChildren(name string, opts ...func(*OneToManyQuery)) *OneToManyQuery {
	query := (&OneToManyClient{config: otmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if otmq.withNamedChildren == nil {
		otmq.withNamedChildren = make(map[string]*OneToManyQuery)
	}
	otmq.withNamedChildren[name] = query
	return otmq
}

// OneToManyGroupBy is the group-by builder for OneToMany entities.
type OneToManyGroupBy struct {
	selector
	build *OneToManyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otmgb *OneToManyGroupBy) Aggregate(fns ...AggregateFunc) *OneToManyGroupBy {
	otmgb.fns = append(otmgb.fns, fns...)
	return otmgb
}

// Scan applies the selector query and scans the result into the given value.
func (otmgb *OneToManyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otmgb.build.ctx, ent.OpQueryGroupBy)
	if err := otmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OneToManyQuery, *OneToManyGroupBy](ctx, otmgb.build, otmgb, otmgb.build.inters, v)
}

func (otmgb *OneToManyGroupBy) sqlScan(ctx context.Context, root *OneToManyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(otmgb.fns))
	for _, fn := range otmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*otmgb.flds)+len(otmgb.fns))
		for _, f := range *otmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*otmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OneToManySelect is the builder for selecting fields of OneToMany entities.
type OneToManySelect struct {
	*OneToManyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (otms *OneToManySelect) Aggregate(fns ...AggregateFunc) *OneToManySelect {
	otms.fns = append(otms.fns, fns...)
	return otms
}

// Scan applies the selector query and scans the result into the given value.
func (otms *OneToManySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otms.ctx, ent.OpQuerySelect)
	if err := otms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OneToManyQuery, *OneToManySelect](ctx, otms.OneToManyQuery, otms, otms.inters, v)
}

func (otms *OneToManySelect) sqlScan(ctx context.Context, root *OneToManyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(otms.fns))
	for _, fn := range otms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*otms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (otms *OneToManySelect) Modify(modifiers ...func(s *sql.Selector)) *OneToManySelect {
	otms.modifiers = append(otms.modifiers, modifiers...)
	return otms
}
