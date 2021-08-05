import "./wasm_exec.js"

if (globalThis.fetch === undefined) {
  // node
  var fs = require('fs');
  globalThis.fetchAsArrayBuffer = fs.promises.readFile
} else {
  // browser
  globalThis.fetchAsArrayBuffer = async (url) => await (await fetch(url)).arrayBuffer()
}

const go = new Go();
fetchAsArrayBuffer("main.wasm")
  .then(wasmcode => WebAssembly.instantiate(wasmcode, go.importObject))
  .then((result) => {
    go.run(result.instance)
    console.log("direct call", mylib.AddInts(2, 3))
    console.log("call returning AddInts function", mylib.GetAddIntsFunction()(2, 3))
    console.log("started async fib",
      mylib.AsyncFib(10, (res) => console.log("async fib return: ", res)))
    console.log(mylib.SumAndProdInts([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))
    console.log(mylib.CombineName({}))
    console.log(mylib.CombineName({first: "Reinoud"}))
    console.log(mylib.CombineName({last: "Elhorst"}))
    console.log(mylib.CombineName({first: "Reinoud", last: "Elhorst"}))
})
