This repo contains the code for a [blog post](https://blog.claude.nl/tech/howto/interface-between-go-1.16-and-javascript-syscall-js/).

In order to run the code (you need to have Go 1.16 installed):

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOOS=js GOARCH=wasm go build -o main.wasm
```

you can now run the examples in node (note: the `.mjs` extension tells node to load this file a as javascript module).

```
node main.mjs
```


To run the code in a browser, you *have to* run a webserver and access your code from there.
Make sure to open the developer tools in your browser; all output and interaction will be in the javascript console.

Way to start a very simple webserver (if you have python3 installed)
```
python3 -m http.server
```

Now go to http://localhost:8000/index.html in a browser
