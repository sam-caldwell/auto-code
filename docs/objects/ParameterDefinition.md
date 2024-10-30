Object: ParameterDefinition
===========================

## Description

The `ParameterDefinition` object defines the parameter, providing a helpful
`helpText` string, a `Value` object (see below) and a collection of data
sources (`DataSourceCollection`) from which the parameter's actual value
can be obtained.

## Structure

```text
  ParameterDefinition
    |
    +-->helpText:  HelpText
	+-->value: *Value
	+-->sources: DataSourceCollection
```

## Dependencies:

  * [HelpText](./HelpText.md)
  * [Value](./Value.md)
  * [DataSourceCollection](./DataSourceCollection.md)

