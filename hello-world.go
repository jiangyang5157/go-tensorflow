package main

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {
	// Construct a graph with an operation that produces a string constant.
	root := op.NewScope()
	msg := op.Const(root, "Hello from TensorFlow version "+tf.Version())
	graph, err := root.Finalize()
	if err != nil {
		panic(err)
	}

	// Execute the graph in a session.
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	output, err := session.Run(nil, []tf.Output{msg}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(output[0].Value())
}
