package main

// This file implements the Nodes of a filesystem tree.

import (
	"fmt"
	"io"
)

// Node in the Library's file tree.
type Node struct {
	payload  inte