# gocmt

go commit lint. Based on Angular commit conventions. Forces to use the following format when committing a message.

```
<type>: <short summary>
```

LINK: [AngularJS Git Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit?usp=sharing)

## Usage

```
$ go install github.com/chaewonkong/gocmt@latest
```

Run:

```
$ gocmt init
```

## Supported types

- feat
- fix
- docs
- style
- refactor
- test
- build
- ci
- perf
- chore
