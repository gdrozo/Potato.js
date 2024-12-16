package main

import (
	"fmt"
	"regexp"
	"strings"
)

func transpileJSF(jsfCode string) (string, error) {
	// Split the code into components
	// Getting the starting index (int) inside jsfCode of the components declaration
	componentRegex := regexp.MustCompile(`(server|client|static)[\s\n]+component[\s\n]+<([^>]+)>[\s\n]*{`)
	componentsIndexes := componentRegex.FindAllStringIndex(jsfCode, -1)

	fmt.Println(componentsIndexes)

	//lines := strings.Split(jsfCode, "\n")

	for i, index := range componentsIndexes {
		component := jsfCode[index[0]:]
		if i < len(componentsIndexes)-1 {
			component = jsfCode[index[0]:componentsIndexes[i+1][0]]
		}

		curlyBracesCount := 0
		started := false
		componentEndIndex := -1

		for j, char := range component {
			if char == '{' {
				curlyBracesCount++
				started = true
			} else if char == '}' {
				curlyBracesCount--
			}
			
			if curlyBracesCount == 0 && started {
				componentEndIndex = j
				break
			}
		}
		component = component[0:componentEndIndex+1]
		fmt.Println(component)
	}

	var transpiledCode strings.Builder

	return transpiledCode.String(), nil
}

func extractCodeBlock(componentContent, keyword string) string {
	// Implement logic to extract code blocks based on keywords
	// This is a placeholder; a more robust solution is needed.
	// ... (Implementation details) ...
	return ""
}

func main() {
	// Example usage
	jsfCode := `
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
        <h1>Jsf Component Example</h1>
    </div>
}

client function otherClientFunction(){
    console.log("Other client function")
}

client component <ClientComponent param1 param2>{

    //All code runs full on the client
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

`
	transpiled, err := transpileJSF(jsfCode)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(transpiled)
	}
}