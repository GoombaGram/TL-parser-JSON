# TL-parser-JSON
Golang program that parses Telegram TL schema to JSON.


## Getting started
1. [Download latest release](https://github.com/ErikPelli/TL-parser-to-JSON/releases/latest)
   or get latest commit using `go get github.com/ErikPelli/TL-parser-to-JSON`

2. Download pre-compiled binary or use `go build` on source files folder.
   Open the .exe (or executable for Linux/MacOS) in a terminal and insert .tl file path.
   
## Result
You will find the JSON result on "result" folder in same directory (latest is currently inside).

JSON structure is identical to [Telegram TL JSON](https://core.telegram.org/schema/json), but, while they are stopped at a lower layer (actually they have JSON schema for layer 121), with this program you can parse a most recent TL layer taken from the [Telegram Desktop repository](https://github.com/telegramdesktop/tdesktop) (actually they have .tl schema of layer 133).
