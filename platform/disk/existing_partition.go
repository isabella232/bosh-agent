package disk

type ExistingPartition struct {
	Index        int
	SizeInBytes  uint64
	StartInBytes uint64
	EndInBytes   uint64
	Type         PartitionType
	Name         string
}
