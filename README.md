<h1 align="center">Config</h1>
<p>
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> A package to handle merged config from the environment & JSON files.

## Example

```go
// Define a single config struct as a mixture of JSON and environment
// based parameters.
// 
// Validation is handled by https://github.com/go-playground/validator
type OurConfig struct {
    FromJSON `json:"from_json" validate:"required"`
    FromEnv `env:"FROM_ENV" validate:"required"`
}

// Specify the location of your configuration files, and the current
// environment.
//
// config will search for a file within Path that matches `config.<environment>.json`
var cfg OurConfig
err := config.NewFromFile(&config.Config{
    Path:        "./config",
    Environment: config.Local,
}, &cfg)

// err will hold any validation errors.
```