package array

import (
	"fmt"
	"os"
	"text/template"
)

// ToDotGraphviz . base template to fill data
func (list *Skiplist) ToDotGraphviz(filename string) error {
	if filename == "" {
		filename = "./skiplist.dot"
	}

	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	tpl, err := template.New("skiplist.dot").Parse(graphvizDotTpl())
	if err != nil {
		return err
	}

	visitNode := list.head.levels[0].forward
	d := &graphvizData{
		Len:      list.length,
		Level:    list.level,
		P:        list.p,
		MaxLevel: list.maxLevel,
		Nodes:    make([]*graphvizNode, 0),
		Edges:    make([]*graphvizEdge, 0),
	}

	// append tail
	d.Nodes = append(d.Nodes, &graphvizNode{
		Key:        "NULL",
		Score:      0,
		LevelCount: list.level,
	})

	// append head
	d.Nodes = append(d.Nodes, &graphvizNode{
		Key:        list.head.key,
		Score:      list.head.score,
		LevelCount: list.level,
	})

	for i := 0; i < list.level; i++ {
		level := list.head.levels[i]
		d.Edges = append(d.Edges, &graphvizEdge{
			HeadKey: list.head.key,
			TailKey: level.forward.key,
			Level:   i + 1,
			Span:    level.span,
		})
	}

	// fill graphvizData
	for visitNode != nil {
		d.Nodes = append(d.Nodes, &graphvizNode{
			Key:        visitNode.key,
			Score:      visitNode.score,
			LevelCount: visitNode.level(),
		})

		for idx, level := range visitNode.levels {
			edge := &graphvizEdge{
				HeadKey: visitNode.key,
				TailKey: "NULL",
				Level:   idx + 1,
				Span:    level.span,
			}

			if level != nil && level.forward != nil {
				edge.TailKey = level.forward.key
			}

			d.Edges = append(d.Edges, edge)
		}

		visitNode = visitNode.levels[0].forward
	}

	// Mock:
	// var d = &graphvizData{
	// 	Nodes: []*graphvizNode{
	// 		&graphvizNode{
	// 			Score:      0.1,
	// 			Key:        "a",
	// 			LevelCount: 3,
	// 		},
	// 		&graphvizNode{
	// 			Score:      0.1,
	// 			Key:        "b",
	// 			LevelCount: 3,
	// 		},
	// 	},
	// 	Edges: []*graphvizEdge{
	// 		&graphvizEdge{
	// 			HeadKey: "a",
	// 			TailKey: "b",
	// 			Level:   1,
	// 		},
	// 		&graphvizEdge{
	// 			HeadKey: "a",
	// 			TailKey: "b",
	// 			Level:   2,
	// 		},
	// 		&graphvizEdge{
	// 			HeadKey: "a",
	// 			TailKey: "b",
	// 			Level:   3,
	// 		},
	// 	},
	// }
	if err = tpl.Execute(fd, d); err != nil {
		return err
	}

	return nil
}

type graphvizData struct {
	Len      uint
	Level    int
	P        float32
	MaxLevel int
	Nodes    []*graphvizNode
	Edges    []*graphvizEdge
}

type graphvizEdge struct {
	HeadKey string
	Level   int
	TailKey string
	Span    uint
}

type graphvizNode struct {
	Score      float32
	LevelCount int
	Key        string
}

func (node *graphvizNode) Label() string {
	s := ""
	for i := node.LevelCount; i > 0; i-- {
		s += fmt.Sprintf("<%d> level_%d |", i, i)
	}

	s += fmt.Sprintf("score: %.4f | key: %s", node.Score, node.Key)

	return s
}

func graphvizDotTpl() string {
	return `digraph skiplist {
	rankdir="LR"
	node [shape="record", height=".1"]

	skiplist [label="<length> length: {{.Len}} | <level> level: {{.Level}} | <p> p: {{.P}} | <max_level> max_level: {{.MaxLevel}} | <head> head_ptr | <tail> tail_ptr"]
	// nodes
	// nodename [label="<3> level 3 | <2> level 2 | <1> level  1 | <0> level 0 | head"]
	{{range $index, $element := .Nodes}}{{$element.Key}} [label="{{$element.Label}}"]
	{{end}}

	// edges
	// head:3 -> a:3
	skiplist:head -> head
	skiplist:tail -> NULL
	{{range $index, $element := .Edges}}{{$element.HeadKey}}:{{$element.Level}}->{{$element.TailKey}}:{{$element.Level}} [label="{{$element.Span}}"]
	{{end}} 
}`
}
