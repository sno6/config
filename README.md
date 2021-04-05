<h1 align="center">Config 📃</h1>
<p>
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> Easily handle JSON and environment based config in one.

## Example

```go
// Define your projects configuration.
// 
// Validation is handled by https://github.com/go-playground/validator
type OurConfig struct {
    FromJSON string `json:"from_json" validate:"required"`
    FromEnv  string `env:"FROM_ENV" validate:"required"`
}

// Specify where to find your projects configuration files.
var cfg OurConfig
err := config.NewFromFile(&config.Config{
    Path:        "./config",
    Environment: config.Local,
}, &cfg)

// If there were no validation errors cfg will now be populated with values from 
// the local JSON file as well as the environment.
```

More examples can be found [`here.`](/examples)

### Dependencies

[`go-env`](https://github.com/syndbg/goenv) For marshalling from the environment. </br>
[`validator`](https://github.com/go-playground/validator) For validation.