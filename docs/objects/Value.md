Object: Value
=============

## Description

The `Value` object defines the default value (and later actual value during the
`Parser.Parse()` method's execution as well as a `Validator` object pointer.

## Structure

```text
  Value
    │
    ├──data: any
    └──validator: *Validator
```

## Dependencies:

* [Validator](./Validator.md)

