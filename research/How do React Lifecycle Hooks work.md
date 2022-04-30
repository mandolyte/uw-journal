# How do React Lifecycle Hooks Work?

For the purpose of this discussion, let's initially, at least, focus on `useEffect()`. 

A useEffect has several parts:
- a body
- a dependency list

Documentation for hooks is at: 
https://reactjs.org/docs/hooks-intro.html

## Array dependency
Dependencies must be primitive values for React.useEffect. A reference (or "pointer") is a primitive value.

This means that arrays and objects, which are references will not cause the hook to run when updated, since the array or object will be at the same memory location after an update.

In the array example below, if the line:
`let newlist = alist;` is used, then the code will not work, even tho the array is updated every time the button is clicked.

On the other hand, if the line:
`let newlist = alist.map( n => n);` is used, then a new array is 
allocated in memory with the same elements. Then when it is used by
`setAlist()` it will have new reference (pointer) and then the useEffect will run.

```js
import "./styles.css";
import React, { useState, useEffect } from "react";
export default function App() {
  const [count, setCount] = useState(0);
  const [alist, setAlist] = useState([]);

  const handleClick = () => {
    // The commented line does not work, since
    // the value of newlist is a pointer to 
    // the same place as alist. I.e., unchanged
    // whereas using the trivial "map" forces
    // a re-allocation for the array and a new 
    // pointer value.
    // let newlist = alist;
    let newlist = alist.map( n => n);
    newlist.push(count+1);
    setAlist(newlist);
  }

  useEffect(() => {
    setCount(alist.length);
  }, [alist, count]);

  return (
    <div className="App">
      <h1>Hello CodeSandbox</h1>
      <h2>Start editing to see some magic happen!</h2>

      <p>You clicked {count} times</p>
      <button onClick={ () => handleClick() } >
        Click me
      </button>
    </div>
  );
}
```


Sandbox link
https://obehi0.csb.app/


## Object dependency

Use the code below to show the same principles using objects.

```js
import "./styles.css";
import React, { useState, useEffect } from "react";
export default function App() {
  const [count, setCount] = useState(0);
  const [obj, setObj] = useState({});

  const handleClick = () => {
    // The commented line does not work, since
    // the value of newobj is a pointer to
    // the same place as obj. I.e., unchanged
    // whereas using Object.assign() will
    // make a new copy at a new location.
    // let newobj = obj;
    let newobj = Object.assign({}, obj);
    newobj[`key${count}`] = `val${count}`;
    setObj(newobj);
  };

  useEffect(() => {
    const keys = Object.keys(obj);
    setCount(keys.length);
  }, [obj, count]);

  return (
    <div className="App">
      <h1>Hello CodeSandbox</h1>
      <h2>Start editing to see some magic happen!</h2>

      <p>You clicked {count} times</p>
      <button onClick={() => handleClick()}>Click me</button>
    </div>
  );
}
```

## References

Per React docs, comparisons are done via the JavaScript Object.is() function. See [here](https://reactjs.org/docs/hooks-reference.html#bailing-out-of-a-dispatch)

And [here](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/is#description) is documentation on Object.is().
