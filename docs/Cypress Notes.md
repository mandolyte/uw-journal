# Cypress Notes

How to fetch a file:
```js
cy.request({yourUri}).as("response"); // This takes care of the async task. 
 
cy.get("@response").should((response) => { 
   // Carry on with the rest of your test
   cy.get("some selector");

});
```

