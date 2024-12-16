**Objective:** Create a transpiler for the JSF language.

**Overview:**
JSF is a language designed for full stack web development, similar to JavaScript, but with some differences.

JSF is divided into three parts:

1. Server code
2. Client code
3. Static code

**Transpilation Rules:**

- **Server code** is transpiled into Go.
- **Client code** is transpiled into JavaScript.
- **Static code** is transpiled into Go.

**Execution:**

- Server code: Executed upon request.
- Static code: Executed during project build, resulting in cached HTML for immediate client delivery.
- Client code: Executed in the client browser.

**JSF Structure:**

- JSF files consist of components (with inner functions) and outer functions.
- Components can be classified as server, client, or static.
- Functions can be server or client.
- Only components can be static.

**Component and Function Rules:**

- **Server components**: May contain both server and client functions.
- **Client components**: May contain only client functions.
- **Static components**: May contain only server functions.

**File Composition:**

- A JSF file can contain multiple components and functions.

**Example of a JSF file:**

```jsf
server component <ComponentName param1 type1 param2 type2>{
    onRender( client (p1, p2) {
          console.log(p1, p2)
    })

    // server code
    fmt.println("Server code")
    serverCode()

    client function clientCode(){
        console.log("Client code")
    }

    server function serverCode(){
        fmt.println("Server function")
    }

    return <div>
        <h1>JSF Component Example</h1>
    </div>
}

client component <ClientComponent param1 param2>{

    // All code runs fully on the client
    console.log("Client component rendered")

    function clientCode(){
        console.log("Client function")
    }

    return <div>
        <h1>Client Component</h1>
    </div>
}

static component <StaticComponent param1 type1 param2 type2>{

    function staticCode(){
        fmt.println("Static function")
    }

    return <div>
        <h1>Static Component</h1>
    </div>
}
```
