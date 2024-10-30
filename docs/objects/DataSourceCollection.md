Object: DataSource
==================

## Description

The `DataSourceCollection` is a list of pointers to the `DataSource` objects which
define the sources from which data values can be extracted by `Parser()`. This
must use a list since the order in which each data source is processed must be
preserved to ensure the last defined `DataSource` has precedence to overwrite the
final value.

## Structure

```text
  DataSourceCollection []*DataSource
```

## Dependencies:

* [DataSource](../interfaces/DataSource.md)

