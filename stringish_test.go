package stringish

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFromByteSlice(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromByteSlice(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_GetString(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.GetString(); got != tt.want {
				t.Errorf("Stringish.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_ReplaceN(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		old   string
		new   string
		count int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.ReplaceN(tt.args.old, tt.args.new, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.ReplaceN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_ReplaceAll(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.ReplaceAll(tt.args.old, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.ReplaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_ToLower(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.ToLower(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_ToUpper(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.ToUpper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_TrimPrefix(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.TrimPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.TrimPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_TrimSuffix(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		suffix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.TrimSuffix(tt.args.suffix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.TrimSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_TrimSpaces(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.TrimSpaces(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.TrimSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_TrimPrefixSpaces(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.TrimPrefixSpaces(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.TrimPrefixSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_TrimSuffixSpaces(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.TrimSuffixSpaces(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.TrimSuffixSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_HasPrefix(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.HasPrefix(tt.args.prefix); got != tt.want {
				t.Errorf("Stringish.HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_HasSuffix(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		suffix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.HasSuffix(tt.args.suffix); got != tt.want {
				t.Errorf("Stringish.HasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Equals(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Equals(tt.args.str); got != tt.want {
				t.Errorf("Stringish.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_IsEmpty(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Stringish.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_IsBlank(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.IsBlank(); got != tt.want {
				t.Errorf("Stringish.IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Contains(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Contains(tt.args.str); got != tt.want {
				t.Errorf("Stringish.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Len(t *testing.T) {
	type fields struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("Stringish.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Index(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Index(tt.args.str); got != tt.want {
				t.Errorf("Stringish.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_LastIndex(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.LastIndex(tt.args.str); got != tt.want {
				t.Errorf("Stringish.LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Map(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		mapper func(string) string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Map(tt.args.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringish_Filter(t *testing.T) {
	type fields struct {
		str string
	}
	type args struct {
		filterer func(string) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stringish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringish{
				str: tt.fields.str,
			}
			if got := s.Filter(tt.args.filterer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stringish.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
