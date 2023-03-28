# Info on OPFS

## Links

https://webkit.org/blog/12257/the-file-system-access-api-with-origin-private-file-system/

https://sqlite.org/wasm/doc/trunk/persistence.md#opfs

https://fs.spec.whatwg.org/ (seems to be location of the standard)

https://chrome.google.com/webstore/detail/opfs-explorer/acndjpgkpaclldomagafnognkcgjignd#:~:text=OPFS%20Explorer%20is%20a%20Chrome,%2F)%20of%20a%20web%20application.

https://chrome.google.com/webstore/detail/opfs-explorer/acndjpgkpaclldomagafnognkcgjignd

https://developer.chrome.com/articles/file-system-access/#accessing-the-origin-private-file-system

https://stackoverflow.com/questions/44094507/how-to-store-large-files-to-web-local-storage

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


Another example from Apple for Safari:
```js
async function doOpfsDemo() {

    // Open the "root" of the website's (origin's) private filesystem:
    let storageRoot = null;
    try {
        storageRoot = await navigator.storage.getDirectory();
    }
    catch( err ) {
        console.error( err );
        alert( "Couldn't open OPFS. See browser console.\n\n" + err );
        return;
    }

    // Get the <canvas> element from the page DOM:
    const canvasElem = document.getElementById( 'myCanvas' );

    // Save the image:
    await saveCanvasToPngInOriginPrivateFileSystem( storageRoot, canvasElem );

    // (Re-)load the image:
    await loadPngFromOriginPrivateFileSystemIntoCanvas( storageRoot, canvasElem );
}

async function saveCanvasToPngInOriginPrivateFileSystem( storageRoot, canvasElem ) {

    // Save the <canvas>'s image to a PNG file to an in-memory Blob object: (see https://stackoverflow.com/a/57942679/159145 ):
    const imagePngBlob = await new Promise(resolve => canvasElem.toBlob( resolve, 'image/png' ) );

    // Create an empty (zero-byte) file in a new subdirectory: "art/mywaifu.png":
    const newSubDir = await storageRoot.getDirectoryHandle( "art", { "create" : true });
    const newFile   = await newSubDir.getFileHandle( "mywaifu.png", { "create" : true });

    // Open the `mywaifu.png` file as a writable stream ( FileSystemWritableFileStream ):
    const wtr = await newFile.createWritable();
    try {
        // Then write the Blob object directly:
        await wtr.write( imagePngBlob );
    }
    finally {
        // And safely close the file stream writer:
        await wtr.close();
    }
}

async function loadPngFromOriginPrivateFileSystemIntoCanvas( storageRoot, canvasElem ) {
    
    const artSubDir = await storageRoot.getDirectoryHandle( "art" );
    const savedFile = await artSubDir.getFileHandle( "mywaifu.png" ); // Surprisingly there isn't a "fileExists()" function: instead you need to iterate over all files, which is odd... https://wicg.github.io/file-system-access/
    
    // Get the `savedFile` as a DOM `File` object (as opposed to a `FileSystemFileHandle` object):
    const pngFile = await savedFile.getFile();
    
    // Load it into an ImageBitmap object which can be painted directly onto the <canvas>. You don't need to use URL.createObjectURL and <img/> anymore. See https://developer.mozilla.org/en-US/docs/Web/API/createImageBitmap
    // But you *do* still need to `.close()` the ImageBitmap after it's painted otherwise you'll leak memory. Use a try/finally block for that.
    try {
        const loadedBitmap = await createImageBitmap( pngFile ); // `createImageBitmap()` is a global free-function, like `parseInt()`. Which is unusual as most modern JS APIs are designed to not pollute the global scope.
        try {
            const ctx = canvasElem.getContext('2d');
            ctx.clearRect( /*x:*/ 0, /*y:*/ 0, ctx.canvas.width, ctx.canvas.height ); // Clear the canvas before drawing the loaded image.
            ctx.drawImage( loadedBitmap, /*x:*/ 0, /*y:*/ 0 ); // https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawImage
        }
        finally {
            loadedBitmap.close(); // https://developer.mozilla.org/en-US/docs/Web/API/ImageBitmap/close
        }
    }
    catch( err ) {
        console.error( err );
        alert( "Couldn't load previously saved image into <canvas>. See browser console.\n\n" + err );
        return;
    }
}
```