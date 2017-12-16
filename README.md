<p align="center">
  <img src="https://i.genius.com/9e52af38092595b40bc8aba4b64ce227d8accf03?url=http%3A%2F%2Fcontent.screencast.com%2Fusers%2Fezrapgenius%2Ffolders%2FJing%2Fmedia%2F32a2be0d-58cd-49eb-a4f9-de3bbe77c964%2F00000265.png" width="300" height="150"><br><br>
  <b>gogenius</b> |
  <i>A Go client for the </i>
  <a href="http://genius.com/developers">Genius API</a>
</p>

## Usage
Just import it and use like this example:
```go
package main

import (
	"fmt"

	"github.com/erbesharat/gogenius"
)

func main() {
	fmt.Println(gogenius.GetReferents("10347"))
	fmt.Println(gogenius.GetArtist("16775"))
	fmt.Println(gogenius.GetSong("16775"))
	fmt.Println(gogenius.GetArtistSongs("16775"))
	fmt.Println(gogenius.GetWebpage("https://docs.genius.com"))
	fmt.Println(gogenius.Search("Another%20brick%20in%20the%20wall"))
	fmt.Println(gogenius.GetAnnotation("10225840"))
}
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## TODO
- [ ] Add JSON to struct feature for resources
- [ ] Organize code
- [ ] Write tests
- [ ] Write better documentation


## License
[GPLv3](https://www.gnu.org/licenses/gpl-3.0.txt)
