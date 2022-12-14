// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types/pgeo"
	"github.com/volatiletech/strmangle"
)

// Geolocation is an object representing the database table.
type Geolocation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	IPAddress    string      `boil:"ip_address" json:"ip_address" toml:"ip_address" yaml:"ip_address"`
	CountryCode  null.String `boil:"country_code" json:"country_code,omitempty" toml:"country_code" yaml:"country_code,omitempty"`
	Country      null.String `boil:"country" json:"country,omitempty" toml:"country" yaml:"country,omitempty"`
	City         null.String `boil:"city" json:"city,omitempty" toml:"city" yaml:"city,omitempty"`
	Coordinates  pgeo.Point  `boil:"coordinates" json:"coordinates" toml:"coordinates" yaml:"coordinates"`
	MysteryValue null.String `boil:"mystery_value" json:"mystery_value,omitempty" toml:"mystery_value" yaml:"mystery_value,omitempty"`
	CreatedAt    null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *geolocationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L geolocationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GeolocationColumns = struct {
	ID           string
	IPAddress    string
	CountryCode  string
	Country      string
	City         string
	Coordinates  string
	MysteryValue string
	CreatedAt    string
}{
	ID:           "id",
	IPAddress:    "ip_address",
	CountryCode:  "country_code",
	Country:      "country",
	City:         "city",
	Coordinates:  "coordinates",
	MysteryValue: "mystery_value",
	CreatedAt:    "created_at",
}

var GeolocationTableColumns = struct {
	ID           string
	IPAddress    string
	CountryCode  string
	Country      string
	City         string
	Coordinates  string
	MysteryValue string
	CreatedAt    string
}{
	ID:           "geolocations.id",
	IPAddress:    "geolocations.ip_address",
	CountryCode:  "geolocations.country_code",
	Country:      "geolocations.country",
	City:         "geolocations.city",
	Coordinates:  "geolocations.coordinates",
	MysteryValue: "geolocations.mystery_value",
	CreatedAt:    "geolocations.created_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperpgeo_Point struct{ field string }

func (w whereHelperpgeo_Point) EQ(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperpgeo_Point) NEQ(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperpgeo_Point) LT(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperpgeo_Point) LTE(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperpgeo_Point) GT(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperpgeo_Point) GTE(x pgeo.Point) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var GeolocationWhere = struct {
	ID           whereHelperint
	IPAddress    whereHelperstring
	CountryCode  whereHelpernull_String
	Country      whereHelpernull_String
	City         whereHelpernull_String
	Coordinates  whereHelperpgeo_Point
	MysteryValue whereHelpernull_String
	CreatedAt    whereHelpernull_Time
}{
	ID:           whereHelperint{field: "\"geolocations\".\"id\""},
	IPAddress:    whereHelperstring{field: "\"geolocations\".\"ip_address\""},
	CountryCode:  whereHelpernull_String{field: "\"geolocations\".\"country_code\""},
	Country:      whereHelpernull_String{field: "\"geolocations\".\"country\""},
	City:         whereHelpernull_String{field: "\"geolocations\".\"city\""},
	Coordinates:  whereHelperpgeo_Point{field: "\"geolocations\".\"coordinates\""},
	MysteryValue: whereHelpernull_String{field: "\"geolocations\".\"mystery_value\""},
	CreatedAt:    whereHelpernull_Time{field: "\"geolocations\".\"created_at\""},
}

// GeolocationRels is where relationship names are stored.
var GeolocationRels = struct {
}{}

// geolocationR is where relationships are stored.
type geolocationR struct {
}

// NewStruct creates a new relationship struct
func (*geolocationR) NewStruct() *geolocationR {
	return &geolocationR{}
}

// geolocationL is where Load methods for each relationship are stored.
type geolocationL struct{}

var (
	geolocationAllColumns            = []string{"id", "ip_address", "country_code", "country", "city", "coordinates", "mystery_value", "created_at"}
	geolocationColumnsWithoutDefault = []string{"ip_address", "coordinates"}
	geolocationColumnsWithDefault    = []string{"id", "country_code", "country", "city", "mystery_value", "created_at"}
	geolocationPrimaryKeyColumns     = []string{"id"}
	geolocationGeneratedColumns      = []string{}
)

type (
	// GeolocationSlice is an alias for a slice of pointers to Geolocation.
	// This should almost always be used instead of []Geolocation.
	GeolocationSlice []*Geolocation
	// GeolocationHook is the signature for custom Geolocation hook methods
	GeolocationHook func(context.Context, boil.ContextExecutor, *Geolocation) error

	geolocationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	geolocationType                 = reflect.TypeOf(&Geolocation{})
	geolocationMapping              = queries.MakeStructMapping(geolocationType)
	geolocationPrimaryKeyMapping, _ = queries.BindMapping(geolocationType, geolocationMapping, geolocationPrimaryKeyColumns)
	geolocationInsertCacheMut       sync.RWMutex
	geolocationInsertCache          = make(map[string]insertCache)
	geolocationUpdateCacheMut       sync.RWMutex
	geolocationUpdateCache          = make(map[string]updateCache)
	geolocationUpsertCacheMut       sync.RWMutex
	geolocationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var geolocationAfterSelectHooks []GeolocationHook

var geolocationBeforeInsertHooks []GeolocationHook
var geolocationAfterInsertHooks []GeolocationHook

var geolocationBeforeUpdateHooks []GeolocationHook
var geolocationAfterUpdateHooks []GeolocationHook

var geolocationBeforeDeleteHooks []GeolocationHook
var geolocationAfterDeleteHooks []GeolocationHook

var geolocationBeforeUpsertHooks []GeolocationHook
var geolocationAfterUpsertHooks []GeolocationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Geolocation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Geolocation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Geolocation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Geolocation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Geolocation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Geolocation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Geolocation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Geolocation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Geolocation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range geolocationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGeolocationHook registers your hook function for all future operations.
func AddGeolocationHook(hookPoint boil.HookPoint, geolocationHook GeolocationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		geolocationAfterSelectHooks = append(geolocationAfterSelectHooks, geolocationHook)
	case boil.BeforeInsertHook:
		geolocationBeforeInsertHooks = append(geolocationBeforeInsertHooks, geolocationHook)
	case boil.AfterInsertHook:
		geolocationAfterInsertHooks = append(geolocationAfterInsertHooks, geolocationHook)
	case boil.BeforeUpdateHook:
		geolocationBeforeUpdateHooks = append(geolocationBeforeUpdateHooks, geolocationHook)
	case boil.AfterUpdateHook:
		geolocationAfterUpdateHooks = append(geolocationAfterUpdateHooks, geolocationHook)
	case boil.BeforeDeleteHook:
		geolocationBeforeDeleteHooks = append(geolocationBeforeDeleteHooks, geolocationHook)
	case boil.AfterDeleteHook:
		geolocationAfterDeleteHooks = append(geolocationAfterDeleteHooks, geolocationHook)
	case boil.BeforeUpsertHook:
		geolocationBeforeUpsertHooks = append(geolocationBeforeUpsertHooks, geolocationHook)
	case boil.AfterUpsertHook:
		geolocationAfterUpsertHooks = append(geolocationAfterUpsertHooks, geolocationHook)
	}
}

// OneG returns a single geolocation record from the query using the global executor.
func (q geolocationQuery) OneG(ctx context.Context) (*Geolocation, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single geolocation record from the query.
func (q geolocationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Geolocation, error) {
	o := &Geolocation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for geolocations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Geolocation records from the query using the global executor.
func (q geolocationQuery) AllG(ctx context.Context) (GeolocationSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Geolocation records from the query.
func (q geolocationQuery) All(ctx context.Context, exec boil.ContextExecutor) (GeolocationSlice, error) {
	var o []*Geolocation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Geolocation slice")
	}

	if len(geolocationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Geolocation records in the query using the global executor
func (q geolocationQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Geolocation records in the query.
func (q geolocationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count geolocations rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q geolocationQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q geolocationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if geolocations exists")
	}

	return count > 0, nil
}

// Geolocations retrieves all the records using an executor.
func Geolocations(mods ...qm.QueryMod) geolocationQuery {
	mods = append(mods, qm.From("\"geolocations\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"geolocations\".*"})
	}

	return geolocationQuery{q}
}

// FindGeolocationG retrieves a single record by ID.
func FindGeolocationG(ctx context.Context, iD int, selectCols ...string) (*Geolocation, error) {
	return FindGeolocation(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindGeolocation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGeolocation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Geolocation, error) {
	geolocationObj := &Geolocation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"geolocations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, geolocationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from geolocations")
	}

	if err = geolocationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return geolocationObj, err
	}

	return geolocationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Geolocation) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Geolocation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no geolocations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(geolocationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	geolocationInsertCacheMut.RLock()
	cache, cached := geolocationInsertCache[key]
	geolocationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			geolocationAllColumns,
			geolocationColumnsWithDefault,
			geolocationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(geolocationType, geolocationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(geolocationType, geolocationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"geolocations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"geolocations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into geolocations")
	}

	if !cached {
		geolocationInsertCacheMut.Lock()
		geolocationInsertCache[key] = cache
		geolocationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Geolocation record using the global executor.
// See Update for more documentation.
func (o *Geolocation) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Geolocation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Geolocation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	geolocationUpdateCacheMut.RLock()
	cache, cached := geolocationUpdateCache[key]
	geolocationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			geolocationAllColumns,
			geolocationPrimaryKeyColumns,
		)
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update geolocations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"geolocations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, geolocationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(geolocationType, geolocationMapping, append(wl, geolocationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update geolocations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for geolocations")
	}

	if !cached {
		geolocationUpdateCacheMut.Lock()
		geolocationUpdateCache[key] = cache
		geolocationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q geolocationQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q geolocationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for geolocations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for geolocations")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o GeolocationSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GeolocationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), geolocationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"geolocations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, geolocationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in geolocation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all geolocation")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Geolocation) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Geolocation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no geolocations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(geolocationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	geolocationUpsertCacheMut.RLock()
	cache, cached := geolocationUpsertCache[key]
	geolocationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			geolocationAllColumns,
			geolocationColumnsWithDefault,
			geolocationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			geolocationAllColumns,
			geolocationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("model: unable to upsert geolocations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(geolocationPrimaryKeyColumns))
			copy(conflict, geolocationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"geolocations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(geolocationType, geolocationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(geolocationType, geolocationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "model: unable to upsert geolocations")
	}

	if !cached {
		geolocationUpsertCacheMut.Lock()
		geolocationUpsertCache[key] = cache
		geolocationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Geolocation record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Geolocation) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Geolocation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Geolocation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Geolocation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), geolocationPrimaryKeyMapping)
	sql := "DELETE FROM \"geolocations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from geolocations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for geolocations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q geolocationQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q geolocationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no geolocationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from geolocations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for geolocations")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o GeolocationSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GeolocationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(geolocationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), geolocationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"geolocations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, geolocationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from geolocation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for geolocations")
	}

	if len(geolocationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Geolocation) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: no Geolocation provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Geolocation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGeolocation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GeolocationSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("model: empty GeolocationSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GeolocationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GeolocationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), geolocationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"geolocations\".* FROM \"geolocations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, geolocationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in GeolocationSlice")
	}

	*o = slice

	return nil
}

// GeolocationExistsG checks if the Geolocation row exists.
func GeolocationExistsG(ctx context.Context, iD int) (bool, error) {
	return GeolocationExists(ctx, boil.GetContextDB(), iD)
}

// GeolocationExists checks if the Geolocation row exists.
func GeolocationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"geolocations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if geolocations exists")
	}

	return exists, nil
}
