Object: ParameterCollection
===========================

## Description

The `ParameterCollection` object is a map of a `ParameterName` object to the associated `ParameterDefinition`, 
where `ParameterName` is the name used by the application to access a given parameter's `Parser` definition or final
`Configuration` value.

This map associates the ParameterName with a pointer to the ParameterDefinition, keeping things simpler for updates.

## Structure 
```text
  ParameterCollection: map(ParameterName -> *ParameterDefinition)
```

## Dependencies: 
  * [ParameterName](./ParameterName.md)
  * [ParameterDefinition](./ParameterDefinition.md)

