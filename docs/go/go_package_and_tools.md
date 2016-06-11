## Package and the go tool
Amodest-size program today might contain 10,000 functions. Yet its author need think about
only a few of them and design even fewer, because the vast majority were written by others
and made available for reuse through packages.

Go comes with over 100 standard packages that provide the foundations for most applications.

Go also comes with the go tool, a sophisticated but simple-to-use command for managing
workspaces of Go packages.

### 1 Introduction
The pur pos e of any package system is to make the design and maintenance of large programs
practical by grouping related features to get her into units that can be easily understood and
changed, independent of the other packages of the program. This modularity al lows packages
to be shared and reused by different projects, distributed within an organization, or made
available to the wider world.

Packages also provide encapsulation by controlling which names are visible or exported
outside the package. Restricting the visibility of package members hides the helper functions
and types behind the package’s API, allowing the package maintainer to change the implementation
with confidence that no code outside the package will be affected. Restricting visibility
also hides variables so that clients can access and update them only through exported functions
that preserve internal invariants or enforce mutual exclusion in a concurrent program.

### 2 Import Paths
Each package is identified by a unique string called its import path. Import paths are the
strings that appear in import declarations.
```
import (
    "fmt"
    "math/rand"
    "encoding/json"
    "golang.org/x/net/html"
)
```

### 3 The Package Declaration
A package declaration is required at the start of every Go source file. Its main purpose is to
deter mine the default identifier for that package (called the package name) when it is imported
by another package.

```
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println(rand.Int())
}
```
There are three maj or exceptions to the "last segment" convention. The first is that a package
defining a command (an exec utable Go program) always has the name main, regardless of the
package’s import path. This is a signal to **go build** that it must invoke the linker to
make an executable file.

The second exception is that some files in the directory may have the suffix _test on their
package name if the file name ends with _test.go. Such a directory may define two packages:
the usual one, plus another one called an external test package. The _test suffix signals to
go test that it must build both packages, and it indicates which files belong to each package.
External test packages are used to avoid cycles in the import graph arising from dependencies
of the test.

The third exception is that some tools for dependency management **append** version number
suffixes to package import paths, such as "gopkg.in/yaml.v2". The package name excludes
the suffix, so in this case it would be just **yaml**.

### 4 Import Declarations
AGo source file may contain zero or more import de clarations immediately after the package
declaration and before the first non-import declaration. Each import declaration may specify
the import path of a single package, or multiple packages in a parenthesized list. The two
forms below are equivalent but the second form is more common.
```
import "fmt"
import "os"

```
The more common is like following style.
```
import (
    "fmt"
    "os"
)
```
Imported packages may be grouped by introducing blank lines; such groupings usually indicate
different domains. The order is not significant, but by convention the lines of each group
are sorted alphabetically.

If we need to import two packages whose names are the same, like **math/rand** and
**crypto/rand**, into a third package, the import declarat ion must specif y an alternative name
for at least one of them to avoid a conflict.
```
import (
    "crypto/rand"
    mrand "math/rand" // alternative name mrand avoids conflict
)
```

### 5 Blank Imports
It is an error to import a package into a file but not refer to the name it defines within that file.
However, on occasion we must import a package merely for the side effects of doing so: evaluation
of the initializer expressions of its package-level variables and execution of its init functions.

To suppress the  "unused import" er ror we would other wise encounter, we must use a renaming
import in which the alternative name is _, the blank identifier.
```
import _ "image/png" // register PNG decoder
```
This is known as a blank import. It is most often used to implement a compile-time
mechanism whereby the main program can enable optional features by blank-importing additional
packages. First we’ll see how to use it, then we’ll see how it works.

```
// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main
import (
    "fmt"
    "image"
    "image/jpeg"
    _ "image/png" // register PNG decoder
    "io"
    "os"
)
func main() {
    if err := toJPEG(os.Stdin, os.Stdout); err != nil {
    fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
    os.Exit(1)
    }
}

func toJPEG(in io.Reader, out io.Writer) error {
    img, kind, err := image.Decode(in)
    if err != nil {
        return err
    }
    fmt.Fprintln(os.Stderr, "Input format =", kind)
    return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
```

It detects the PNG input format and writes a JPEG version.

### 6 Packages and Naming
In this section, we’ll offer some advice on how to follow Go’s distinctive conventions for naming
packages and their members.

When creating a package, keep its name short, but not so short as to be cryptic. The most
frequently used packages in the standard library are named **bufio, bytes, flag, fmt,
http, io, json, os, sort, sync, and time**.



