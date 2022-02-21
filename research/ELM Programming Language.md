# ELM Programming Language

# Links

- Easy intro: https://elmprogramming.com/
- Installation on Linux: https://github.com/elm/compiler/blob/master/installers/linux/README.md
- VS Code syntax etc: https://github.com/elm-tooling/elm-language-client-vscode
- Online playground: https://elm-lang.org/try
- Slack: https://elmlang.slack.com
- Package Manager: https://package.elm-lang.org/
- Youtube playlist: https://www.youtube.com/playlist?list=PL-cYi7I913S-VgTSUKWhrUkReM_vMNQxG
- 

# Objectives
1. What is elm like as a language? what can we learn from it?
2. How well supported is it? Is it risky?
3. Is it adequate to replace some/all of current tech stack?
4. Community support (see slack)

# Diary

## 2022-02-21

Note: People coming from languages like JavaScript may be surprised that functions look different here:
```elm
madlib "cat" "ergonomic"                  -- Elm
madlib("cat", "ergonomic")                // JavaScript

madlib ("butter" ++ "fly") "metallic"      -- Elm
madlib("butter" + "fly", "metallic")       // JavaScript
```
A **function** using the elm repl tool:
```
> greet name =
|   "Hello " ++ name ++ "!"
| 
<function>

> greet "Bob"
"Hello Bob!"
```

**Lists**:
Lists can hold many values. Those values must all have the same type. Here are a few examples that use functions from the List module (https://package.elm-lang.org/packages/elm/core/latest/List)

**Tuples** are another useful data structure. A tuple can hold two or three values, and each value can have any type. A common use is if you need to return more than one value from a function. The following function gets a name and gives a message for the user:
```elm
> isGoodName name =
|   if String.length name <= 20 then
|     (True, "name accepted!")
|   else
|     (False, "name was too long; please limit it to 20 characters")
| 
<function>
```

A record (sort of like JSON) can hold many values, and each value is associated with a name.
Here is a record that represents British economist John A. Hobson:

> john =
|   { first = "John"
|   , last = "Hobson"
|   , age = 81
|   }
| 
{ age = 81, first = "John", last = "Hobson" }

> john.last
"Hobson"


It is often useful to update values in a record (note immutability, "john" is not actually updated; instead a new value is produced):
```
> john = { first = "John", last = "Hobson", age = 81 }
{ age = 81, first = "John", last = "Hobson" }

> { john | last = "Adams" }
{ age = 81, first = "John", last = "Adams" }

> { john | age = 22 }
{ age = 22, first = "John", last = "Hobson" }
```
Notice that when we update some fields of john we create a whole new record. It does not overwrite the existing one. Elm makes this efficient by sharing as much content as possible. If you update one of ten fields, the new record will share the nine unchanged values.

So a function to update ages might look like this:
```elm
> celebrateBirthday person =
|   { person | age = person.age + 1 }
| 
<function>

> john = { first = "John", last = "Hobson", age = 81 }
{ age = 81, first = "John", last = "Hobson" }

> celebrateBirthday john
{ age = 82, first = "John", last = "Hobson" }


>  
```


### Architecture
The basic pattern of an elm program is described here: 
https://guide.elm-lang.org/architecture/

What happens within the Elm program though? It always breaks into three parts:

- Model — the state of your application
- View — a way to turn your state into HTML
- Update — a way to update your state based on messages

These three concepts are the core of The Elm Architecture.



## 2022-02-18

1. Created repo: https://github.com/mandolyte/elm-programming-language
2. Linux install script tweaked and in the bin folder of preceding.
3. Version of Linux compiler:
```sh
$ elm --version
0.19.1
```
4. Installation page at: https://guide.elm-lang.org/install/elm.html. This has a number of examples of commands I should know.

