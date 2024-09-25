package node

// Node is an object representing any machine in the cluster.
// This defines the actual physical machine. Manager is a node, Workers are nodes
type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Role            string
	TaskCount       int
}
