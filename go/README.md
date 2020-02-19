# goImage

This Go package implements the host-side of the Flutter [goImage](https://github.com/itMicmouse/goImage) plugin.

## Usage

Import as:

```go
import goImage "github.com/itMicmouse/goImage/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&goImage.GoImagePlugin{}),
```
