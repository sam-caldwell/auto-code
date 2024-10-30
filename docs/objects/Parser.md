Object: Parser
==============

## Description

The `Parser` object defines how the application will compile and validate its
data sources to create a configuration object. This will consist of global
program properties like `programName` as well as a collection (map) of parameter
definition objects.

## Structure

```text
Parser
  │
  ├── programName: string
  ├── programDescription: string
  ├── programCopyright: string
  └── parameters: ParameterCollection
```

## Dependencies:
* [ParameterCollection](./ParameterCollection.md)
