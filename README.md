# kitty

Your regular crap bot you all know and love.

# Installation
1. First you'll need [Go](https://golang.org), 1.4+ should be fine.
2. Then setup your GOPATH to where ever you want it to go.
3. Run `go install github.com/acdenisSK/kitty` 
4. Go to `GOPATH/bin`
5. place `config.toml.example` there (don't forget to remove `.example`)
6. replace the required fields and *finally* run the binary.
7. You're good to go!

# Notes
- Please don't mind the unnecessary comments on some structs/functions. They're just there to comply with the linters, who unironically ask you to purposely add docs for publicly exported stuff.

# Credits
- [iopred/bruxism](https://github.com/iopred/bruxism) (septapus) for some parts of the stat command and motivations.
