package myeasysql

import (
	"testing"
	"time"
)

type TestStruct struct {
	IntField    int
	BoolField   bool
	StringField string
	TimeField   time.Time
	AnyField    any
}

func Test_GetPointers(t *testing.T) {
	testTime := time.Now()
	testStruct := &TestStruct{
		IntField:    42,
		BoolField:   true,
		StringField: "test",
		TimeField:   testTime,
		AnyField:    123,
	}

	pointers, err := getPointers(testStruct)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(pointers) != 5 {
		t.Fatalf("expected 5 pointers, got %d", len(pointers))
	}

	if *pointers[0].(*int) != testStruct.IntField {
		t.Errorf("expected %d, got %d", testStruct.IntField, *pointers[0].(*int))
	}
	if *pointers[1].(*bool) != testStruct.BoolField {
		t.Errorf("expected %v, got %v", testStruct.BoolField, *pointers[1].(*bool))
	}
	if *pointers[2].(*string) != testStruct.StringField {
		t.Errorf("expected %s, got %s", testStruct.StringField, *pointers[2].(*string))
	}
	if !(*pointers[3].(*time.Time)).Equal(testStruct.TimeField) {
		t.Errorf("expected %v, got %v", testStruct.TimeField, *pointers[3].(*time.Time))
	}
	if *pointers[4].(*any) != testStruct.AnyField {
		t.Errorf("expected %v, got %v", testStruct.AnyField, *pointers[4].(*any))
	}
}

func Test_GetPointers_NotPointer(t *testing.T) {
	testStruct := TestStruct{}

	_, err := getPointers(testStruct)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "v is not a pointer" {
		t.Errorf("expected 'v is not a pointer', got %v", err)
	}
}

func Test_GetPointers_NotStruct(t *testing.T) {
	testInt := 42

	_, err := getPointers(&testInt)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "v is not a struct" {
		t.Errorf("expected 'v is not a struct', got %v", err)
	}
}
