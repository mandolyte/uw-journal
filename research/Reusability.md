# Links

The Epic is here:
https://github.com/unfoldingWord/gateway-edit/issues/423

This is the "checklist" issue in the epic:
https://github.com/unfoldingWord/gateway-edit/issues/430

# Checklist for Ease of Reuse

This will be a checklist for *easily* reusing a component library. A *library* may expose multiple, related components.

## Scope

- Reusability will be limited to React applications and React-based frameworks, such as Next.js.

## Prerequisites

1. NPM will be used as the reuse-infrastructure.
2. Semver will be used to signal the kinds of changes over time made to the component,
3. The language will be JavaScript or TypeScript.
4. Dependencies will kept to a minimum to minimize impact on size and performance of apps using the component

## Outstanding Questions

1. Use of Material-UI: does it meet prerequisite 4?
2. Use of Tailwind: does it meet prerequisite 4?
3. Type safety is widely considered a critical aspect of ease of reuse. Plus IDE's will enforce type safety so that mistakes are prevented at edit time. This would suggest that we use TypeScript as the preferred language. This deserves some discussion!
4. Documentation options:
	- Styleguidist (or other doc tool)
	- Example folder with instructions on how to run locally
	- Examples built with code sandbox or stackblitz
	- Some combination of the above

## Checklist

- [ ] Component built using functions, not classes.
- [ ] Clear and documented API (with type definitions as needed for TypeScript).
- [ ] Complete examples showing use of the entire API (see OQ#4).
- [ ] Minimal use of external dependencies.
- [ ] For required external dependencies (such as React, MUI, etc.):
	- [ ] Tested versions (and combination of versions if needed)
	- [ ] Use of "peer" dependencies to let app chose versions (ie, avoid mixing different versions of library in the same app)
- [ ] Do not rely on "tree-shaking" to prune unused code in the library. Instead, make each library component directly accessible.