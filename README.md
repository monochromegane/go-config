# go-config

A go package for parsing JSON config file.

## Features

- Parsing JSON config file.
- Config with default value.

## Installation

```sh
$ go get github.com/monochromegane/go-config
```

## Usage

This package uses struct for storing parsed JSON config data.

```go
type Config struct {
	Name string `json:"name" default:"default name"`
}
```

```json
{
  "name": "name"
}
```
And, you can get config by the following.

```go
var conf Config
config.Parse("config.json", &conf)
```

## Tasks

- Show config description.
- Set default value to slice.

## Code status

- [![Build Status](https://travis-ci.org/monochromegane/go-config.svg?branch=master)](https://travis-ci.org/monochromegane/go-config)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

