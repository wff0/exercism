package nucleotidecount

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
// type Histogram ...
type Histogram map[byte]int

func NewHistogram() Histogram {
	h := make(map[byte]int)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	return h
}

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
// type DNA ...
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()
	for i := range d {
		if d.valid(d[i]) {
			h[d[i]]++
		} else {
			return nil, errors.ErrUnsupported
		}
	}
	return h, nil
}

func (d DNA) valid(nucl byte) bool {
	if nucl == 'A' || nucl == 'C' || nucl == 'G' || nucl == 'T' {
		return true
	}
	return false
}
