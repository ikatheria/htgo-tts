![HTGO-TTS](https://banners.beyondco.de/HTGO-TTS.png?theme=light&packageManager=&packageName=go+get+%22github.com%2Fhegedustibor%2Fhtgo-tts%22&pattern=bamboo&style=style_1&description=Text+to+Speech+Package+for+GoLang&md=1&showWatermark=0&fontSize=100px&images=volume-up)

# htgo-tts
[https://hegedustibor.github.io/htgo-tts/](https://hegedustibor.github.io/htgo-tts/)

### Requirement:
- mplayer (optional)

### Install
```
go get "github.com/ikatheria/htgo-tts"
```

### Update
```
go get -u "github.com/ikatheria/htgo-tts"
```

### Remove
```
go clean -i "github.com/ikatheria/htgo-tts"
```

### Import
```go
import "github.com/ikatheria/htgo-tts"
import "github.com/ikatheria/htgo-tts/voices"
```

### Use
```go
speech := htgotts.Speech{Folder: "audio", Language: voices.English}
speech.Speak("Your sentence.")
```

### Use with Handlers
```go
import (
    htgotts "github.com/ikatheria/htgo-tts"
    handlers "github.com/ikatheria/htgo-tts/handlers"
    voices "github.com/ikatheria/htgo-tts/voices"
)

speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.MPlayer{}}
speech.Speak("Your sentence.")
```

### Use tts without external player

Note: The "Native" handler works around the Oto librery, see the ["Prerequisite"](https://github.com/hajimehoshi/oto/blob/main/README.md#prerequisite) section to work with this handler 

```go
import (
    htgotts "github.com/ikatherisa/htgo-tts"
    handlers "github.com/ikatheria/htgo-tts/handlers"
    voices "github.com/ikatheria/htgo-tts/voices"
)

speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
speech.Speak("Your sentence.")
```

### Use with Proxy
```go
import (
    htgotts "github.com/ikatheria/htgo-tts"
    handlers "github.com/ikatheria/htgo-tts/handlers"
    voices "github.com/ikatheria/htgo-tts/voices"
)

speech := htgotts.Speech{Folder: "audio", Language: voices.English, Proxy: "https://..."}
speech.Speak("Your sentence.")
```

### Support and Contributions

If you encounter issues using HTGO-TTS or would like to suggest improvements to the source code, you can create an issue on the ["Issues"](https://github.com/hegedustibor/htgo-tts/issues) tab. If you'd like to contribute to the HTGO-TTS source code, please submit a pull request.

### License

HTGO-TTS is free software and is available under the MIT license. For more information, please see the LICENSE file in the source code repository.


Have Fun!
