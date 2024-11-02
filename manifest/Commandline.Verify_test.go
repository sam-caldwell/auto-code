package manifest

import (
	"github.com/sam-caldwell/auto-code/manifest/messages"
	"testing"
)

func TestCommandline_Verify(t *testing.T) {

	t.Run("create a Commandline argument", func(t *testing.T) {
		var cmd Commandline
		cmd.Required = true
		cmd.Short = "-a"
		cmd.Long = "--long_arg"
	})

	t.Run("sad path: empty short and long should fail", func(t *testing.T) {
		var cmd Commandline
		cmd.Required = true
		cmd.Short = EmptyString
		cmd.Long = EmptyString
		if err := cmd.Verify(nil); err == nil {
			t.Fatalf("Verify() should have failed since Short and Long are empty")
		} else {
			if msg := err.Error(); msg != messages.ErrMissingCommandlineArgument {
				t.Fatalf("Verify() error mismatch.\n"+
					"Got %s\n"+
					"Want %s",
					msg, messages.ErrMissingCommandlineArgument)
			}
		}
	})

	t.Run("test arguments", func(t *testing.T) {
		type TestData struct {
			short bool
			long  bool
			valid bool
		}
		testData := map[string]TestData{
			"-a": {short: true, valid: true}, "-b": {short: true, valid: true},
			"-c": {short: true, valid: true}, "-d": {short: true, valid: true},
			"-e": {short: true, valid: true}, "-f": {short: true, valid: true},
			"-g": {short: true, valid: true}, "-h": {short: true, valid: true},
			"-i": {short: true, valid: true}, "-j": {short: true, valid: true},
			"-k": {short: true, valid: true}, "-l": {short: true, valid: true},
			"-m": {short: true, valid: true}, "-n": {short: true, valid: true},
			"-o": {short: true, valid: true}, "-p": {short: true, valid: true},
			"-q": {short: true, valid: true}, "-r": {short: true, valid: true},
			"-s": {short: true, valid: true}, "-t": {short: true, valid: true},
			"-u": {short: true, valid: true}, "-v": {short: true, valid: true},
			"-w": {short: true, valid: true}, "-x": {short: true, valid: true},
			"-y": {short: true, valid: true}, "-z": {short: true, valid: true},
			"-A": {short: true, valid: true}, "-B": {short: true, valid: true},
			"-C": {short: true, valid: true}, "-D": {short: true, valid: true},
			"-E": {short: true, valid: true}, "-F": {short: true, valid: true},
			"-G": {short: true, valid: true}, "-H": {short: true, valid: true},
			"-I": {short: true, valid: true}, "-J": {short: true, valid: true},
			"-K": {short: true, valid: true}, "-L": {short: true, valid: true},
			"-M": {short: true, valid: true}, "-N": {short: true, valid: true},
			"-O": {short: true, valid: true}, "-P": {short: true, valid: true},
			"-Q": {short: true, valid: true}, "-R": {short: true, valid: true},
			"-S": {short: true, valid: true}, "-T": {short: true, valid: true},
			"-U": {short: true, valid: true}, "-V": {short: true, valid: true},
			"-W": {short: true, valid: true}, "-X": {short: true, valid: true},
			"-Y": {short: true, valid: true}, "-Z": {short: true, valid: true},
			"-0": {short: true, valid: true}, "-1": {short: true, valid: true},
			"-2": {short: true, valid: true}, "-3": {short: true, valid: true},
			"-4": {short: true, valid: true}, "-5": {short: true, valid: true},
			"-6": {short: true, valid: true}, "-7": {short: true, valid: true},
			"-8": {short: true, valid: true}, "-9": {short: true, valid: true},
			"-_": {short: true, valid: false}, "-": {short: true, valid: false},
			"-!": {short: true, valid: false}, "-*": {short: true, valid: false},
			"-@": {short: true, valid: false}, "-(": {short: true, valid: false},
			"-#": {short: true, valid: false}, "-)": {short: true, valid: false},
			"-$": {short: true, valid: false}, "-+": {short: true, valid: false},
			"-%": {short: true, valid: false}, "-{": {short: true, valid: false},
			"-^": {short: true, valid: false}, "-}": {short: true, valid: false},
			"-;": {short: true, valid: false}, "-[": {short: true, valid: false},
			"-:": {short: true, valid: false}, "-]": {short: true, valid: false},
			"-'": {short: true, valid: false}, "-\"": {short: true, valid: false},
			"-`": {short: true, valid: false}, "-~": {short: true, valid: false},
			"-<": {short: true, valid: false}, "->": {short: true, valid: false},
			"-,": {short: true, valid: false}, "-.": {short: true, valid: false},
			"-\\": {short: true, valid: false}, "-?": {short: true, valid: false},
			"-/": {short: true, valid: false}, "--": {short: true, long: true, valid: false},
			"--a": {long: true, valid: false}, "--b": {long: true, valid: false},
			"--c": {long: true, valid: false}, "--d": {long: true, valid: false},
			"--e": {long: true, valid: false}, "--f": {long: true, valid: false},
			"--g": {long: true, valid: false}, "--h": {long: true, valid: false},
			"--i": {long: true, valid: false}, "--j": {long: true, valid: false},
			"--k": {long: true, valid: false}, "--l": {long: true, valid: false},
			"--m": {long: true, valid: false}, "--n": {long: true, valid: false},
			"--o": {long: true, valid: false}, "--p": {long: true, valid: false},
			"--q": {long: true, valid: false}, "--r": {long: true, valid: false},
			"--s": {long: true, valid: false}, "--t": {long: true, valid: false},
			"--u": {long: true, valid: false}, "--v": {long: true, valid: false},
			"--w": {long: true, valid: false}, "--x": {long: true, valid: false},
			"--y": {long: true, valid: false}, "--z": {long: true, valid: false},
			"--A": {long: true, valid: false}, "--B": {long: true, valid: false},
			"--C": {long: true, valid: false}, "--D": {long: true, valid: false},
			"--E": {long: true, valid: false}, "--F": {long: true, valid: false},
			"--G": {long: true, valid: false}, "--H": {long: true, valid: false},
			"--I": {long: true, valid: false}, "--J": {long: true, valid: false},
			"--K": {long: true, valid: false}, "--L": {long: true, valid: false},
			"--M": {long: true, valid: false}, "--N": {long: true, valid: false},
			"--O": {long: true, valid: false}, "--P": {long: true, valid: false},
			"--Q": {long: true, valid: false}, "--R": {long: true, valid: false},
			"--S": {long: true, valid: false}, "--T": {long: true, valid: false},
			"--U": {long: true, valid: false}, "--V": {long: true, valid: false},
			"--W": {long: true, valid: false}, "--X": {long: true, valid: false},
			"--Y": {long: true, valid: false}, "--Z": {long: true, valid: false},
			"--0": {long: true, valid: false}, "--1": {long: true, valid: false},
			"--2": {long: true, valid: false}, "--3": {long: true, valid: false},
			"--4": {long: true, valid: false}, "--5": {long: true, valid: false},
			"--6": {long: true, valid: false}, "--7": {long: true, valid: false},
			"--8": {long: true, valid: false}, "--9": {long: true, valid: false},
			"--_": {long: true, valid: false}, "---": {long: true, valid: false},
			"--!": {long: true, valid: false}, "--*": {long: true, valid: false},
			"--@": {long: true, valid: false}, "--(": {long: true, valid: false},
			"--#": {long: true, valid: false}, "--)": {long: true, valid: false},
			"--$": {long: true, valid: false}, "--+": {long: true, valid: false},
			"--%": {long: true, valid: false}, "--{": {long: true, valid: false},
			"--^": {long: true, valid: false}, "--}": {long: true, valid: false},
			"--;": {long: true, valid: false}, "--[": {long: true, valid: false},
			"--:": {long: true, valid: false}, "--]": {long: true, valid: false},
			"--'": {long: true, valid: false}, "--\"": {long: true, valid: false},
			"--`": {long: true, valid: false}, "--~": {long: true, valid: false},
			"--<": {long: true, valid: false}, "-->": {long: true, valid: false},
			"--,": {long: true, valid: false}, "--.": {long: true, valid: false},
			"--\\": {long: true, valid: false}, "--?": {long: true, valid: false},
			"--/": {long: true, valid: false}, "-- ": {long: true, valid: false},
			"--foo_bar": {long: true, valid: true}, "--foo-bar": {long: true, valid: true},
			"--foo.bar": {long: true, valid: true}, "--foo;bar": {long: true, valid: false},
			"--foobar_": {long: true, valid: false}, "--foobar,": {long: true, valid: false},
			"--foobar`": {long: true, valid: false}, "--foobar$": {long: true, valid: false},
			"--foobar~": {long: true, valid: false}, "--foobar%": {long: true, valid: false},
			"--foobar!": {long: true, valid: false}, "--foobar^": {long: true, valid: false},
			"--foobar@": {long: true, valid: false}, "--foobar&": {long: true, valid: false},
			"--foobar#": {long: true, valid: false}, "--foobar*": {long: true, valid: false},
			"--foobar(": {long: true, valid: false}, "--foobar{": {long: true, valid: false},
			"--foobar)": {long: true, valid: false}, "--foobar}": {long: true, valid: false},
			"--foobar-": {long: true, valid: false}, "--foobar[": {long: true, valid: false},
			"--foobar+": {long: true, valid: false}, "--foobar]": {long: true, valid: false},
			"--foobar=": {long: true, valid: false}, "--foobar|": {long: true, valid: false},
			"--foobar'": {long: true, valid: false}, "--foobar\\": {long: true, valid: false},
			"--foobar?": {long: true, valid: false}, "--foobar;": {long: true, valid: false},
			"--foobar/": {long: true, valid: false}, "--foobar:": {long: true, valid: false},
			"--foobar>": {long: true, valid: false}, "--foobar.": {long: true, valid: false},
			"--foobar<": {long: true, valid: false}, "--foobar ": {long: true, valid: true},
			"--~foobar": {long: true, valid: false}, "--}foobar": {long: true, valid: false},
			"--:foobar": {long: true, valid: false}, "--|foobar": {long: true, valid: false},
			"--!foobar": {long: true, valid: false}, "--\\foobar": {long: true, valid: false},
			"--@foobar": {long: true, valid: false}, "--;foobar": {long: true, valid: false},
			"--#foobar": {long: true, valid: false}, "--'foobar": {long: true, valid: false},
			"--$foobar": {long: true, valid: false}, "--\"foobar": {long: true, valid: false},
			"--%foobar": {long: true, valid: false}, "--<foobar": {long: true, valid: false},
			"--^foobar": {long: true, valid: false}, "-->foobar": {long: true, valid: false},
			"--&foobar": {long: true, valid: false}, "--,foobar": {long: true, valid: false},
			"--*foobar": {long: true, valid: false}, "--.foobar": {long: true, valid: false},
			"--(foobar": {long: true, valid: false}, "--?foobar": {long: true, valid: false},
			"--)foobar": {long: true, valid: false}, "--/foobar": {long: true, valid: false},
			"--_foobar": {long: true, valid: false}, "-- foobar": {long: true, valid: false},
			"--+foobar": {long: true, valid: false}, "--1foobar": {long: true, valid: false},
			"---foobar": {long: true, valid: false}, "--2foobar": {long: true, valid: false},
			"--=foobar": {long: true, valid: false}, "--3foobar": {long: true, valid: false},
			"--{foobar": {long: true, valid: false}, "--4foobar": {long: true, valid: false},
			"--5foobar": {long: true, valid: false}, "--8foobar": {long: true, valid: false},
			"--6foobar": {long: true, valid: false}, "--9foobar": {long: true, valid: false},
			"--7foobar": {long: true, valid: false}, "--0foobar": {long: true, valid: false},
			"--foo~bar": {long: true, valid: false}, "--foo:bar": {long: true, valid: false},
			"--foo`bar": {long: true, valid: false}, "--foo'bar": {long: true, valid: false},
			"--foo!bar": {long: true, valid: false}, "--foo\"bar": {long: true, valid: false},
			"--foo@bar": {long: true, valid: false}, "--foo?bar": {long: true, valid: false},
			"--foo#bar": {long: true, valid: false}, "--foo/bar": {long: true, valid: false},
			"--foo$bar": {long: true, valid: false}, "--foo>bar": {long: true, valid: false},
			"--foo%bar": {long: true, valid: false}, "--foo7bar0": {long: true, valid: true},
			"--foo^bar": {long: true, valid: false}, "--foo<bar": {long: true, valid: false},
			"--foo&bar": {long: true, valid: false}, "--foo,bar": {long: true, valid: false},
			"--foo*bar": {long: true, valid: false}, "--foo|bar": {long: true, valid: false},
			"--foo(bar": {long: true, valid: false}, "--foo\\bar": {long: true, valid: true},
			"--foo)bar": {long: true, valid: false}, "--foo1bar": {long: true, valid: true},
			"--foo+bar": {long: true, valid: false}, "--foo2bar": {long: true, valid: true},
			"--foo=bar": {long: true, valid: false}, "--foo3bar": {long: true, valid: true},
			"--foo[bar": {long: true, valid: false}, "--foo4bar": {long: true, valid: true},
			"--foo]bar": {long: true, valid: false}, "--foo5bar": {long: true, valid: true},
			"--foo{bar": {long: true, valid: false}, "--foo6bar": {long: true, valid: true},
			"--foo}bar": {long: true, valid: false}, "--foo7bar": {long: true, valid: true},
			"--foo8bar": {long: true, valid: true}, "--foo9bar": {long: true, valid: true},
			"--foo0bar": {long: true, valid: true}, "--foo7bar1": {long: true, valid: true},
			"--foo7bar6": {long: true, valid: true}, "--foo7bar2": {long: true, valid: true},
			"--foo7bar7": {long: true, valid: true}, "--foo7bar3": {long: true, valid: true},
			"--foo7bar8": {long: true, valid: true}, "--foo7bar4": {long: true, valid: true},
			"--foo7bar9": {long: true, valid: true}, "--foo7bar5": {long: true, valid: true},
			"--foobar": {long: true, valid: true}, " --foobar": {long: true, valid: true},
			" -f": {short: true, valid: true}, " -f  ": {short: true, valid: true},
		}

		for arg, expectation := range testData {
			var cmd Commandline
			if expectation.short {
				cmd.Short = ShortArgumentString(arg)
			}
			if expectation.long {
				cmd.Long = LongArgumentString(arg)
			}

			err := cmd.Verify(nil)
			if expectation.valid {

				if err == nil {
					continue
				} else {
					t.Fatalf("expected valid, got error: {short:'%s', long:'%s', error: %v}",
						cmd.Short, cmd.Long, err)
				}

			} else { //expect invalid / error

				if err == nil {
					t.Fatalf("expect error. got none: {short:'%s', long:'%s', error: %v}",
						cmd.Short, cmd.Long, err)
				} else {
					continue
				}

			}

		}
	})
}
