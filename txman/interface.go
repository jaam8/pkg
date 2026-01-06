package txman

import "context"

type Txman interface {
	Do(ctx context.Context, h Handler, opts ...TxOption) error
}
