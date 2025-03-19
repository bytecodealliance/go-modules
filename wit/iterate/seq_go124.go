//go:build go1.24 && !tinygo

package iterate

import "iter"

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// Requires generic type aliases: https://github.com/golang/go/issues/46477
type Seq[V any] = iter.Seq[V]

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
//
// Requires generic type aliases: https://github.com/golang/go/issues/46477
type Seq2[K, V any] = iter.Seq2[K, V]
