# Milagro crypto Go wrapper

[![pipeline status](https://gitlab.corp.miracl.com/mfa/pkg/gomiracl/badges/master/pipeline.svg)](https://gitlab.corp.miracl.com/mfa/pkg/gomiracl/commits/master)
[![coverage report](https://gitlab.corp.miracl.com/mfa/pkg/gomiracl/badges/master/coverage.svg)](https://gitlab.corp.miracl.com/mfa/pkg/gomiracl/commits/master)

# Build

The package by default is build with all supported curves. To change that you can use the name of the curve as build tag.
Special tags:
- ignoredefaultcurves - ignores the default set of curves and uses only the one specified with tags


## All curves
```
go build
```

## Build with subset curve
```
go build -tags 'ignoredefaultcurves ANSII'
```
