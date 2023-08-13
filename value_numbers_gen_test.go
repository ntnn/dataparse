package dataparse

import (
	"testing"
	"log"
	"fmt"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestValue_Int_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Int()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Int8_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Int8()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Int16_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Int16()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Int32_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Int32()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Int64_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Int64()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Uint_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Uint()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Uint8_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Uint8()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Uint16_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Uint16()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Uint32_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Uint32()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Uint64_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Uint64()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Float32_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Float32()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

func TestValue_Float64_Fuzz(t *testing.T) {
	t.Parallel()

	for i := 0; i < parallelFuzzTests; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			val := Value{}
			fuzz.New().Fuzz(&val)
			parsed, err := val.Float64()
			log.Printf("%v -> %v", val.Data, parsed)
			assert.Nil(t, err)
		})
	}
}

