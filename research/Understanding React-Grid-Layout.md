# Understanding React-Grid-Layout
## Links
- Link to the demos: https://strml.github.io/react-grid-layout/examples/1-basic.html
- Link to the "showcase": https://strml.github.io/react-grid-layout/examples/0-showcase.html (this has few restrictions of moving and sizing)


## breakpoint

A breakpoint specifies where to re-layout the children components into columns based on screen sizes.

Example:
```js
const breakpoints = { lg: 900, md: 700, sm: 500 };
```

This set of breakpoints uses 900, 700, and 500. The numbers are relative and the above could be re-written as:

```js
const breakpoints = { lg: 9, md: 7, sm: 5};
```

For each size of display, this specifies how many columns should be used to layout the children components. Thus:
- `lg` will allow 9 columns
- `sm` will allow only 5

## layout

Since the placement of the children components is constrained by the size of the screen, then a layout is associated with a "screen size", which must match one of the breakpoint property names.

Example:
```js
const layout = {
  lg: [
    {"w":6,"h":1,"x":0,"y":0,"i":"1"},
    {"w":6,"h":1,"x":6,"y":2,"i":"2"},
    {"w":6,"h":1,"x":0,"y":2,"i":"3"},
    {"w":6,"h":1,"x":6,"y":0,"i":"4",static:true},
    {"w":12,"h":1,"x":0,"y":1,"i":"5"},
  ]
};
```
This layout only provides placements for a large ("lg") screen size.

A *layout* is an object of breakpoint names, which have an array of objects, one per child component, with positioning information, namely (using the breakpoints above):

- "w": the width in number of columns to take up (0 to 8 for the "lg" example above)
- "h": the height in number of rows, which is inferred from the max "y" coordinate in the layout.
- - "x": the horizontal coordinate. Zero based. Starting from left side of the container.
	- A number from 0 to 8 for the "large" (lg) 
- "y": the vertical coordinate. Zero based. Starting from the top of the container.
- "i": the unique number for the component. In the above, it goes from "1" to "5" (strings, not numbers).
- "static": this is optional and defaults to false. If set to true, then that child component cannot be moved or resized.

## layoutWidths and layoutHeights
The documentation says these are arrays of arrays of numbers.
Example:
```js
const layoutWidths = [[1, 1], [1, 1], [1]];
```

Perhaps these arrays determine how many children can fit into each row?

No idea how LayoutHeights might be used...

## minH and minW

The component properties says these are numbers, but no other info is provided.

## rows

The documentation says this is a number and is *The number of lines in the grid that fit on 1 screen.* The default is 12.

## totalGridUnits

A number defaulting to 12. But do not know its purpose.

## rowHeight

A number defaulting to 100. But no info is provided as to its purpose or its unit of measure.

## gridMargin

An array of 2 numbers, defaulting to `[0,0]`. Unit of measure is unknown. Controls the margin on left and right with first number and top and bottom with second number.

## corrrectHeight

A number defaulting to zero. The documentation states: *If there is an appbar or footer, then you can specify how much to reduce the screen size.*

## columns

A "shape" (whatever that is) and the documentation states:
```
lg: number
md: number
sm: number
xs: number
xxs: number
```

How is this different than a set of breakpoints? (next)

## breakpoints

Is also a "shape" and is described exactly like the `columns` property.

## autoResize

A Boolean, defaulting to false, and is described as:
> Whether it is necessary to automatically recalculate the rowHeight of the cards when the screen is resized so that the cards maintain their proportions relative to the screen.