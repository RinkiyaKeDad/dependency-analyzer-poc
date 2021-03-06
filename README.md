# depstat

`depstat` is a CLI for analyzing dependencies of Go modules enabled projects. 

![depstat demo with k8s repo](./depstat-k8s-demo.gif)

## Installation 
Currently, to test `depstat` out you can generate the binary directly by running `go build` in the project directory and then pointing to the location of the binary from the project whose dependencies you want to analyze. Releases will be supported once the repository for `depstat` gets finalized.

## Usage
`depstat` can be used as a standalone command-line application. It is also to be used in the [Kubernetes](https://github.com/kubernetes/kubernetes) CI pipeline. For more info on the latter see [#100792](https://github.com/kubernetes/kubernetes/pull/100792).

## Commands

To see the list of commands `depstat` offers you can run `depstat help`. `depstat` currently supports the following commands:

### cycles

`depstat cycles` shows all the cycles present in the dependencies of the project.

An example of a cycle in project dependenies is:
`golang.org/x/net -> golang.org/x/crypto -> golang.org/x/net`

`--json` prints the output of the cycles command in JSON format. For the above example the JSON output would look like this:
```
{
  "cycles": [
    [
      "golang.org/x/net",
      "golang.org/x/crypto",
      "golang.org/x/net"
    ]
  ]
}
```

### graph

`depstat graph` will generate a `graph.dot` file which can be used with [Graphviz](https://graphviz.org)'s dot command to visualize the dependencies of a project.

For example, after running `depstat graph`, an SVG can be created using:
`twopi -Tsvg -o dag.svg graph.dot`

### list

`depstat list` shows a sorted list of all project dependencies. These include both direct and transitive dependencies.

1. Direct dependencies: Dependencies that are directly used in the code of the project. These do not include standard go packages like `fmt`, etc.

2. Transitive dependencies: These are dependencies that get imported because they are needed by some direct dependency of the project.

### stats

`depstat stats` will provide the following metrics about the dependencies of the project:

1. Total Dependencies: Total number of dependencies of the project. This is the sum of direct and transitive dependencies.

2. Max Depth of Dependencies: Number of dependencies in the longest dependency chain present in the project.

3. Transitive Dependencies: Total number of transitive dependencies.

- The `--json` flag gives this output in a JSON format.
- `--verbose` mode will help provide you with the list of all the dependencies and will also print the longest dependency chain.

## Project Goals
`depstat` is being developed under the code organization sub-project under [SIG Architecture](https://github.com/kubernetes/community/tree/master/sig-architecture). The goal is to make it easy to evaluate dependency updates to Kubernetes. This will be done by running `depstat` as part of the Kubernetes CI pipeline.

## Community Contact Information
You can reach the maintainers of this project at:

[#k8s-code-organization](https://kubernetes.slack.com/messages/k8s-code-organization) on the [Kubernetes slack](http://slack.k8s.io).

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

