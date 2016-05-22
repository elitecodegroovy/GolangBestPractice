
对于未来的编程形式，会不会像太极（**tai-chi**）一样充分发挥自己的设立能力和想象能力呢？
![](https://d1ohg4ss876yi2.cloudfront.net/blog/programming-productivity-depends-on-your-tools-and-language-choice/future-of-programming.jpg)
[YouTube – Iron Man 2 – SFX montage by Prologue Films](https://www.youtube.com/watch?v=VB3w5NhCicU). <br>

It's too far away to think about this. ^_^哈哈！

<br>

![weixing publiv account](http://img.blog.csdn.net/20160424104206329)

下面的文章就是对Go语言的基本介绍。

We will look at how Go does objects, polymorphism and inheritance and allow you to make your own conclusion.

Objects In Go
Go doesn’t have a something called ‘object’, but ‘object’ is only a word that connotes a meaning. It’s the meaning that matters, not the term itself.

While Go doesn’t have a type called ‘object’ it does have a type that matches the same definition of a data structure that integrates both code and behavior. In Go this is called a ‘struct’.

A ‘struct’ is a kind of type which contains named fields and methods.

Let’s use an example to illustrate this:

type rect struct {
    width int
    height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
}
There’s a lot we could talk about here. It’s probably best to walk through the code line by line and explain what is happening.

The first block is defining a new type called a ‘rect’. This is a struct type. The struct has two fields, both of which are type int.

The next block is defining a method bound to this struct. This is accomplished by defining a function and attaching (binding) it to a rect. Technically, in our example it is really attached to a pointer to a rect. While the method is bound to the type, Go requires us to have a value of that type to make the call, even if the value is the zero value for that type (in the case of a struct the zero value is nil).

The final block is our main function. The first line creates a value of type rect. There are other syntaxes we could use to do this, but this is the most idiomatic way. The second line prints to the output the result of calling the area function on our rect ‘r’.

To me this feels very much like an object. I am able to create a structured data type and then define methods that interact with that specific data.

What haven’t we done? In most object oriented languages we would be using the ‘class’ keyword to define our objects. When using inheritance it is a good practice to define interfaces for those classes. In doing so we would be defining an inheritance hierarchy tree (in the case of single inheritance).

An additional thing worth noting is that in Go any named type can have methods, not only structs. For example I can define a new type ‘Counter’ which is of type int and define methods on it. See an example at http://play.golang.org/p/LGB-2j707c

Inheritance And Polymorphism
There are a few different approaches to defining the relationships between objects. While they differ quite a bit from each other, all share a common purpose as a mechanism for code reuse.

Inheritance
Multiple Inheritance
Subtyping
Object composition
Single & Multiple Inheritance
Inheritance is when an object is based on another object, using the same implementation. Two different implementations of inheritance exist. The fundamental distinction between them is whether an object can inherit from a single object or from multiple objects. This is a seemingly small distinction, but with large implications. The hierarchy in single inheritance is a tree, while in multiple inheritance it is a lattice. Single inheritance languages include PHP, C#, Java and Ruby. Multiple inheritance languages include Perl, Python and C++.

Subtyping (Polymorphism)
In some languages subtyping and inheritance are so interwoven that this may seem redundant to the previous section if your particular perspective comes from a language where they are tightly coupled. Subtyping establishes an is-a relationship, while inheritance only reuses implementation. Subtyping defines a semantic relationship between two (or more) objects. Inheritance only defines a syntactic relationship.

Object Composition
Object composition is where one object is defined by including other objects. Rather than inheriting from them, the object contains them. Unlike the is-a relationship of subtyping, object composition defines a has-a relationship.

Inheritance In Go
Go is intentionally designed without any inheritance at all. This does not mean that objects (struct values) do not have relationships, instead the Go authors have chosen to use a alternative mechanism to connote relationships. To many encountering Go for the first time this decision may appear as if it cripples Go. In reality it is one of the nicest properties of Go and it resolves decade old issues and arguments around inheritance.

Inheritance Is Best Left Out
The following really drives home this point. It comes from a JavaWorld article titled why extends is evil :

The Gang of Four Design Patterns book discusses at length replacing implementation inheritance (extends) with interface inheritance (implements).

I once attended a Java user group meeting where James Gosling (Java’s inventor) was the featured speaker. During the memorable Q&A session, someone asked him: “If you could do Java over again, what would you change?” “I’d leave out classes,” he replied. After the laughter died down, he explained that the real problem wasn’t classes per se, but rather implementation inheritance (the extends relationship). Interface inheritance (the implements relationship) is preferable. You should avoid implementation inheritance whenever possible.

Polymorphism & Composition In Go
Instead of inheritance Go strictly follows the composition over inheritance principle. Go accomplishes this through both subtyping (is-a) and object composition (has-a) relationships between structs and interfaces.

Object Composition In Go
The mechanism Go uses to implement the principle of object composition is called embedded types. Go permits you to embed a struct within a struct giving them a has-a relationship.

An good example of this would be the relationship between a Person and an Address.

type Person struct {
   Name string
   Address Address
}

type Address struct {
   Number string
   Street string
   City   string
   State  string
   Zip    string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
    fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func main() {
    p := Person{
        Name: "Steve",
        Address: Address{
            Number: "13",
            Street: "Main",
            City:   "Gotham",
            State:  "NY",
            Zip:    "01313",
        },
    }

    p.Talk()
    p.Location()
}
Output

>  Hi, my name is Steve
>  I’m at 13 Main Gotham NY 01313
http://play.golang.org/p/LigPIVT2mf

Important things to realize from this example is that Address remains a distinct entity, while existing within the Person. In the main function we demonstrate that you can set the p.Address field to an address, or simply set the fields by accessing them via dot notation.

Pseudo Subtyping In Go
AUTHORS NOTE:
In the first version of this post it made the incorrect claim that Go supports the is-a relationship via Anonymous fields. In reality Anonymous fields appear to be an is-a relationship by exposing embedded methods and properties as if they existed on the outer struct. This falls short of being an is-a relationship for reasons now provided below. Go does have support for is-a relationships via interfaces, covered below. The current version of this post refers to Anonymous fields as a pseudo is-a relationship because it looks and behaves in some ways like subtyping, but isn’t.

The pseudo is-a relationship works in a similar and intuitive way. To extend our example above. Let’s use the following statement. A Person can talk. A Citizen is a Person therefore a citizen can Talk.

This code depends on and adds to the code in the example above.

type Citizen struct {
   Country string
   Person
}

func (c *Citizen) Nationality() {
    fmt.Println(c.Name, "is a citizen of", c.Country)
}

func main() {
    c := Citizen{}
    c.Name = "Steve"
    c.Country = "America"
    c.Talk()
    c.Nationality()
}
Output

>  Hi, my name is Steve
>  Steve is a citizen of America
http://play.golang.org/p/eCEpLkQPR3

We accomplish this pseudo is-a relationship in go using what is called an Anonymous field. In our example Person is an anonymous field of Citizen. Only the type is given, not the field name. It assumes all of the properties and methods of a Person and is free to use them or promote it’s own.

Promoting Methods Of Anonymous Fields

An example of this would be that citizens Talk just as People do, but in a different way.

For this we simply define Talk for *Citizen, and run the same main function as defined above. Now instead of calling *Person.Talk(), *Citizen.Talk() will be called instead.

func (c *Citizen) Talk() {
    fmt.Println("Hello, my name is", c.Name, "and I'm from", c.Country)
}
Output

>  Hello, my name is Steve and I'm from America
>  Steve is a citizen of America