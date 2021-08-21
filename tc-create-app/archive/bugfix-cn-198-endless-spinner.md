# bugfix-cn-198-endless-spinner

Reference: https://forum.door43.org/t/rcl-app-development-process/605

## 2020-07-22

Findings so far at:
https://github.com/unfoldingWord/tc-create-app/issues/198

And text snippet:
>I have been stepping thru the debugger trying to find out why filepath is persisted on a logout, when "keep me logged in" is NOT checked. All the other keys are removed from IndexedDB, but not this one. I found that it is being removed but is immediately put back by Gitea React Toolkit's useFile.js -- that was a surprise. For documentation purposes and in case I do not figure it out before lunch, I wanted to put this here so you can take it up if you wish. Just before this, I saw it go into the save state function and remove it. Then the debugger paused here. If you look at the call stack, you will see that somehow, the custom hook usefile is putting it back.


1. use yalc to publish GRT locally
2. use yalc to link local store into tc-create-app