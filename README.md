# bob-logging

## Description

This repository is a logging package created using ubuer/zap.

## Install
Using this package is straightforward. First, you install the package using <code>go get</code>.


```sh

go get -u github.com/tae2089/bob-logging

```

## Usage

```go
package main

import "github.com/tae2089/bob-logging/logger"

func main() {

  logger.Debug("Debug!!")

  logger.Info("Info!!")
	
  logger.Warn("Warn!!")
	
  logger.Error("Error!!")
  
}
```

## Author
👤 tae2089 tae2089002@gmail.com

Github: @tae2089

## 🤝 Contributing
Contributions, issues and feature requests are welcome!
Feel free to check issues page.
