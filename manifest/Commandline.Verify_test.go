package manifest

import "testing"

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
			if msg := err.Error(); msg != errMissingCommandlineArgument {
				t.Fatalf("Verify() error mismatch.\n"+
					"Got %s\n"+
					"Want %s",
					msg, errMissingCommandlineArgument)
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
			"-a": TestData{short: true, valid: true}, "-b": TestData{short: true, valid: true},
			"-c": TestData{short: true, valid: true}, "-d": TestData{short: true, valid: true},
			"-e": TestData{short: true, valid: true}, "-f": TestData{short: true, valid: true},
			"-g": TestData{short: true, valid: true}, "-h": TestData{short: true, valid: true},
			"-i": TestData{short: true, valid: true}, "-j": TestData{short: true, valid: true},
			"-k": TestData{short: true, valid: true}, "-l": TestData{short: true, valid: true},
			"-m": TestData{short: true, valid: true}, "-n": TestData{short: true, valid: true},
			"-o": TestData{short: true, valid: true}, "-p": TestData{short: true, valid: true},
			"-q": TestData{short: true, valid: true}, "-r": TestData{short: true, valid: true},
			"-s": TestData{short: true, valid: true}, "-t": TestData{short: true, valid: true},
			"-u": TestData{short: true, valid: true}, "-v": TestData{short: true, valid: true},
			"-w": TestData{short: true, valid: true}, "-x": TestData{short: true, valid: true},
			"-y": TestData{short: true, valid: true}, "-z": TestData{short: true, valid: true},
			"-A": TestData{short: true, valid: true}, "-B": TestData{short: true, valid: true},
			"-C": TestData{short: true, valid: true}, "-D": TestData{short: true, valid: true},
			"-E": TestData{short: true, valid: true}, "-F": TestData{short: true, valid: true},
			"-G": TestData{short: true, valid: true}, "-H": TestData{short: true, valid: true},
			"-I": TestData{short: true, valid: true}, "-J": TestData{short: true, valid: true},
			"-K": TestData{short: true, valid: true}, "-L": TestData{short: true, valid: true},
			"-M": TestData{short: true, valid: true}, "-N": TestData{short: true, valid: true},
			"-O": TestData{short: true, valid: true}, "-P": TestData{short: true, valid: true},
			"-Q": TestData{short: true, valid: true}, "-R": TestData{short: true, valid: true},
			"-S": TestData{short: true, valid: true}, "-T": TestData{short: true, valid: true},
			"-U": TestData{short: true, valid: true}, "-V": TestData{short: true, valid: true},
			"-W": TestData{short: true, valid: true}, "-X": TestData{short: true, valid: true},
			"-Y": TestData{short: true, valid: true}, "-Z": TestData{short: true, valid: true},
			"-0": TestData{short: true, valid: true}, "-1": TestData{short: true, valid: true},
			"-2": TestData{short: true, valid: true}, "-3": TestData{short: true, valid: true},
			"-4": TestData{short: true, valid: true}, "-5": TestData{short: true, valid: true},
			"-6": TestData{short: true, valid: true}, "-7": TestData{short: true, valid: true},
			"-8": TestData{short: true, valid: true}, "-9": TestData{short: true, valid: true},
			"-_": TestData{short: true, valid: false}, "-": TestData{short: true, valid: false},
			"-!": TestData{short: true, valid: false}, "-*": TestData{short: true, valid: false},
			"-@": TestData{short: true, valid: false}, "-(": TestData{short: true, valid: false},
			"-#": TestData{short: true, valid: false}, "-)": TestData{short: true, valid: false},
			"-$": TestData{short: true, valid: false}, "-+": TestData{short: true, valid: false},
			"-%": TestData{short: true, valid: false}, "-{": TestData{short: true, valid: false},
			"-^": TestData{short: true, valid: false}, "-}": TestData{short: true, valid: false},
			"-;": TestData{short: true, valid: false}, "-[": TestData{short: true, valid: false},
			"-:": TestData{short: true, valid: false}, "-]": TestData{short: true, valid: false},
			"-'": TestData{short: true, valid: false}, "-\"": TestData{short: true, valid: false},
			"-`": TestData{short: true, valid: false}, "-~": TestData{short: true, valid: false},
			"-<": TestData{short: true, valid: false}, "->": TestData{short: true, valid: false},
			"-,": TestData{short: true, valid: false}, "-.": TestData{short: true, valid: false},
			"-\\": TestData{short: true, valid: false}, "-?": TestData{short: true, valid: false},
			"-/": TestData{short: true, valid: false}, "--": TestData{short: true, long: true, valid: false},
			"--a": TestData{long: true, valid: false}, "--b": TestData{long: true, valid: false},
			"--c": TestData{long: true, valid: false}, "--d": TestData{long: true, valid: false},
			"--e": TestData{long: true, valid: false}, "--f": TestData{long: true, valid: false},
			"--g": TestData{long: true, valid: false}, "--h": TestData{long: true, valid: false},
			"--i": TestData{long: true, valid: false}, "--j": TestData{long: true, valid: false},
			"--k": TestData{long: true, valid: false}, "--l": TestData{long: true, valid: false},
			"--m": TestData{long: true, valid: false}, "--n": TestData{long: true, valid: false},
			"--o": TestData{long: true, valid: false}, "--p": TestData{long: true, valid: false},
			"--q": TestData{long: true, valid: false}, "--r": TestData{long: true, valid: false},
			"--s": TestData{long: true, valid: false}, "--t": TestData{long: true, valid: false},
			"--u": TestData{long: true, valid: false}, "--v": TestData{long: true, valid: false},
			"--w": TestData{long: true, valid: false}, "--x": TestData{long: true, valid: false},
			"--y": TestData{long: true, valid: false}, "--z": TestData{long: true, valid: false},
			"--A": TestData{long: true, valid: false}, "--B": TestData{long: true, valid: false},
			"--C": TestData{long: true, valid: false}, "--D": TestData{long: true, valid: false},
			"--E": TestData{long: true, valid: false}, "--F": TestData{long: true, valid: false},
			"--G": TestData{long: true, valid: false}, "--H": TestData{long: true, valid: false},
			"--I": TestData{long: true, valid: false}, "--J": TestData{long: true, valid: false},
			"--K": TestData{long: true, valid: false}, "--L": TestData{long: true, valid: false},
			"--M": TestData{long: true, valid: false}, "--N": TestData{long: true, valid: false},
			"--O": TestData{long: true, valid: false}, "--P": TestData{long: true, valid: false},
			"--Q": TestData{long: true, valid: false}, "--R": TestData{long: true, valid: false},
			"--S": TestData{long: true, valid: false}, "--T": TestData{long: true, valid: false},
			"--U": TestData{long: true, valid: false}, "--V": TestData{long: true, valid: false},
			"--W": TestData{long: true, valid: false}, "--X": TestData{long: true, valid: false},
			"--Y": TestData{long: true, valid: false}, "--Z": TestData{long: true, valid: false},
			"--0": TestData{long: true, valid: false}, "--1": TestData{long: true, valid: false},
			"--2": TestData{long: true, valid: false}, "--3": TestData{long: true, valid: false},
			"--4": TestData{long: true, valid: false}, "--5": TestData{long: true, valid: false},
			"--6": TestData{long: true, valid: false}, "--7": TestData{long: true, valid: false},
			"--8": TestData{long: true, valid: false}, "--9": TestData{long: true, valid: false},
			"--_": TestData{long: true, valid: false}, "---": TestData{long: true, valid: false},
			"--!": TestData{long: true, valid: false}, "--*": TestData{long: true, valid: false},
			"--@": TestData{long: true, valid: false}, "--(": TestData{long: true, valid: false},
			"--#": TestData{long: true, valid: false}, "--)": TestData{long: true, valid: false},
			"--$": TestData{long: true, valid: false}, "--+": TestData{long: true, valid: false},
			"--%": TestData{long: true, valid: false}, "--{": TestData{long: true, valid: false},
			"--^": TestData{long: true, valid: false}, "--}": TestData{long: true, valid: false},
			"--;": TestData{long: true, valid: false}, "--[": TestData{long: true, valid: false},
			"--:": TestData{long: true, valid: false}, "--]": TestData{long: true, valid: false},
			"--'": TestData{long: true, valid: false}, "--\"": TestData{long: true, valid: false},
			"--`": TestData{long: true, valid: false}, "--~": TestData{long: true, valid: false},
			"--<": TestData{long: true, valid: false}, "-->": TestData{long: true, valid: false},
			"--,": TestData{long: true, valid: false}, "--.": TestData{long: true, valid: false},
			"--\\": TestData{long: true, valid: false}, "--?": TestData{long: true, valid: false},
			"--/": TestData{long: true, valid: false}, "-- ": TestData{long: true, valid: false},
			"--foo_bar": TestData{long: true, valid: true}, "--foo-bar": TestData{long: true, valid: true},
			"--foo.bar": TestData{long: true, valid: true}, "--foo;bar": TestData{long: true, valid: false},
			"--foobar_": TestData{long: true, valid: false}, "--foobar,": TestData{long: true, valid: false},
			"--foobar`": TestData{long: true, valid: false}, "--foobar$": TestData{long: true, valid: false},
			"--foobar~": TestData{long: true, valid: false}, "--foobar%": TestData{long: true, valid: false},
			"--foobar!": TestData{long: true, valid: false}, "--foobar^": TestData{long: true, valid: false},
			"--foobar@": TestData{long: true, valid: false}, "--foobar&": TestData{long: true, valid: false},
			"--foobar#": TestData{long: true, valid: false}, "--foobar*": TestData{long: true, valid: false},
			"--foobar(": TestData{long: true, valid: false}, "--foobar{": TestData{long: true, valid: false},
			"--foobar)": TestData{long: true, valid: false}, "--foobar}": TestData{long: true, valid: false},
			"--foobar-": TestData{long: true, valid: false}, "--foobar[": TestData{long: true, valid: false},
			"--foobar+": TestData{long: true, valid: false}, "--foobar]": TestData{long: true, valid: false},
			"--foobar=": TestData{long: true, valid: false}, "--foobar|": TestData{long: true, valid: false},
			"--foobar'": TestData{long: true, valid: false}, "--foobar\\": TestData{long: true, valid: false},
			"--foobar?": TestData{long: true, valid: false}, "--foobar;": TestData{long: true, valid: false},
			"--foobar/": TestData{long: true, valid: false}, "--foobar:": TestData{long: true, valid: false},
			"--foobar>": TestData{long: true, valid: false}, "--foobar.": TestData{long: true, valid: false},
			"--foobar<": TestData{long: true, valid: false}, "--foobar ": TestData{long: true, valid: true},
			"--~foobar": TestData{long: true, valid: false}, "--}foobar": TestData{long: true, valid: false},
			"--:foobar": TestData{long: true, valid: false}, "--|foobar": TestData{long: true, valid: false},
			"--!foobar": TestData{long: true, valid: false}, "--\\foobar": TestData{long: true, valid: false},
			"--@foobar": TestData{long: true, valid: false}, "--;foobar": TestData{long: true, valid: false},
			"--#foobar": TestData{long: true, valid: false}, "--'foobar": TestData{long: true, valid: false},
			"--$foobar": TestData{long: true, valid: false}, "--\"foobar": TestData{long: true, valid: false},
			"--%foobar": TestData{long: true, valid: false}, "--<foobar": TestData{long: true, valid: false},
			"--^foobar": TestData{long: true, valid: false}, "-->foobar": TestData{long: true, valid: false},
			"--&foobar": TestData{long: true, valid: false}, "--,foobar": TestData{long: true, valid: false},
			"--*foobar": TestData{long: true, valid: false}, "--.foobar": TestData{long: true, valid: false},
			"--(foobar": TestData{long: true, valid: false}, "--?foobar": TestData{long: true, valid: false},
			"--)foobar": TestData{long: true, valid: false}, "--/foobar": TestData{long: true, valid: false},
			"--_foobar": TestData{long: true, valid: false}, "-- foobar": TestData{long: true, valid: false},
			"--+foobar": TestData{long: true, valid: false}, "--1foobar": TestData{long: true, valid: false},
			"---foobar": TestData{long: true, valid: false}, "--2foobar": TestData{long: true, valid: false},
			"--=foobar": TestData{long: true, valid: false}, "--3foobar": TestData{long: true, valid: false},
			"--{foobar": TestData{long: true, valid: false}, "--4foobar": TestData{long: true, valid: false},
			"--5foobar": TestData{long: true, valid: false}, "--8foobar": TestData{long: true, valid: false},
			"--6foobar": TestData{long: true, valid: false}, "--9foobar": TestData{long: true, valid: false},
			"--7foobar": TestData{long: true, valid: false}, "--0foobar": TestData{long: true, valid: false},
			"--foo~bar": TestData{long: true, valid: false}, "--foo:bar": TestData{long: true, valid: false},
			"--foo`bar": TestData{long: true, valid: false}, "--foo'bar": TestData{long: true, valid: false},
			"--foo!bar": TestData{long: true, valid: false}, "--foo\"bar": TestData{long: true, valid: false},
			"--foo@bar": TestData{long: true, valid: false}, "--foo?bar": TestData{long: true, valid: false},
			"--foo#bar": TestData{long: true, valid: false}, "--foo/bar": TestData{long: true, valid: false},
			"--foo$bar": TestData{long: true, valid: false}, "--foo>bar": TestData{long: true, valid: false},
			"--foo%bar": TestData{long: true, valid: false}, "--foo7bar0": TestData{long: true, valid: true},
			"--foo^bar": TestData{long: true, valid: false}, "--foo<bar": TestData{long: true, valid: false},
			"--foo&bar": TestData{long: true, valid: false}, "--foo,bar": TestData{long: true, valid: false},
			"--foo*bar": TestData{long: true, valid: false}, "--foo|bar": TestData{long: true, valid: false},
			"--foo(bar": TestData{long: true, valid: false}, "--foo\\bar": TestData{long: true, valid: true},
			"--foo)bar": TestData{long: true, valid: false}, "--foo1bar": TestData{long: true, valid: true},
			"--foo+bar": TestData{long: true, valid: false}, "--foo2bar": TestData{long: true, valid: true},
			"--foo=bar": TestData{long: true, valid: false}, "--foo3bar": TestData{long: true, valid: true},
			"--foo[bar": TestData{long: true, valid: false}, "--foo4bar": TestData{long: true, valid: true},
			"--foo]bar": TestData{long: true, valid: false}, "--foo5bar": TestData{long: true, valid: true},
			"--foo{bar": TestData{long: true, valid: false}, "--foo6bar": TestData{long: true, valid: true},
			"--foo}bar": TestData{long: true, valid: false}, "--foo7bar": TestData{long: true, valid: true},
			"--foo8bar": TestData{long: true, valid: true}, "--foo9bar": TestData{long: true, valid: true},
			"--foo0bar": TestData{long: true, valid: true}, "--foo7bar1": TestData{long: true, valid: true},
			"--foo7bar6": TestData{long: true, valid: true}, "--foo7bar2": TestData{long: true, valid: true},
			"--foo7bar7": TestData{long: true, valid: true}, "--foo7bar3": TestData{long: true, valid: true},
			"--foo7bar8": TestData{long: true, valid: true}, "--foo7bar4": TestData{long: true, valid: true},
			"--foo7bar9": TestData{long: true, valid: true}, "--foo7bar5": TestData{long: true, valid: true},
			"--foobar": TestData{long: true, valid: true}, " --foobar": TestData{long: true, valid: true},
			" -f": TestData{short: true, valid: true}, " -f  ": TestData{short: true, valid: true},
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
