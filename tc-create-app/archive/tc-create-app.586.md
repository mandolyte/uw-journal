# Add a unique identifier to tC Create develop builds

Issue: [586](https://github.com/unfoldingword/tc-create-app/issues/586)
PR: [599](https://github.com/unfoldingWord/tc-create-app/pull/599)

## Notes from review

New start and build scripts in package json:
```
    "start": "cross-env NODE_ENV=development REACT_APP_BUILD_NUMBER=$(cat public/build_number) rescripts start",
    "build": "cross-env NODE_ENV=production REACT_APP_BUILD_NUMBER=$(cat public/build_number) rescripts build && yarn styleguide:build",
```

A build number is captured and made available to the build and run process.
As env variable `REACT_APP_BUILD_NUMBER`

Package cross-env is a cross platform way of setting env vars.

Two more new scripts:
```
    "postinstall": "husky install",
    "increment-build": "bash scripts/increment-build.sh && git add -A"
```

The bash increment script:
```sh
#!/bin/sh
build_hash_dir=public/build_number
build_hash=`cat "$build_hash_dir"`

build_hash_split=(${build_hash//-/ })
build_number=(${build_hash_split[0]})
build_commit_hash=$(git rev-parse --short HEAD)
build_number=$((build_number + 1))
echo "$build_number-$build_commit_hash" > "$build_hash_dir"
```
The build number file contains:
```
12-6d5f98d
```
So it takes the build number file, splits into `12` and `6d5f98d`,
then increments the number, re-combines them and over writes the file.

There is a new postinstall step, which runs "husky install"

*What is husky?* After you run `yarn install`, this:
```
$ husky install
husky - installing git hooks...
husky - done
Done in 152.97s.
```

At this point, husky is in node_modules. And it has a "scripts" folder with husky.sh in it.

There is a ".husky" folder with a pre-commit script that runs `yarn increment-build`. 
So I guess git runs the increment build script runs just prior to a commit.

I made a small change to the README and then committed:
```
$ git commit -a -m "link added to readme"
yarn run v1.22.5
$ bash scripts/increment-build.sh && git add -A
Done in 0.48s.
[feature-jay-build-number 713808e] link added to readme
 2 files changed, 5 insertions(+), 1 deletion(-)
$ 
```
So the pre-commit hooks works. It ran the script and then did `git add -A`.
This means you will never catch public/build_number in uncommitted state.

## Proposed new placement

translationCore Create - v1.1.0-rc.1 build 13-6777a79

I had to update the Application Bar component in the `gitea-react-toolkit` to accept a new attribute 'build'.
This is used then to form the title on the app bar.
