# zapgorm2

:smile: zapgorm2 is a zap logging driver for gorm v2

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/zapgorm2)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/zapgorm2/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/zapgorm2.svg)](https://github.com/moul/zapgorm2/releases)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![CI](https://github.com/moul/zapgorm2/workflows/CI/badge.svg)](https://github.com/moul/zapgorm2/actions?query=workflow%3ACI)
[![Release](https://github.com/moul/zapgorm2/workflows/Release/badge.svg)](https://github.com/moul/zapgorm2/actions?query=workflow%3ARelease)
[![GolangCI](https://golangci.com/badges/github.com/moul/zapgorm2.svg)](https://golangci.com/r/github.com/moul/zapgorm2)
[![codecov](https://codecov.io/gh/moul/zapgorm2/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/zapgorm2)
[![Go Report Card](https://goreportcard.com/badge/moul.io/zapgorm2)](https://goreportcard.com/report/moul.io/zapgorm2)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/zapgorm2/badge)](https://www.codefactor.io/repository/github/moul/zapgorm2)

If you're using gorm v1, you can use https://github.com/moul/zapgorm instead.

## Usage

```go
import "moul.io/zapgorm2"

logger := zapgorm2.New(zap.L())
logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
db, err = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{Logger: logger})
```

## Install

### Using go

```console
$ go get -u moul.io/zapgorm2
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/zapgorm2.svg)](https://starchart.cc/moul/zapgorm2)

## License

© 2020-2021 [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
