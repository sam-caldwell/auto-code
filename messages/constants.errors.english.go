package messages

//ToDo: add build configs so this only builds if english is specified to allow language portability

const (
	ErrDefaultValueTypeMismatch        = `config.property.Type must match config.property.default (property:%s)`
	ErrDuplicateConfigSource           = `config.sources cannot contain duplicate elements`
	ErrEmptyGitRepoUrl                 = `empty git repo url`
	ErrEmptyRegularExpression          = `empty regular expression is not allowed`
	ErrExpectedEmptyValidatorParameter = `config.properties[].validator should be nil if class is null (property:%s)`
	ErrExpectedRegexString             = `expected regular expression string (property: %s)`
	ErrGitInitFailed                   = `'git init' failed: %v`
	ErrGitRemoteAddFailed              = `'git remote add' failed: %v`
	ErrInvalidConfigSource             = `unsupported or unknown config.source value (expect: %s)`
	ErrInvalidCopyrightPattern         = `global copyright string must be properly formatted`
	ErrInvalidEnvironmentVariableName  = `invalid environment variable name (%s)`
	ErrInvalidEmailAddress             = `email address must be valid`
	ErrInvalidFileName                 = `invalid file name (%s)`
	ErrInvalidFileFormat               = `invalid file format (%s)`
	ErrInvalidGitRepoUrl               = `invalid git url (%s)`
	ErrInvalidLanguagePattern          = `language must be supported by auto-code`
	ErrInvalidLicensePattern           = `license must match pattern (%s)`
	ErrInvalidMinMaxValue              = `min/max values must be numeric where min<max`
	ErrInvalidNonEmptyText             = `a non-empty/whitespace text string is expected`
	ErrInvalidNameIdentifier           = `name must match pattern (%s)`
	ErrInvalidObjectPropertyString     = `invalid object property string (%s)`
	ErrInvalidLongArgument             = `long argument '%s' is not valid`
	ErrInvalidShortArgument            = `short argument '%s' is not valid`
	ErrInvalidValidatorRegex           = `invalid validator regular expression (property: %s, regex: %s, err: %v)`
	ErrInvalidVersionPattern           = `version must match pattern (%s)`
	ErrGitRepoUrlMismatch              = `git repo url mismatch (local != remote)`
	ErrMissingCommandlineArgument      = `missing argument: must have either short, long or both argument type`
	ErrMissingConfigProperties         = `config.properties must have at least one element`
	ErrMissingConfigSource             = `config.sources must have at least one element`
	ErrMissingLocalGitRepoUrl          = `missing GitRepoUrl (git remote -v)`
	ErrMissingGitRepoUrl               = `missing local repo and manifest (global.git_repo): '%v'`
	ErrGitRemoteUrlFailed              = `'git remote -v' failed: 'Error:%v'`
	ErrTypeMismatchMin                 = `type mismatch(Min)`
	ErrTypeMismatchMax                 = `type mismatch(Max)`
	ErrUnknownProperty                 = `unknown property (%s)`
	ErrUnsupportedPropertyType         = `unsupported property type for '%s'`
	ErrUnsupportedPropertyValidator    = `unsupported property validator for '%s'`
)
