package ledger

// IndexableAttr represents an indexable attribute
type IndexableAttr string

// constants for indexable attributes
const (
	IndexableAttrBlockNum         = IndexableAttr("BlockNum")
	IndexableAttrBlockHash        = IndexableAttr("BlockHash")
	IndexableAttrTxID             = IndexableAttr("TxID")
	IndexableAttrBlockNumTranNum  = IndexableAttr("BlockNumTranNum")
	IndexableAttrBlockTxID        = IndexableAttr("BlockTxID")
	IndexableAttrTxValidationCode = IndexableAttr("TxValidationCode")
)

// IndexConfig - a configuration that includes a list of attributes that should be indexed
type IndexConfig struct {
	AttrsToIndex []IndexableAttr
}
