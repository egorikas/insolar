// Code generated by "stringer -type=NodeState"; DO NOT EDIT.

package insolar

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NodeUndefined-0]
	_ = x[NodePending-1]
	_ = x[NodeReady-2]
	_ = x[NodeLeaving-3]
}

const _NodeState_name = "NodeUndefinedNodePendingNodeReadyNodeLeaving"

var _NodeState_index = [...]uint8{0, 13, 24, 33, 44}

func (i NodeState) String() string {
	if i >= NodeState(len(_NodeState_index)-1) {
		return "NodeState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeState_name[_NodeState_index[i]:_NodeState_index[i+1]]
}