package router

import (
	"net/http"
)

// Node is a struct representing one node of our router-tree
type node struct {
	component		string					// The path of this node
	isNamedParam	bool					// Whether it's a path with variables or not
	methods			map[string]http.Handler	// All methods "known" to this node, represented like ["GET"] >> HandlerGet
	children		[]*node					// All child-nodes
}

// Router is a struct containing a tree of nodes which can be traversed to come to the right handler and the defaultHandler
type Router struct {
	tree 					*node		 // Tree is a trie of nodes containing all of our paths
	pathNotFoundHandler 	http.Handler // Handler called when no node-path matches the requested url
	methodNotAllowedHandler	http.Handler // Handler called when the request method does not match the route
}