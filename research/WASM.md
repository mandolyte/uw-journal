# Objective

Write a component in Go that:
- converted to WASM
- published to NPM
- imported into a Javascript application

# Links

This one seems to focus on what I am interested in
https://github.com/sanathkr/go-npm

# Example 1
from: https://blog.suborbital.dev/foundations-wasm-in-golang-is-fantastic

I completed  much of the tutorial. The examples are focused on application level use of WASM.
But instructive nonetheless. So I need to keep looking...

# first-component

Using notes in appendix C.

First, wasm-pack init is deprecated... probably can skip this step. Here is the output:
```text
$ wasm-pack init first-component
wasm-pack init is deprecated, consider using wasm-pack build
Error: crate directory is missing a `Cargo.toml` file; is `first-component` the wrong directory?
Caused by: crate directory is missing a `Cargo.toml` file; is `first-component` the wrong directory?
```
Makes me wonder if this is a Rust only tool?? It did not actually create a folder. So I created it myself.

Next I copied main.go from the second-function folder.

Since this is pure code, not intended access browser features, this will be a plain old Go program, which takes a string and returns an array. So modifying the main.go as needed.

So far, the above is not working. Appears to assume that Rust is being used.


https://developer.mozilla.org/en-US/docs/WebAssembly/Rust_to_Wasm
This link has a detailed use of Rust to make an NPM component. It may be possible to insert Go generated WASM code into the steps outlined. I went thru this tutorial without any issues.

Now this tutorial using Go: https://golangbot.com/webassembly-using-go/
**Results:** The wasm code is wrapped in order to be used by a Javascript application. So perhaps the route forward to making an NPM package is to write a wrapper (perhaps a class like Sqlite3?) and then use NPM to deploy that wrapper. 

#todo study the Rust example and compare with Go example and see if I can apply the approach to Go that Rust is using.




# Appendix A
Here is a prompt:
> can i deploy a wasm component to npm

Here is Bard's response:
Yes, it is possible to deploy a WebAssembly (wasm) component to npm. To do this, you will need to create a package.json file that specifies the name of your component, the version number, and the dependencies that your component requires. You will also need to create a build script that compiles your wasm component into a format that can be deployed to npm. Once you have created these files, you can use the npm publish command to deploy your component to the npm registry.

Here are the steps on how to deploy a wasm component to npm:
1. Create a directory for your component.
2. Create a package.json file in the directory.
3. In the package.json file, specify the name of your component, the version number, and the dependencies that your component requires.
4. Create a build script that compiles your wasm component into a format that can be deployed to npm.
5. Run the npm publish command to deploy your component to the npm registry.

***wasm-bindgen does not exist!*
Here is an example of a package.json file for a wasm component:
```json
{
  "name": "my-wasm-component",
  "version": "1.0.0",
  "dependencies": {
    "wasm-bindgen": "^0.2.79"
  }
} 
```
Here is an example of a build script that compiles a wasm component:
```sh
#!/bin/sh

# Build the wasm component.
wasm-bindgen -o my-wasm-component.wasm my-wasm-component.rs

# Deploy the wasm component to npm.
npm publish my-wasm-component.wasm 
```

Once you have completed these steps, your wasm component will be deployed to the npm registry and can be used by other developers.

# Appendix B - using go to create an NPM package

Prompt: go lang to wasm publish to npm

Response:
Here are the steps on how to go from Golang to WASM and publish it to npm:

- Install the Golang compiler and the wasm-pack tool.
- Create a Golang project.
- Write your Golang code.
- Compile your Golang code to WASM.


- Create an npm package.
- Publish your npm package.

Here are the details for each step:
To install the Golang compiler, run the following command:
> go install golang.org/dl/go 

To install the wasm-pack tool, run the following command:
> npm install -g wasm-pack 

To create a Golang project, create a directory for your project and initialize it as a Golang project:
- mkdir my-project
- cd my-project
- go init 

To write your Golang code, create a file called main.go and write your code in it. For example, the following code prints "Hello, world!" to the console:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
} 
```

To compile your Golang code to WASM, run the following command:
> wasm-pack build 

This will create a directory called pkg in your project directory. The pkg directory will contain the compiled WASM file.

To create an npm package, run the following command:
> npm init 

This will create a file called package.json in your project directory. The package.json file is used to specify the metadata for your npm package.

To publish your npm package, run the following command:
> npm publish 

This will publish your npm package to the npm registry. Other developers can then install your npm package by running the following command:
> npm install my-package 

Once your npm package is installed, other developers can use it in their projects. For example, the following code imports the my-package npm package and prints "Hello, world!" to the console:

```js
import "my-package"

func main() {
  mypackage.PrintHelloWorld()
} 
```

# Appendix  C - publish wasm to npm

Prompt: publish wasm to npm

Sure. Here are the steps on how to publish a WebAssembly (wasm) to npm:
- Install the `wasm-pack` tool.
- Initialize a wasm-pack project.
- Write your wasm code.
- Compile your wasm code.
- Create an npm package.
- Publish your npm package.

Here are the detailed steps:

Install the `wasm-pack` tool:
> npm install -g wasm-pack 

Initialize a wasm-pack project:
> wasm-pack init my-wasm-project 

Write your wasm code.

You can write your wasm code in any language that supports WebAssembly, such as Rust, C++, or AssemblyScript.

Compile your wasm code:
> wasm-pack build 


This will compile your wasm code and generate a `pkg` directory.

Create an npm package:
> npm init 

This will create a `package.json` file.

Publish your npm package:
> npm publish 

This will publish your npm package to the npm registry.
Once you have published your npm package, other JavaScript developers can install it using the `npm install` command.