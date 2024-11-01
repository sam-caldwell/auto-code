package manifest

//ToDo: add build configs so this only builds if english is specified to allow language portability

const (
	errDefaultValueTypeMismatch        = `config.property.Type must match config.property.default (property:%s)`
	errDuplicateConfigSource           = `config.sources cannot contain duplicate elements`
	errEmptyGitRepoUrl                 = `empty git repo url`
	errExpectedEmptyValidatorParameter = `config.properties[].validator should be nil if class is null (property:%s)`
	errExpectedRegexString             = `expected regular expression string (property: %s)`
	errGitInitFailed                   = `'git init' failed: %v`
	errGitRemoteAddFailed              = `'git remote add' failed: %v`
	errInvalidConfigSource             = `unsupported or unknown config.source value (expect: %s)`
	errInvalidCopyrightPattern         = `global copyright string must be properly formatted`
	errInvalidEnvironmentVariableName  = `invalid environment variable name (%s)`
	errInvalidEmailAddress             = `email address must be valid`
	errInvalidFileName                 = `invalid file name (%s)`
	errInvalidFileFormat               = `invalid file format (%s)`
	errInvalidGitRepoUrl               = `invalid git url (%s)`
	errInvalidLanguagePattern          = `language must be supported by auto-code`
	errInvalidLicensePattern           = `license must match pattern (%s)`
	errInvalidMinMaxValue              = `min/max values must be numeric where min<max`
	errInvalidNonEmptyText             = `a non-empty/whitespace text string is expected`
	errInvalidNameIdentifier           = `name must match pattern (%s)`
	errInvalidObjectPropertyString     = `invalid object property string (%s)`
	errInvalidLongArgument             = `long argument '%s' is not valid`
	errInvalidShortArgument            = `short argument '%s' is not valid`
	errInvalidValidatorRegex           = `invalid validator regular expression (property: %s, regex: %s, err: %v)`
	errInvalidVersionPattern           = `version must match pattern (%s)`
	errGitRepoUrlMismatch              = `git repo url mismatch (local != remote)`
	errMissingCommandlineArgument      = `missing argument: must have either short, long or both argument type`
	errMissingConfigProperties         = `config.properties must have at least one element`
	errMissingConfigSource             = `config.sources must have at least one element`
	errMissingLocalGitRepoUrl          = `missing GitRepoUrl (git remote -v)`
	errMissingGitRepoUrl               = `missing local repo and manifest (global.git_repo): '%v'`
	errGitRemoteUrlFailed              = `'git remote -v' failed: 'Error:%v'`
	errUnknownProperty                 = `unknown property (%s)`
	errUnsupportedPropertyType         = `unsupported property type for '%s'`
	errUnsupportedPropertyValidator    = `unsupported property validator for '%s'`
)
