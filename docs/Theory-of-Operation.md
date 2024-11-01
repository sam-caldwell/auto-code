Theory of Operation
===================

Given a manifest YAML file (below), we should be able to generate code in the language specified in `globals.language`
which will--
- include a `main()` entrypoint function with a set of global constants.
- create a `-h` and `--help` argument to print usage information formatted as:
  ```text
  #program -h
  PROGRAM_NAME (PROGRAM_AUTHOR)
  PROGRAM_DESCRIPTION
  
  Usage:

      PROGRAM_NAME [options]
  
  options:
    -h | --help          : return help
    -v | --version       : print the program version
    -c | --config <file> : identifies a configuration file
    -d | --debug         : print debug messages
    <...any defined program command-line arguments...>  
  
  PROGRAM_COPYRIGHT  (PROGRAM_LICENSE)
  ```
- create a `-v` and `--version` argument to print `global.version` to stdout.
- create an extensible `configuration` software package/module which defines how the program will be configured
  using configuration files, environment variables, command-line arguments and other sources.  This should include--`
    - a main `Configuration` class/struct/object which can be used by the larger program to access configuration
      properties as key-value pairs, accessed through a getter function.
        - The `Configuration` object should initialize and load all configuration files, environment variables and other
          data sources (validate them and merge them in source order) during startup without being invoked in the `main()`
          function.
        - Further, the generated program should make `Configuration` available globally to the program.
    - a `Validator` class/struct/object, defined by the manifest YAML file, which will validate that a given property
      has the expected data type and a value matching the expected value parameters.
    - a `configFile` module/package which can read a JSON/YAML configuration file and return a map of key-value pairs
      based on the YAML manifest file mapping definition
    - an `environment` module/package, which can read a list of environment variable names and return a map of
      corresponding key-value pairs containing the values of any environment variables.
    - a `commandline` module/package which can parse the command-line arguments passed to a program and return the
      result as a map of key-value pairs where each map is based on the YAML manifest file.
- ensure that `configuration.Configuration` uses `configuration.Validator` to validate the key-value pairs received by
  the sources defined in manifest YAML (e.g. `file` using `configFile`, `environment`, or `commandline`) and merge them
  in the order of the `sources` field of the manifest YAML.
- Finally, if executed with a `--debug` flag, the program should start by dumping the key-value pairs contained in the
  `Configuration` object.

This is an example `manifest.yaml` file (referred to above):
---
#
# This YAML manifest defines a software project and should be consumable by a code generator to produce the
# project's basic implementation.
#

#
# global:
#   defines program global constants used to identify the program.
#
global:
# name must be non-empty, alphanumeric (with _ and . also allowed) and not contain spaces.
name: myDemo
# description must be a non-empty string containing unicode characters
description: This is a demonstration of the code generator
# author must be a non-empty string
author: Sam Caldwell
# email must be a valid email address
email: scaldwell@asymmetric-effort.com
# copyright must be a valid US Copyright string
copyright: (c) 2024 Asymmetric Effort, LLC
# license must be a valid license type (proprietary, MIT, BSD, etc.)
license: MIT
# language must be one of the supported languages (e.g. go)
language: go
# version must follow semantic version formatting
version: 0.0.1
#
# Config section:
#   defines the Configuration struct used to configure the program runtime using a variety of
#   sources (files, environment variables, cli, etc.)
#
config:
#
# sources:
#   define the input sources used to configure the application in the order which they will be applied,
#   such that the last source will override its predecessors. Each source must be defined in a named
#   subsection (e.g. file, environment, commandline)
#
#   each source type can only appear once in the list.
#
sources:
- file
- environment
- commandline
#
# properties:
#   define the configuration object's properties.  This will result in a property and a getter/setter function pair
#   as well as property type and validator specification to ensure data quality.
#
#   property type must be one of string, int, uint, etc.
#
properties:
- listenAddress:
type: string
validator:
class: pattern
parameter: <ip address regex> # ToDo: create an ip address regex
default: 0.0.0.0
- listenPort:
type: uint16
validator:
class: null
default: 8080
- protocol:
type: string
validator:
class: pattern
parameter: `^http[s]{0,1}$`
default: http
#
# file:
#   defines a configuration file data source by fileName, format (e.g. yaml) and a mapping of the yaml file content
#   to a property listed in the properties section.  When the file is read, property values will be stored as the
#   associated (mapped) property name and validated using the property validator and typing.
#
file:
file: config.yaml
# if required==true, then the file must exist or an error will be thrown.
required: true
# only yaml is supported at this time, but this should be extensible to include json in the future.
format: yaml
# the `map` associates a `property` to a config file record, where the config file record uses a dot-delimited
# object notation.
map:
- listenAddress: network.address
- listenPort: network.port
- protocol: network.protocol
#
# environment:
#   defines one or more environment variables whose values map to configuration properties.
#   An environment variable is applied only if it is non-empty.
#
environment:
- listenAddress: LISTEN_ADDRESS
- listenPort: LISTEN_PORT
- protocol: LISTEN_PROTOCOL
#
# commandline:
#   defines one or more command-line arguments and maps them to a parameter.  This list is a list of properties
#   which has one or both of the short and long argument types.  If the `required` property is true and the program
#   runs without this commandline argument, an error should be thrown.
#
commandline:
- listenAddress:
required: false
short: -a
long: --listen-address
- listenPort:
required: false
short: -p
long: --listen-port
- protocol:
required: false
long: --protocol
