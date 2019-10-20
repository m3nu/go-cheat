# go-cheat (depreceated)
**The original `cheat` was rewritten in Go and has many more features than this projects. Recommend to use that. See also [this issue](https://github.com/cheat/cheat/issues/470) or the [main project page](https://github.com/cheat/cheat)**


`go-cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

`go-cheat` is a Go implementation of [Chris Allen Lane](https://github.com/chrisallenlane)'s great [`cheat`](https://github.com/cheat/cheat) Python package. It aims to be fully compatible with the Python version and solves those issues:

- Many users don't maintain a Python installation and `pip` is often not available. `go-cheat` comes as a single binary for most architectures. See [releases](https://github.com/m3nu/go-cheat/releases).
- `cheat` tries to install system-wide cheat sheets and [fails(https://github.com/cheat/cheat/issues/431)] without root permissions. `go-cheat` includes a [default set](https://github.com/cheat/cheat/tree/master/cheat/cheatsheets) of cheat sheets from `cheat` in the binary.

## Todo:
- [ ] Tests
- [ ] Nicer release script and more architectures

## Example
The next time you're forced to disarm a nuclear weapon without consulting
Google, you may run:

```sh
cheat tar
```

You will be presented with a cheatsheet resembling:

```sh
# To extract an uncompressed archive: 
tar -xvf '/path/to/foo.tar'

# To extract a .gz archive:
tar -xzvf '/path/to/foo.tgz'

# To create a .gz archive:
tar -czvf '/path/to/foo.tgz' '/path/to/foo/'

# To extract a .bz2 archive:
tar -xjvf '/path/to/foo.tgz'

# To create a .bz2 archive:
tar -cjvf '/path/to/foo.tgz' '/path/to/foo/'
```

To see what cheatsheets are available, run `cheat -l`.

Note that, while `cheat` was designed primarily for \*nix system administrators,
it is agnostic as to what content it stores. If you would like to use `cheat`
to store notes on your favorite cookie recipes, feel free.


## Installing
Download the binary release or package (`rpm`/`deb`) for your platform from [Releases](https://github.com/m3nu/go-cheat/releases).

To install the `.deb`:
`$ apt install ./cheat-*.deb`

Or the `.rpm`:
`$ yum localinstall ./cheat-*.rpm`

## Modifying Cheatsheets
The value of `cheat` is that it allows you to create your own cheatsheets - the
defaults are meant to serve only as a starting point, and can and should be
modified.

Cheatsheets are stored in the `~/.cheat/` directory, and are named on a
per-keyphrase basis. In other words, the content for the `tar` cheatsheet lives
in the `~/.cheat/tar` file.


## Configuring

### Setting a CHEAT_USER_DIR ###
Personal cheatsheets are saved in the `~/.cheat` directory by default, but you
can specify a different default by exporting a `CHEAT_USER_DIR` environment
variable:

```sh
export CHEAT_USER_DIR='/path/to/my/cheats'
```

### Setting a CHEAT_PATH ###
You can additionally instruct `cheat` to look for cheatsheets in other
directories by exporting a `CHEAT_PATH` environment variable:

```sh
export CHEAT_PATH='/path/to/my/cheats'
```

You may, of course, append multiple directories to your `CHEAT_PATH`:

```sh
export CHEAT_PATH="$CHEAT_PATH:/path/to/more/cheats"
```

You may view which directories are on your `CHEAT_PATH` with `cheat -d`.


### Enabling Syntax Highlighting ###
`cheat` can optionally apply syntax highlighting to your cheatsheets. To
enable syntax highlighting, export a `CHEAT_COLORS` environment variable:

```sh
export CHEAT_COLORS=true
```

`cheat` ships with both light and dark colorschemes to support terminals with
different background colors. A colorscheme may be selected via the
`CHEAT_COLORSCHEME` envvar. Valid values can be found [here](https://xyproto.github.io/splash/docs/):

```sh
export CHEAT_COLORSCHEME=pygments
```


## See Also:
- [Enabling Command-line Autocompletion][autocompletion]
- [Related Projects][related-projects]


[autocompletion]:   https://github.com/cheat/cheat/wiki/Enabling-Command-line-Autocompletion
[dotfiles]:         http://dotfiles.github.io/
[gfm]:              https://help.github.com/articles/creating-and-highlighting-code-blocks/
[installing]:       https://github.com/cheat/cheat/wiki/Installing
[pygments]:         http://pygments.org/
[related-projects]: https://github.com/cheat/cheat/wiki/Related-Projects
