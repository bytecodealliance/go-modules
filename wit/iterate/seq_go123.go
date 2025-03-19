//go:build !go1.24 || tinygo

package iterate

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// TODO: delete this once we drop support for Go 1.23.
type Seq[V any] func(yield func(V) bool)

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
//
// TODO: delete this once we drop support for Go 1.23.
type Seq2[K, V any] func(yield func(K, V) bool)
