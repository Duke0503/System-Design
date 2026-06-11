// Day 1 — What is a System?
//
// The simplest possible system: something receives input,
// does something with it, and returns output.
//
// That's it. That's a system.
// Everything we learn from here is just this idea getting more complicated.

package main

import "fmt"

// A Request is what comes IN to our system.
type Request struct {
	From    string
	Message string
}

// A Response is what goes OUT.
type Response struct {
	To      string
	Message string
}

// process is the system doing its job.
// Input goes in. Something happens. Output comes out.
func process(req Request) Response {
	return Response{
		To:      req.From,
		Message: "Got your message: " + req.Message,
	}
}

func main() {
	// A user sends a request
	req := Request{
		From:    "Duke",
		Message: "hello, system",
	}

	fmt.Printf("→ Request  from: %s | message: %q\n", req.From, req.Message)

	// The system processes it
	res := process(req)

	fmt.Printf("← Response to:   %s | message: %q\n", res.To, res.Message)

	// That's a system. One request. One response. One purpose.
	//
	// Tomorrow's question: what happens when TWO users send requests
	// at the same time? And where does the data go when the program stops?
}
