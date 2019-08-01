// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testServices(t *testing.T) {
	t.Parallel()

	query := Services()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServicesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Services().DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceSlice{o}

	if rowsAff, err := slice.DeleteAll(tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceExists(tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Service exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceExists to return true, but got false.")
	}
}

func testServicesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceFound, err := FindService(tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServicesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Services().Bind(nil, tx, o); err != nil {
		t.Error(err)
	}
}

func testServicesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Services().One(tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServicesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceOne := &Service{}
	serviceTwo := &Service{}
	if err = randomize.Struct(seed, serviceOne, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceTwo, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = serviceOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Services().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServicesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceOne := &Service{}
	serviceTwo := &Service{}
	if err = randomize.Struct(seed, serviceOne, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceTwo, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = serviceOne.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceTwo.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBeforeInsertHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterInsertHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterSelectHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeUpdateHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterUpdateHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeDeleteHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterDeleteHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceBeforeUpsertHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func serviceAfterUpsertHook(e boil.Executor, o *Service) error {
	*o = Service{}
	return nil
}

func testServicesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Service{}
	o := &Service{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Service object: %s", err)
	}

	AddServiceHook(boil.BeforeInsertHook, serviceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBeforeInsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterInsertHook, serviceAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceAfterInsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterSelectHook, serviceAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceAfterSelectHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeUpdateHook, serviceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBeforeUpdateHooks = []ServiceHook{}

	AddServiceHook(boil.AfterUpdateHook, serviceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceAfterUpdateHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeDeleteHook, serviceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBeforeDeleteHooks = []ServiceHook{}

	AddServiceHook(boil.AfterDeleteHook, serviceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceAfterDeleteHooks = []ServiceHook{}

	AddServiceHook(boil.BeforeUpsertHook, serviceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBeforeUpsertHooks = []ServiceHook{}

	AddServiceHook(boil.AfterUpsertHook, serviceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceAfterUpsertHooks = []ServiceHook{}
}

func testServicesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Whitelist(serviceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testServicesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceSlice{o}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testServicesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Services().All(tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Description`: `text`, `ServicePath`: `varchar`, `TargetHost`: `varchar`, `TargetPath`: `varchar`, `TargetPort`: `int`, `Methods`: `set`}
	_              = bytes.MinRead
)

func testServicesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceAllColumns) == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDBTypes, true, servicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	if rowsAff, err := o.Update(tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServicesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceAllColumns) == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Service{}
	if err = randomize.Struct(seed, o, serviceDBTypes, true, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDBTypes, true, servicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceAllColumns, servicePrimaryKeyColumns) {
		fields = serviceAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceAllColumns,
			servicePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ServiceSlice{o}
	if rowsAff, err := slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServicesUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceAllColumns) == len(servicePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Service{}
	if err = randomize.Struct(seed, &o, serviceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Service: %s", err)
	}

	count, err := Services().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceDBTypes, false, servicePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	if err = o.Upsert(tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Service: %s", err)
	}

	count, err = Services().Count(tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
