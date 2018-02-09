# golor
Rainbows as lolcat

# Screenshot

```
$fortune | cowsay | golor
 ___________________________________
/ After all, all he did was string  \
| together a lot of old, well-known |
| quotations.                       |
|                                   |
\ -- H.L. Mencken, on Shakespeare   /
 -----------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
$fortune | cowsay | golor aurora
 _________________________________________
/ It may be that your whole purpose in    \
| life is simply to serve as a warning to |
\ others.                                 /
 -----------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

![Screenshot](https://user-images.githubusercontent.com/10125386/36017805-9e80736e-0dbc-11e8-87be-69522e5d3677.png)


# Install

`go get -u github.com/khsk/golor`

or you can download binary https://github.com/khsk/golor/releases

# Help

```
NAME:
   golor - Colorize inputs

USAGE:
   golor [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     rainbow, r  Same columns, same colors
     aurora, a   Gradationed text
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input, -i, --in               Input from stdin
   --file value, -f value          Input from file
   --bold, -b                      Set bold
   --italic, --it                  Set italic
   --underline, -u                 Set underline
   --reverse, -r, --invert, --inv  Set reverse
   --cancel, -c                    Set cancealed
   --help, -h                      show help
   --version, -v                   print the version
```

```
NAME:
   golor rainbow - Same columns, same colors

USAGE:
   golor rainbow [command options] [arguments...]

OPTIONS:
   --deviation value, --dev value, -d value  Each line deviates from the color cycle (default: 0)
```

```
NAME:
   golor aurora - Gradationed text

USAGE:
   golor aurora [command options] [arguments...]

OPTIONS:
   --freq value, -f value    Aurora frequency  (default: 0.01)
   --spread value, -s value  Aurora spread (default: 3)
   --seed value              Set seed (default: 0)
```

# Usage

`golor Hello World`

`golor rainbow Hello World`

`echo "Hello World" | golor`

```
echo "Hello World" > text
golor -f text
```

# Special Thanks

* [busyloop/lolcat: Rainbows and unicorns!](https://github.com/busyloop/lolcat)
* [Code-Hex/Neo-cowsay: 🐮 cowsay written in Go](https://github.com/Code-Hex/Neo-cowsay)
