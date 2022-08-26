package bigint_test

import (
	"bootcamp/bigint/bigint"
	"testing"
)

func TestAbs(t *testing.T) {
	a := bigint.Bigint{
		Value: "-1",
	}

	b := a.Abs()
	if b.Value != "1" {
		t.Errorf("Abs FAILED: expected 1 but got %v", b)
	}

	a = bigint.Bigint{
		Value: "+1",
	}

	b = a.Abs()
	if b.Value != "1" {
		t.Errorf("Abs FAILED: expected 1 but got %v", b)
	}

	a = bigint.Bigint{
		Value: "1",
	}

	b = a.Abs()
	if b.Value != "1" {
		t.Errorf("Abs FAILED: expected 1 but got %v", b)
	}
}

func TestNewInt(t *testing.T) {
	str := "123123132"
	a, err := bigint.NewInt(str)
	if err != nil {
		t.Errorf("NewInt FAILED: %v", err)
	}
	if a.Value != str {
		t.Errorf("NewInt FAILED: expected %s but got %s", str, a.Value)
	}

	str2 := "123asd123"
	_, err2 := bigint.NewInt(str2)
	if err2 != bigint.ErrorBadInput {
		t.Errorf("NewInt FAILED: it should return %v", bigint.ErrorBadInput)
	}
}
