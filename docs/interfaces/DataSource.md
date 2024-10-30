Object: DataSource
==================

## Description

A `DataSource` is an implementation of this standard interface. This may be an
object such as `YamlFile`, `JsonFile`, `KeyValueFile`, `Environment`, `CliOption`
or maybe even a network-aware `AwsSecretManager` or `OnePassword` secure vault
integration.

The MVP `DataSource` objects are--

* `YamlFile`: Consumes a YAML configuration file and returns the appropriate
  parameter value.

* `JsonFile`: Consumes a JSON configuration file and returns the appropriate
  parameter value.

* `KeyValueFile`: Consumes a key-value text file and returns a corresponding
  value for a given key.

* `Environment`: Returns the value for a given environment variable value.

* `CliOption`: Returns the corresponding command-line option flag value.

## Structure

```text
  DataSource (interface)
```

## Dependencies:

* [YamlFile](../objects/YamlFile.md)

* [JsonFile](../objects/JsonFile.md)

* [KeyValueFile](../objects/KeyValueFile.md)

* [Environment](../objects/Environment.md)

* [CliOption](../objects/CliOption.md)

