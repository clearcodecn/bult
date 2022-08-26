package blut

import (
	"context"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

const (
	bucketMeta = "blut_meta"
)

type Engine struct {
	db *bolt.DB

	codec Codec
}

func (e *Engine) createNs(ctx context.Context, ns *Namespace) error {
	return e.db.Update(func(tx *bolt.Tx) error {
		// 1. 创建 bucket
		if _, err := tx.CreateBucket(s2b(ns.Name)); err != nil {
			return errors.Wrap(err, "failed to create ns")
		}
		// 2. 往 meta 表写数据.
		metaBucket, err := tx.CreateBucketIfNotExists(s2b(bucketMeta))
		if err != nil {
			return errors.Wrap(err, "failed to get meta bucket")
		}
		id, err := metaBucket.NextSequence()
		if err != nil {
			return errors.Wrap(err, "failed to get next sequence")
		}
		m := newMeta(id, ns)
		data, err := e.codec.Marshal(m)
		if err != nil {
			return errors.Wrap(err, "failed to marshal meta data")
		}
		metaBucket.Put(m.key())
	})
}
