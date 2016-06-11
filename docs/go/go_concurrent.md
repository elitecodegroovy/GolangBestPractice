## Mutex

## 1 sync.Mutex
There is a good reason Go’s mutexes are not re-entrant. The pur pos e of a mutex is to ensure
that certain invariants of the shared variables are maintained at critical points dur ing program
execution. One of the invar iants is "no goroutine is accessing the shared variables," but there
may be additional invariants specific to the dat a structures that the mutex guards.When a
goroutine acquires a mutex lock, it may assume that the invariants hold. While it holds the
lock, it may update the shared variables so that the invariants are temporarily violated.
However, when it releases the lock, it must guarantee that order has been restored and the
invariants hold once again. Although a re-entrant mutex would ensure that no other
goroutines are accessing the shared variables, it cannot protect the additional invariants of
those variable.

We can use the following methods to sync transaction.
```
mu.Lock()
defer mu.Unlock()
```

e.g.



### 2 Read/WriteMutexes:sync.RWMutex
Sometimes we need a speci al kind of lock that allows read-only operat ions to
proceed in paral lel with each other, but write operations to have fully exclusive access. This
lock is called a multiple readers, single writer lock, and in Go it’s provide d by sync.RWMutex:

```
var mu sync.RWMutex
var balance int
func Balance() int {
    mu.RLock() // readers lock
    defer mu.RUnlock()
    return balance
}
```

It’s only profitable to use an RWMutex when most of the goroutines that acquire the lock are
readers, and the lock is under contention, that is, goroutines routinely have to wait to acquire
it. An RWMutex requires more complex internal bookkeeping, making it slower than a regular
mutex for uncontended locks.

### 3 Memory Synchronization
There are two reasons we need a
mutex. The first is that it’s equally important that Balance not execute in the middle of some
other operation like Withdraw. The second (and more subtle) reason is that synchronization
is about more than just the order of execution of multiple goroutines; synchronization also
affects memory.

### 4 Lazy Initialization: sync.Once
Using sync.Once in this way, we can avoid sharing variables with other goroutines until they have
 been properly constructed.
 ```
 var loadIconsOnce sync.Once
 ....
 loadIconsOnce.Do(InitMethod)
 ```

 ### 5 The Race Detector
Even with the greatest of care, it’s all too easy to make concurrency mistakes. Fortunately, the
Go runtime and toolchain are equipped with a sophisticated and easy-to-use dynamic analysis
tool, the race detector.

### 6 Goroutines and Threads
Although the dif ferences bet ween them are essential
ly quantitative, a big enough quantitative dif ference becomes a qualitative one, and so it is
with goroutines and threads. The time has now come to distinguish them.

#### 6.1 Growable Stacks
Each OS thread has a fixed-size blo ck of memory (often as large as 2MB) for its stack, the work
area where it saves the local variables of function cal ls that are in progress or temporarily
suspended while another function is called.It’s not uncommon for a Go program to create hundreds
of thousands of goroutines at one time, which would be impossible with stacks this large.

Changing the fixed size can improvespace efficiency and allow more threads to be created, or it
can enable more deeply recursive functions, but it cannot do both.

In contrast, a goroutine starts life with a small stack, typic ally 2KB.A goroutine’s stack is
not fixed, it grows and shr inks as needed.The size limit for a goroutine stack may be as much
as 1GB, orders of magnitude larger than a typical fixed-size thread stack, though of course
few goroutines use that much.

#### 6.2 Goroutine Scheduling
**OS threads** are scheduled by the OS ker nel. Every few milliseconds, a hardware timer interrupts
the processor, which causes a ker nel function cal le d the scheduler to be invoked. This
function suspends the cur recently executing thread and saves its registers in memory, looks over
the list of threads and decides which one should run next, restores that thread’s registers from
memory, then resumes the execution of that thread. Because OS threads are scheduled by the
kernel, passing control from one thread to another requires a full context switch, that is, saving
the state of one user thread to memory, restoring the state of another, and updating the
scheduler’s dat a structures. This operation is slow, due to its poor locality and the number of
memory accesses required, and has historically only gotten worse as the number of CPU cycles
required to access memory has increased.

The Go runtime contains its own scheduler that uses a technique known as m:n scheduling,
because it multiplexes (or schedules) m goroutines on n OS threads. The job of the Go
scheduler is analogous to that of the ker nel scheduler, but it is concerned only with the
goroutines of a single Go program.

Unlike the operating system’s thread scheduler, the Go scheduler is not invoked periodic ally
by a hardware timer, but implicitly by certain Go language constructs. For example, when a
goroutine cal ls time.Sleep or blo cks in a channel or mutex operation, the scheduler puts it to
sleep and runs another goroutine until it is time to wake the first one up. Because it doesn’t
need a switch to kernel context, rescheduling a goroutine is much cheaper than rescheduling a
thread.

#### 6.3 GOMAXPROCS
The Go scheduler uses a parameter cal le d GOMAXPROCS to deter mine how many OS threads
may be actively executing Go code simultaneously. Its default value is the number of CPUs on
the machine, so on a machine with 8 CPUs, the scheduler will schedule Go code on up to 8 OS
threads at once. (GOMAXPROCS is the n in m:n scheduling.)

You can explicitly control this parameter using the GOMAXPROCS environment variable or the
runtime.GOMAXPROCS function.

```
for {
go fmt.Print(0)
fmt.Print(1)
}
$ GOMAXPROCS=1 go run hackerclich.go
go
111111111111111111110000000000000000000011111...
$ GOMAXPROCS=2 go run hackerclich.go
go
010101010101010101011001100101011010010100110...
```
In the first run, at most one goroutine was executed at a time. Initial ly, it was the main
goroutine, which prints ones. After a period of time, the Go scheduler put it to sleep and woke
up the goroutine that prints zeros, giving it a tur n to run on the OS thread. In the second run,
there were two OS threads available, so both goroutines ran simultaneously, printing digits at
about the same rate.

#### 6.4 Goroutines Have No Identity
In most operating systems and programming languages that support multithreading, the current
thread has a distinct identity that can be easily obtained as an ordinar y value, typic ally an
integer or pointer. This makes it easy to build an abstraction called thread-local storage, which
is essentially a global map keyed by thread identity, so that each thread can store and retrieve
values independent of other threads.

Goroutines have no notion of identity that is accessible to the programmer.
