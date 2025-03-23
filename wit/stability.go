package wit

import (
	"github.com/coreos/go-semver/semver"
)

// Stability represents the version or feature-gated stability of a given feature.
type Stability interface {
	Node
	isStability()
	Versioned() bool
}

// _stability is an embeddable type that conforms to the [Stability] interface.
type _stability struct{}

func (_stability) isStability() {}

// Stable represents a stable WIT feature, for example: @since(version = 1.2.3)
//
// Stable features have an explicit since version and an optional feature name.
type Stable struct {
	_stability
	Since      semver.Version
	Deprecated *semver.Version
}

// Versioned always returns true for [Stable].
func (*Stable) Versioned() bool {
	return true
}

// Unstable represents an unstable WIT feature defined by name.
type Unstable struct {
	_stability
	Feature    string
	Deprecated *semver.Version
}

// Versioned returns whether [Unstable] is deprecated.
func (u *Unstable) Versioned() bool {
	return u.Deprecated != nil
}
