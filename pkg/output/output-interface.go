package output

import (
	"github.com/DrSmithFr/go-console/pkg/formatter"
	"github.com/DrSmithFr/go-console/pkg/verbosity"
)

// OutputInterface is the interface implemented by all Output classes
type OutputInterface interface {
	// Formats a message according to the current formatter styles.
	format(message string) string

	// Writes a message to the output.
	Write(message string)

	// Writes a message to the output and adds a newline at the end.
	Writeln(message string)

	// Writes a message to the output.
	WriteOnVerbose(message string, verbosity verbosity.Level)

	// Writes a message to the output and adds a newline at the end.
	WritelnOnVerbose(message string, verbosity verbosity.Level)

	// Sets the decorated flag.
	SetDecorated(decorated bool)

	// Gets the decorated flag.
	IsDecorated() bool

	// Sets current output formatter instance.
	SetFormatter(formatter *formatter.OutputFormatter)

	// Gets current output formatter instance.
	GetFormatter() *formatter.OutputFormatter

	// Sets the verbosity of the output.
	SetVerbosity(verbosity verbosity.Level)

	// Gets the current verbosity of the output.
	GetVerbosity() verbosity.Level

	// Returns whether verbosity is quiet (-q)
	IsQuiet() bool

	// Returns whether verbosity is verbose (-v)
	IsVerbose() bool

	// Returns whether verbosity is very verbose (-vv)
	IsVeryVerbose() bool

	// Returns whether verbosity is debug (-vvv)
	IsDebug() bool
}
