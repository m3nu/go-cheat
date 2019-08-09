# go-cheat
`go-cheat` allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

`go-cheat` is a Go implementation of [Chris Allen Lane](https://github.com/chrisallenlane)'s great [`cheat`](https://github.com/cheat/cheat) Python package. It aims to be compatible with the Python version and solves those issues:

- Many users don't maintain a Python installation and `pip` is often not available. `go-cheat` comes as a single 
  binary for most architectures and OS. See [releases]().
- `cheat` requires root permissions to install system-wide cheat sheets. `go-cheat` includes a 
  [default set](https://github.com/cheat/cheat/tree/master/cheat/cheatsheets) of cheat sheets from `cheat` in the binary.

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
Download the binary release for your platform from the Release page.

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

## See Also:
- [Enabling Command-line Autocompletion][autocompletion]
- [Related Projects][related-projects]


[autocompletion]:   https://github.com/cheat/cheat/wiki/Enabling-Command-line-Autocompletion
[dotfiles]:         http://dotfiles.github.io/
[gfm]:              https://help.github.com/articles/creating-and-highlighting-code-blocks/
[installing]:       https://github.com/cheat/cheat/wiki/Installing
[pygments]:         http://pygments.org/
[related-projects]: https://github.com/cheat/cheat/wiki/Related-Projects
