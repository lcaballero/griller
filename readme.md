# Introduction

`griller` creates a starter Go project with some basic structure.
Basic structure includes the start of command line processing using
go-flags.

## Installation

```
%> go get -u github.com/lcaballero/griller
```

## Usage

```
griller --name=newlib --remote github.com/saber --dest .
```

A directory for the project will be created in the current directory
containing the boiler plate code for the project.

A ~/.griller JSON file can be created with the default values for the
--dest and --remote flags.  For example:

```
{
  "Remote": "github.com/saber",
  "Dest": "$GOPATH/src/github.com/saber"
}
```

Once the .griller file is present you can then run (assuming that
the griller executable is on the PATH):

```
griller --name=newlib
```

## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file
'license' at the root of this distribution. By using this software in
any fashion, you are agreeing to be bound by the terms of this
license. You must not remove this notice, or any other, from this
software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
