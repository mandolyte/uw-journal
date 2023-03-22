# Info on OPFS

## Links

https://webkit.org/blog/12257/the-file-system-access-api-with-origin-private-file-system/

https://sqlite.org/wasm/doc/trunk/persistence.md#opfs

https://fs.spec.whatwg.org/ (seems to be location of the standard)

https://chrome.google.com/webstore/detail/opfs-explorer/acndjpgkpaclldomagafnognkcgjignd#:~:text=OPFS%20Explorer%20is%20a%20Chrome,%2F)%20of%20a%20web%20application.

https://chrome.google.com/webstore/detail/opfs-explorer/acndjpgkpaclldomagafnognkcgjignd

https://developer.chrome.com/articles/file-system-access/#accessing-the-origin-private-file-system

Some example code here I can use...
https://developer.mozilla.org/en-US/docs/Web/API/File_System_Access_API#synchronously_reading_and_writing_files_in_opfs
Here is the code snippet:
```js
onmessage = async (e) => {
  // retrieve message sent to work from main script
  const message = e.data;

  // Get handle to draft file in OPFS
  const root = await navigator.storage.getDirectory();
  const draftHandle = await root.getFileHandle("draft.txt", { create: true });
  // Get sync access handle
  const accessHandle = await draftHandle.createSyncAccessHandle();

  // Get size of the file.
  const fileSize = accessHandle.getSize();
  // Read file content to a buffer.
  const buffer = new DataView(new ArrayBuffer(fileSize));
  const readBuffer = accessHandle.read(buffer, { at: 0 });

  // Write the message to the end of the file.
  const encoder = new TextEncoder();
  const encodedMessage = encoder.encode(message);
  const writeBuffer = accessHandle.write(encodedMessage, { at: readBuffer });

  // Persist changes to disk.
  accessHandle.flush();

  // Always close FileSystemSyncAccessHandle if done.
  accessHandle.close();
};

```