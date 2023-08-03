# Objective

Write a component in Go that:
- converted to WASM
- published to NPM
- imported into a Javascript application

# To Do

This seems to be fairly complete... cross my fingers:
https://dev.to/royhadad/how-to-create-a-react-app-with-go-support-using-webassembly-in-under-60-seconds-4oa3
- note that he uses a tweaked version of `wasm_exec.js` -- I need to study that since the one for tinygo gets warnings about the requires. Namely:

```
./pages/wasm_exec.js
Critical dependency: require function is used in a way in which dependencies cannot be statically extracted
```


This uses rollup to package the wasm:
https://stackoverflow.com/questions/75989746/how-to-export-wasm-from-an-npm-package-in-a-way-that-react-js-would-be-able-to-u

This has some ideas that might work:
https://codeburst.io/getting-started-with-react-and-webassembly-using-hooks-441818c91608

This might work using React with some other of the ideas above:
https://github.com/mbasso/react-wasm

# Links

This might be useful to learn from:
https://davetayls.me/blog/2022-11-24-use-wasm-compiled-golang-functions-in-nextjs

Notice this proposal!
https://github.com/WebAssembly/esm-integration/tree/main/proposals/esm-integration
(referenced in https://dev.to/sendilkumarn/loading-wasm-as-esm-in-nodejs-2gdb)

This one seems to focus on what I am interested in
https://github.com/sanathkr/go-npm

This one has a section about publishing to NPM:
https://david.coffee/react-native-wasm-golang/

# Tutorial
from: https://blog.suborbital.dev/foundations-wasm-in-golang-is-fantastic

I completed  much of the tutorial. The examples are focused on application level use of WASM.
But instructive nonetheless. So I need to keep looking...

# Making a Component

A lot of tweaks were made - see repo `experiment-wasm` in the `pretty-print-go-wasm` folder.
- added babel config
- patched the `prepublishOnly` script to force how I wanted the "dist" folder to look.

I am getting this error:
```
Error: globalThis.crypto is not available, polyfill required (crypto.getRandomValues only)
```

After some searches, trying the solution in second link above, namely:

```
import getRandomValues from "polyfill-crypto.getrandomvalues";
if (!globalThis.crypto) {
  globalThis.crypto = {
    getRandomValues,
  };
}
```
This is a modification to Go's wasm file.

*Got past that...*

Next error: 
```
error - unhandledRejection: TypeError: _wasm_exec.default is not a constructor
```

Advice in stack overflow at https://stackoverflow.com/questions/42652423/error-default-is-not-a-constructor is to do this:

```
export default class Go
```



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