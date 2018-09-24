package stringish

import (
	"strings"
)

// Stringish is the main exposed type with all implemented methods
type Stringish struct {
	str string
}

// New creates an instance of Stringish and returns a pointer to it
func New(str string) *Stringish {
	return &Stringish{str}
}

// NewFromByteSlice creates an instance of Stringish from a slice of bytes and returns a pointer to it
func NewFromByteSlice(bytes []byte) *Stringish {
	return &Stringish{string(bytes)}
}

// GetString returns the final string result
func (s *Stringish) GetString() string {
	return s.str
}

// ReplaceN replaces old with new with limit of count time[s]
func (s *Stringish) ReplaceN(old, new string, count int) *Stringish {
	s.str = strings.Replace(s.str, old, new, count)
	return s
}

// ReplaceAll replaces all occurrences of old with new
func (s *Stringish) ReplaceAll(old, new string) *Stringish {
	return s.ReplaceN(old, new, -1)
}

// ToLower lowers all characters of the string
func (s *Stringish) ToLower() *Stringish {
	s.str = strings.ToLower(s.str)
	return s
}

// ToUpper lowers all characters of the string
func (s *Stringish) ToUpper() *Stringish {
	s.str = strings.ToUpper(s.str)
	return s
}

// TrimPrefix trims prefix from the start of the string
func (s *Stringish) TrimPrefix(prefix string) *Stringish {
	s.str = strings.TrimPrefix(s.str, prefix)
	return s
}

// TrimSuffix trims suffix from the end of the string
func (s *Stringish) TrimSuffix(suffix string) *Stringish {
	s.str = strings.TrimSuffix(s.str, suffix)
	return s
}

// TrimSpaces trims space characters from both sides of the string
func (s *Stringish) TrimSpaces() *Stringish {
	s.str = strings.TrimSpace(s.str)
	return s
}

// TrimPrefixSpaces trims space characters from the start of the string
func (s *Stringish) TrimPrefixSpaces() *Stringish {
	return s.TrimPrefix(" ")
}

// TrimSuffixSpaces trims space characters from the end of the string
func (s *Stringish) TrimSuffixSpaces() *Stringish {
	return s.TrimSuffix(" ")
}

// HasPrefix returns true if string starts with prefix
func (s *Stringish) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.str, prefix)
}

// HasSuffix returns true if string ends with suffix
func (s *Stringish) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.str, suffix)
}

// Equals returns true if string equals str
func (s *Stringish) Equals(str string) bool {
	return s.str == str
}

// IsEmpty returns true if string is empty (strictly equal to "")
func (s *Stringish) IsEmpty() bool {
	return s.Equals("")
}

// IsBlank returns true if string is blank (equal to "" or some space chars like "  ")
func (s *Stringish) IsBlank() bool {
	return s.TrimSpaces().IsEmpty()
}

// Contains returns true if string contains the str value
func (s *Stringish) Contains(str string) bool {
	return strings.Contains(s.str, str)
}

// Len returns the len of the string
func (s *Stringish) Len() int {
	return len(s.str)
}

// Index returns the index of first occurrence of str in string (-1 if no occurrence found)
func (s *Stringish) Index(str string) int {
	return strings.Index(s.str, str)
}

// LastIndex returns the index of last occurrence of str in string (-1 if no occurrence found)
func (s *Stringish) LastIndex(str string) int {
	return strings.LastIndex(s.str, str)
}

// Map maps every character of string with given function and replace them with return value of the function
func (s *Stringish) Map(mapper func(string) string) *Stringish {
	s.str = strings.Map(func(char rune) rune {
		resultChar := mapper(string(char))
		return []rune(resultChar)[0]
	}, s.str)
	return s
}

// Filter runs filterer on characters of the string and holds the character if the filterer returns true
func (s *Stringish) Filter(filterer func(string) bool) *Stringish {
	s.str = strings.Map(func(char rune) rune {
		result := filterer(string(char))
		if result {
			return char
		}
		return -1
	}, s.str)
	return s
}
