## Go reflection
Go provides a mechanism to update variables and inspect their values at run time, to call their
methods, and to apply the operations intrinsic to their representation, all without knowing
their types at compile time. This mechanism is called reflection. Reflection also lets us treat
types themselves as first-class values.

###1 Why Reflection?
Sometimes we need to write a function cap able of dealing uniformly with values of typ es that
don’t satisfy a common interface, don’t have a known representation, or don’t exist at the time
we design the function—or even all three.

We start with a typ e switch that tests whether the argument defines a String method, and call
it if so. We then add switch cases that test the value’s dynamic type against each of the basic
types—string, int, bool, and so on—and perform the appropriate formatting operation in
each case.

    func Sprint(x interface{}) string {
        type stringer interface {
            String() string
        }
        switch x := x.(type) {
            case stringer:
            return x.String()
            case string:
            return x
            case int:
            return strconv.Itoa(x)
            // ...similar cases for int16, uint32, and so on...
        case bool:
        if x {
            return "true"
        }
        return "false"
        default:
        // array, chan, func, map, pointer, slice, struct
        return "???"
        }
    }

Without a way to inspect the representation of values of unknown types, we quickly get stuck.
What we need is reflection.

### 2 reflect.Type and reflect.Value
Reflection is provide d by the reflect package. It defines two important typ es, Type and
Value. A Type represents a Go typ e. It is an interface with many met hods for discriminating
among types and inspecting their components, like the fields of a struct or the parameters of a
function. The sole implementation of reflect.Type is the type descriptor , the same
entity that identifies the dynamic type of an interface value.

The reflect.TypeOf function accepts any interface{} and returns its dynamic type as a
reflect.Type:

    t := reflect.TypeOf(3) // a reflect.Type
    fmt.Println(t.String()) // "int"
    fmt.Println(t) // "int"
