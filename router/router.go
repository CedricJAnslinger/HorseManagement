package router

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Slice of strings containing all supported HTTPMethods of our router.
var SupportedHttpMethods = []string{"GET"}

// NewRouter creates a new instance of the type router.
// @parameter - NFHandler: func(http.ResponseWriter, *http.Request) >> Function to be executed when a path was not found.
// @parameter - NAdHandler: func(http.ResponseWriter, *http.Request) >> Function to be executed when a path was not found.
// @return - *Router(Pointer) >> The new instance of router.
func NewRouter(NFHandler func(http.ResponseWriter, *http.Request), NAHandler func(http.ResponseWriter, *http.Request)) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]http.Handler)}
	return &Router{&node, http.HandlerFunc(NFHandler), http.HandlerFunc(NAHandler)}
}

// HandleFunc converts a func to a http.Handler and calls Handle afterwards
// @caller - r: *Router(Pointer): Instance of a router to which we add the new path.
// @param - method: string: The HttpMethod which is getting used for this path.
// @param - path: string: The path of the node we add to our api.
// @param - handler: func(http.ResponseWriter, *http.Request): The actual function which is getting executed at this endpoint.
func (r *Router) HandleFunc(method string, path string, handler func(http.ResponseWriter, *http.Request)) {
	// Use adapter http.HandlerFunc to allow the use of functions as HTTP handlers
	r.Handle(method, path, http.HandlerFunc(handler))
}

// Handle adds a new path to our router.
// @caller - r: *Router(Pointer): Instance of a router to which we add the new path.
// @param - method: string: The HttpMethod which is getting used for this path.
// @param - path: string: The path of the node we add to our api.
// @param - handler: http.Handler: The actual handler which is getting executed at this endpoint.
func (r *Router) Handle(method string, path string, handler http.Handler) {
	if len(path) == 0 || path[0] != '/' {
		log.Fatal("Couldn't add path: ", path, ", path has to start with a /")
	}

	// Check if the method is supported
	if !isSupportedHttpMethod(method) {
		log.Fatal("Couldn't add path: ", path, ", with method ", method, " ,method not supported")
	}
	// Use adapter http.HandlerFunc to allow the use of functions as HTTP handlers
	r.tree.addNode(method, path, handler)
}

// Implements the from the interface Handler needed ServeHTTP(ResponseWriter, *Request) method.
// From the docs: ServeHTTP should write reply headers and data to the ResponseWriter and then return.
// @caller - r: *Router(Pointer): Instance of a router
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Parse raw query from URL
	req.ParseForm()
	params := req.Form
	// Used to split url in "pure path" and params-part and then search for the needed node
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	handler := node.methods[req.Method]

	// Handler is not available,
	if handler == nil {
		// methods exist >> Requested method is not allowed
		if len(node.methods) > 0 {
			handler = r.methodNotAllowedHandler
		} else {
			// No method at all exists >> Path is not allowed
			handler = r.pathNotFoundHandler
		}
	}

	handler.ServeHTTP(w, req)
}

// addNode adds a new node to the tree and updates all existing nodes (where necessary).
// @caller - n: *node(Pointer) >> Instance of a node. Needed because the tree of a router has type node.
// @param - method: string: The HttpMethod which is getting used for this path.
// @param - path: string: The path of the node we add to our api.
// @param - handler: http.Handler >> The actual handler of the new added node.
func (n *node) addNode(method string, path string, handler http.Handler) {
	// Split path in several sub-components to check for each level what to do
	components := strings.Split(path, "/")[1:]
	for count := len(components); count > 0; count-- {
		aNode, component := n.traverse(components, nil)
		newNode := node{component: component, isNamedParam: false, methods: make(map[string]http.Handler)}

		// If component starts with : >> there is a named parameter in the route
		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}

		// Last iteration of the loop, decide where to put the handler
		if count == 1 {
			// Update an existing node when the path of aNode == component
			if aNode.component == component {
				aNode.methods[method] = handler
				return
			}
			// If not, attach the handler directly to the newNode
			newNode.methods[method] = handler
		}
		aNode.children = append(aNode.children, &newNode)
	}
}

// traverse iterates through the "child nodes" of the caller and checks which node are available,
// which node is the parent node of our new node and under which component(sub-path) we shall add the new node
// @caller - n: *node(Pointer) >> Instance of a node. Needed because the tree of a router has type node.
// @param - components: array(string) >> All sub-parts of an url
// @param - params: url.Values >> Named parameters to be added where needed when traversing
// @return - *node(Pointer) >> Instance of node which is the parent node >> Place where the new component shall be added as child
// @return - string >> String under which the new node shall be added
// @example >> components [website :name test1] >> components [:name test1] >> components [test1]
func (n *node) traverse(components []string, params url.Values) (*node, string) {
	// If component exists, it is always the parent component. exp: /tests/test1 >> /website = component
	// Like that, we can walk through the tree in recursive call
	component := components[0]
	// Iterate over all children of n, if none exists, directly return
	if len(n.children) > 0 {
		for _, child := range n.children {
			// If we need to update an existing component >> Recognized because component(path) == child.component(path)
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 {
					// recursive call to go through the tree
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}

// isSupportedHttpMethod checks whether a given method is supported by this router
// @param - method: string >> The method which we have to website
// @return - isSupported: bool >> The bool value whether the given method is supported or not
func isSupportedHttpMethod(method string) (isSupported bool) {
	isSupported = false
	for _, k := range SupportedHttpMethods {
		if k == method {
			isSupported = true
			return
		}
	}
	return isSupported
}


// AddDirectoryWeb makes a complete directory accessible by adding all elements to the api
// @param - directory: string >> The path to the directory we want to add
// @param - handler: http.Handler: The actual handler which is getting executed at this endpoint.
func (r *Router) AddDirectoryWeb(directory string, handler http.Handler) {
	listOfElements := make([]string, 0)

	err := filepath.Walk(directory, func(path string, f os.FileInfo, err error) error {
		// Ignore all directories and hidden files
		if f.IsDir() || f.Name()[0] == '.' {
			return nil
		}
		listOfElements = append(listOfElements, strings.TrimPrefix(path, directory))
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range listOfElements {
		r.Handle("GET", file, handler)
	}
}

// Redirect wraps a func(w http.ResponseWriter, r *http.Request) around a http.Redirect so that it can be used
// as handler for router.HandleFunc
// @param - newPath: string >> The path we want to redirect to
func Redirect(newPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, newPath, http.StatusMovedPermanently)
	}
}