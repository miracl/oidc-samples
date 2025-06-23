# Milagro crypto Go wrapper

## Build

The package by default is build with all supported curves. To change that you
can use the name of the curve as build tag. Special tags:

- ignoredefaultcurves - ignores the default set of curves and uses only the one
  specified with tags

### All curves

```bash
go build
```

### Build with subset curve

```bash
go build -tags 'ignoredefaultcurves ANSII'
```
