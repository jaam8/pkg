package txman

import (
	"time"

	"github.com/jackc/pgx/v5"
)

type TxOption func(*TxConfig)

func WithRetry(n uint) TxOption               { return func(c *TxConfig) { c.Retry = n } }
func WithIso(lvl pgx.TxIsoLevel) TxOption     { return func(c *TxConfig) { c.IsoLevel = lvl } }
func ReadOnly(on bool) TxOption               { return func(c *TxConfig) { c.ReadOnly = on } }
func WithMaxBackoff(d time.Duration) TxOption { return func(c *TxConfig) { c.MaxBackoff = d } }
