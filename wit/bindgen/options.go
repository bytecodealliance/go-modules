package bindgen

import (
	"go.bytecodealliance.org/wit/logging"
)

// Option represents a single configuration option for this package.
type Option interface {
	applyOption(*options) error
}

type optionFunc func(*options) error

func (f optionFunc) applyOption(opts *options) error {
	return f(opts)
}

type options struct {
	logger logging.Logger

	// generatedBy is the name of the program that generates code with this package.
	generatedBy string

	// world is the name of the WIT world to generate, e.g. "command" or "wasi:cli/command".
	// Default: the first world in the Resolve will be generated.
	world string

	// packageRoot is the root Go package or module path used in generated code.
	packageRoot string

	// cmPackage is the package path to the "cm" or Component Model package with basic types.
	// Default: go.bytecodealliance.org/cm.
	cmPackage string

	// versioned determines if Go packages are generated with version numbers.
	versioned bool

	// generateWIT determines if WIT files will be generated for each world and interface.
	generateWIT bool
}

func (opts *options) apply(o ...Option) error {
	for _, o := range o {
		err := o.applyOption(opts)
		if err != nil {
			return err
		}
	}
	return nil
}

// Logger returns an [Option] that specifies a [logging.Logger] for logging.
func Logger(logger logging.Logger) Option {
	return optionFunc(func(opts *options) error {
		opts.logger = logger
		return nil
	})
}

// GeneratedBy returns an [Option] that specifies the name of the program or package
// that will appear in the "Code generated by ..." header on generated files.
func GeneratedBy(name string) Option {
	return optionFunc(func(opts *options) error {
		opts.generatedBy = name
		return nil
	})
}

// World returns an [Option] that specifies the WIT world to generate.
// By default, the first world will be generated.
func World(world string) Option {
	return optionFunc(func(opts *options) error {
		opts.world = world
		return nil
	})
}

// PackageRoot returns an [Option] that specifies the root Go package path for generated Go packages.
func PackageRoot(path string) Option {
	return optionFunc(func(opts *options) error {
		opts.packageRoot = path
		return nil
	})
}

// CMPackage returns an [Option] that specifies the package path to the
// Component Model utility package (default: go.bytecodealliance.org/cm).
func CMPackage(path string) Option {
	return optionFunc(func(opts *options) error {
		opts.cmPackage = path
		return nil
	})
}

// Versioned returns an [Option] that specifies that all generated Go packages
// will have versions that match WIT versions.
func Versioned(versioned bool) Option {
	return optionFunc(func(opts *options) error {
		opts.versioned = versioned
		return nil
	})
}

// WIT returns an [Option] that specifies that a WIT file will be generated
// for each Go package corresponding to each WIT world and interface.
func WIT(generateWIT bool) Option {
	return optionFunc(func(opts *options) error {
		opts.generateWIT = generateWIT
		return nil
	})
}
