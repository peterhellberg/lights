# 💡💡 `lights`

A command line tool for circadian lighting at my desk.

> Note: This code is meant for my own use, written as a lab day project, do not expect much.

## Circadian Lighting

![Circadian Lighting Diagram](https://aws1.discourse-cdn.com/free1/uploads/wled/optimized/2X/2/24c03aa4e803b47a723a0a28737b7047fffb2cfd_2_1380x780.jpeg)

![Temperature Chart](https://i.imgur.com/NiAdlGj.png)

## Usage

```
lights      # Toggles both of the lights ON or OFF
lights -i   # Show information about the current state
```

```go
$ lights --help
Usage of lights:
  -af string
    	the address of the Fill Light's HTTP API (default "http://filllight:9123")
  -ak string
    	the address of the Key Light's HTTP API (default "http://keylight:9123")
  -bf value
    	set Fill Light brightness to an absolute (between 0 and 100) or relative (-N or +N) percentage
  -bk value
    	set Key Light brightness to an absolute (between 0 and 100) or relative (-N or +N) percentage
  -c	calculate and set the appropriate circadian lighting values
  -i	display the current status of the lights without changing their state
  -tf value
    	set Fill Light temperature to an absolute (between 2900 and 7000) or relative (-N or +N) degrees
  -tk value
    	set Key Light temperature to an absolute (between 2900 and 7000) or relative (-N or +N) degrees
```

> I have also bound `Ctrl+F11` to toggle the Key Light, and `Ctrl+F12` to toggle the Fill Light.

## Technical details

The heavy lifting is performed by [github.com/mdlayher/keylight](https://github.com/mdlayher/keylight/)

<img src="https://assets.c7.se/svg/viking-gopher.svg" align="right" width="30%" height="300">

## License (MIT)

Copyright (c) 2022 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
