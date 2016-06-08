## Mutex

##sync.Mutex
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



###Read/WriteMutexes:sync.RWMutex
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

### Memory Synchronization
There are two reasons we need a
mutex. The first is that it’s equally important that Balance not execute in the middle of some
other operation like Withdraw. The second (and more subtle) reason is that synchronization
is about more than just the order of execution of multiple goroutines; synchronization also
affects memory.