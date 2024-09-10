# Project scouts-risk

One Paragraph of project description goes here

## Rough notes

### Get up and running

#### TailwindCSS

* Install the latest binary, instructions here
  * https://tailwindcss.com/blog/standalone-cli
  * BUT to include DaisyUI use the following binary instead
  * https://github.com/dobicinaitis/tailwind-cli-extra

### Template

* GoShip.it
  * Comes with Templ first components that employ DaisyUI components
    * https://goship.it
  * DaisyUI components (for reference)
    * https://daisyui.com/components

### Notion

* Don't forget to
  * [Create your integration](https://www.notion.so/profile/integrations)
  * Add your integration to the databases you want it to have access to

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```