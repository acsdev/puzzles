package mirror

import (
	"testing"

	"gotest.tools/assert"
)

func TestMirrorIs(t *testing.T) {

	root := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 4,
			},
		},
		right: &Node{
			val: 2,
			left: &Node{
				val: 4,
			},
			right: &Node{
				val: 3,
			},
		},
	}

	result := IsMirror(root)
	assert.Equal(t, result, true)
}

func TestMirrorIsNot(t *testing.T) {

	root := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 4,
			},
		},
		right: &Node{
			val: 2,
			left: &Node{
				val: 3,
			},
			right: &Node{
				val: 4,
			},
		},
	}

	result := IsMirror(root)
	assert.Equal(t, result, false)

}
