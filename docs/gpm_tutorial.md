
gpm is a minimalist package manager for Go that leverages the power of the `go get` command and the underlying version control systems used by it to set your Go dependencies to desired versions, thus allowing easily reproducible builds in your Go projects.

Go Package Manager makes no assumptions about your dependencies and supports Git, Bazaar and Mercurial hosted Go packages, for a smoother workflow be sure to check out [gvp](https://github.com/pote/gvp) - the Go Versioning Packager which  provides dependency isolation for your projects.

<div align="center">
  <img src="https://raw.githubusercontent.com/pote/gpm/master/gpm_logo.png">
</div>
###1 一行安装gpm
```
wget https://raw.githubusercontent.com/pote/gpm/v1.4.0/bin/gpm --no-check-certificate && chmod +x gpm && sudo mv gpm /usr/local/bin
```

###2 The Godeps file
gpm expects you to have a file called Godeps in the root of your Go application in the format <import path> <tag/revision>.

Once this file is in place, running the gpm tool will download those packages and check out the specified versions.

####2.1 Packages
You can specify packages with the <import path> <version> format, where version can be a revision number (a git/bazaar/mercurial/svn revision hash) or a tag.

	$ ls .
	Godeps  foo.go  foo_test.go
	
	$ cat Godeps
	github.com/nu7hatch/gotrail               v0.0.2
	github.com/replicon/fast-archiver         v1.02
	launchpad.net/gocheck                     r2013.03.03   # Bazaar repositories are supported
	code.google.com/p/go.example/hello/...    ae081cd1d6cc  # And so are Mercurial ones
	
####Comments
The Godeps file accepts comments using a # symbol. Everything to the right of a # will be ignored by gpm, as well as empty lines.


####Extensibility

As a convention comments can be used to specify lines that gpm core should ignore but are instead intended to affect how a given gpm plugin behaves.

For example: a hypothetical gpm-track plugin that makes sure a given package is always updated to its last possible version would leverage a line like this one:

	[gpm-track] github.com/nu7hatch/gotrail
	
This convention makes the Godeps file format extensible, just as with plugins this can help identify common needs that might later on be merged into core without having to sacrifice code simplicity in order to explore new features.

#### Private Repos

Both gpm and `go get` support using private GitHub repositories! Here's what you need to do in order for a specific machine to be able to access them:

* Generate a GitHub access token by following [these instructions](https://help.github.com/articles/creating-an-access-token-for-command-line-use/).
* Add the following line to the `~/.netrc` file in your home directory.

```bash
machine github.com login <token>
```

You can now use gpm (and `go get`) to install private repositories to which your user has access! :)

#### Completeness

It is recommended to keep a healthy and exhaustive `Godeps` file in the root of all Go project that use external dependencies, remember every package that you add to the Godeps file will be installed along with its dependencies when gpm runs `go get` on it, so if you don't include these dependencies in your Godeps file you are losing the ability to reproduce a build with 100% reliability.

Make sure your Godeps file is exhaustive, this way any project includes the documentation required to be built reliably at any point in time.

### Commands

gpm has the following commands:

```bash
$ gpm             # Same as 'install'.
$ gpm get         # Parses the Godeps file, gets dependencies and sets them
                  # to the appropriate version but does not install them.
$ gpm install     # Parses the Godeps file, installs dependencies and sets
                  # them to the appropriate version.
$ gpm version     # Outputs version information
$ gpm help        # Prints this message
```

### Plugins

As of version [v1.1.1](https://github.com/pote/gpm/releases/tag/v1.1.1) gpm supports plugins, the intent of which is the ability to add powerful non-core features to gpm without compromising the simplicity of its codebase.

The way gpm plugin works is simple: whenever an unknown command is passed into gpm it will look for an executable in your `$PATH` called `gpm-<command>` and if it exists it will run it while passing all extra arguments to it, simple yet powerful.

This brings a lot to the table: plugins can be written in anything, they can be Go binaries, bash scripts, Ruby gems, Python packages, you name it. gpm wants to make it easy for you to extend it. :)

#### Installing plugins through Homebrew

I maintain a [repository with homebrew formulae for gpm plugins](https://github.com/pote/homebrew-gpm_plugins) that you can add to your system with the `brew tap` command:

```bash
$ brew tap pote/gpm_plugins
```

After you've done this you can install plugins as you would with any other homebrew packge.

```bash
$ brew install gpm-bootstrap
```

#### Known Plugins

If you have written a gpm plugin and want it included please send a pull request to the repo! I love how people have taken to explore possible features using plugins so if you've written one there is about a 99% chance I will include it here. :)

| Name and Link                   | Author                               | Short Description                 | Type        |
|:-------------------------------:|:------------------------------------:|:----------------------------------|:-----------:|
| [gpm-bootstrap][plugin-boot]    | [pote][author-pote]                  | Creates an initial Godeps file    | official    |
| [gpm-git][plugin-git]           | [technosophos][author-technosophos]  | Git management helpers            | third party |
| [gpm-link][plugin-link]         | [elcuervo][author-elcuervo]          | Dependency vendoring              | third party |
| [gpm-local][plugin-local]       | [technosophos][author-technosophos]  | Usage of local paths for packages | third party |
| [gpm-prebuild][plugin-prebuild] | [technosophos][author-technosophos]  | Improves building performance     | third party |
| [gpm-all][plugin-all]           | [pote][author-pote]                  | Installs multiple sets of deps    | official    |
| [gpm-lock][plugin-lock]         | [zeeyang][author-zeeyang]            | Lock down dependency versions     | third party |

There is no real difference on official/third party plugins other than the willingness of the gpm core team to support each, plugins labeled as third party will be supported (or not) by their authors.

[plugin-boot]: https://github.com/pote/gpm-bootstrap
[plugin-git]: https://github.com/technosophos/gpm-git
[plugin-link]: https://github.com/elcuervo/gpm-link
[plugin-local]: https://github.com/technosophos/gpm-local
[plugin-prebuild]: https://github.com/technosophos/gpm-prebuild
[plugin-all]: https://github.com/pote/gpm-all
[plugin-lock]: https://github.com/zeeyang/gpm-lock

[author-pote]: https://github.com/pote
[author-technosophos]: https://github.com/technosophos
[author-elcuervo]: https://github.com/elcuervo
[author-zeeyang]: https://github.com/zeeyang


### Further Reading

The creator for the [gpm-git](https://github.com/technosophos/gpm-git) and [gpm-local](https://github.com/technosophos/gpm-local) wrote a [fantastic blog post explaining the usage and rationale](http://technosophos.com/2014/05/29/why-gpm-is-the-right-go-package-manager.html) of gpm and [gvp](https://github.com/pote/gvp), it sums up explanations for several of the design decisions behind both tools.

### Contributing

Lots of people have contributed to make gpm what it is today, if you want to take your time to play around
with the code please do so! Opening issues on bugs, feature requests or simple food for thought are a great
way to contribute, if you send a pull request please be a good citizen and do things in a tidy manner.

* Create a feature branch with a meaningful name.
* Make sure your commit messages and PR comments are informative.
* Write a test for your feature if applicable.
* Always remember to run the test suite with `make test` before comitting.

Either way, thank you **very** much for any form of contribution, even if a patch ends up not being merged
the fact that it was sent and forced us to think about it is a contribution in itself.