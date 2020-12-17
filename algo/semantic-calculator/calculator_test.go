package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc_Sample1(t *testing.T) {
	v, e := calc("1+2")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(3), *v)
}

func TestCalc_Sample2(t *testing.T) {
	v, e := calc("2*2+3")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(7), *v)
}

func TestCalc_Sample3(t *testing.T) {
	v, e := calc("2+2*3")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(8), *v)
}

func TestCalc_Sample4(t *testing.T) {
	v, e := calc("2*(2+3)")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(10), *v)
}

func TestCalc_Sample5(t *testing.T) {
	v, e := calc("(3+3)*(2+3)")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(30), *v)
}

func TestCalc_Sample6(t *testing.T) {
	v, e := calc("(3+3)*(2+3)+10-2")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(38), *v)
}

func TestCalc_Sample7(t *testing.T) {
	v, e := calc("(3+3)*(2+3)-10-2")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(18), *v)
}

func TestCalc_Sample8(t *testing.T) {
	v, e := calc("8:4:2")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(1), *v)
}

func TestCalc_Sample9(t *testing.T) {
	v, e := calc("20-8:4:2-2-3")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(14), *v)
}

func TestCalc_Sample10(t *testing.T) {
	v, e := calc("40:(15+5)+(3+2)*10")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(52), *v)
}

func TestCalc_Fail_invalidSymbol(t *testing.T) {
	_, e := calc("1+1a")
	assert.Error(t, e)
}

func TestCalc_Whitespaces(t *testing.T) {
	v, e := calc(" 1 + 10")
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, float64(11), *v)
}
