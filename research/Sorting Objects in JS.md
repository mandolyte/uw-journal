# Sorting Objects in JS

```js
// both work below

const array1 = [
	{n:1,x:"one"}, 
	{n:30,x:"thirty"}, 
	{n:4,x:"four"}, 
	{n:21,x:"twenty"}, 
	{n:100000,x:"big"}
];
array1.sort(
	(a,b) => {
		if ( a.n === b.n ) return 0;
		if ( a.n > b.n ) return 1;
		return -1;
	}
);

array1.sort(
	(a,b) => {
	  return a.n - b.n
	}
);
```