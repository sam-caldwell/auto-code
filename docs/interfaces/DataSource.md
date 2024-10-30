Object: DataSource
==================

## Description

A `DataSource` is an implementation of this standard interface. This may be an
object such as `YamlFile`, `JsonFile`, `KeyValueFile`, `Environment`, `CliOption`
or maybe even a network-aware `AwsSecretManager` or `OnePassword` secure vault
integration.

The MVP `DataSource` objects are--

* `dsCliOption`: Returns the corresponding command-line option flag value.

* `dsEnvironment`: Returns the value for a given environment variable value.

* `dsJsonFile`: Consumes a JSON configuration file and returns the appropriate
  parameter value.

* `dsKeyValueFile`: Consumes a key-value text file and returns a corresponding
  value for a given key.

* `dsYamlFile`: Consumes a YAML configuration file and returns the appropriate
  parameter value.

## Structure

```text
  DataSource (interface)
```

## Dependencies:

* [dsCliOption](../objects/dsCliOption.md)

* [dsEnvironment](../objects/dsEnvironment.md)

* [dsJsonFile](../objects/dsJsonFile.md)

* [dsKeyValueFile](../objects/dsKeyValueFile.md)

* [dsYamlFile](../objects/dsYamlFile.md)



