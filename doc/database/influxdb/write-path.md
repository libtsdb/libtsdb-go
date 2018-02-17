# Write Path

Based on original notes from Xephon-K

- tsdb/engine/tsm1/engine.go `func (e *Engine) WritePoints(points []models.Point)`
  - `err := e.Cache.WriteMulti(values)` it writes to cache before write to WAL
  - `_, err = e.WAL.WriteMulti(values)`
- tsdb/engine/tsm1/cache.go use `storer` interface
  - `guru -scope . implements cache.go:#4432` 4432 is byte offset ... not line number, which works for editors, but not human ...
  - implemented by `ring`, `TestStore`, `emptyStore`
- first snapshot, then compactor, them tsmWriter, this is executed async, not sync when request come in
- tsdb/engine/tsm1/engine.go
  - `func (e *Engine) enableSnapshotCompactions()`
  - `func (e *Engine) compactCache()` 'continually checks if the WAL cache should be written to disk'
  - `func (e *Engine) WriteSnapshot() error`
  - `func (e *Engine) writeSnapshotAndCommit(closedFiles []string, snapshot *Cache) (err error)`
- tsdb/engine/tsm1/compact.go
  - `func (c *Compactor) WriteSnapshot(cache *Cache) ([]string, error) {`
  - `func (c *Compactor) writeNewFiles(generation, sequence int, iter KeyIterator, throttle bool) ([]string, error) {`
  - `func (c *Compactor) write(path string, iter KeyIterator, throttle bool) (err error) {`
- tsdb/engine/tsm1/writer.go
  - `func (t *tsmWriter) Write(key []byte, values Values)` saw this function several times when writing Xephon-K(S)
- [] TODO: it calls `Encode` etc. which is where all compression algorithm kicked in
   
````go
// NOTE: some branches and error handling are removed
func (e *Engine) WritePoints(points []models.Point) error {
    values := make(map[string][]Value, len(points))
    var keyBuf []byte
    for _, p := range points {
            keyBuf = append(keyBuf[:0], p.Key()...)
            keyBuf = append(keyBuf, keyFieldSeparator...)
            baseLen = len(keyBuf)
    		iter := p.FieldIterator()
    		t := p.Time().UnixNano()
    		for iter.Next() {
    			keyBuf = append(keyBuf[:baseLen], iter.FieldKey()...)
    			var v Value
    			switch iter.Type() {
    			case models.Float:
    				fv, err := iter.FloatValue()
    				v = NewFloatValue(t, fv)
    			case models.Integer:
    				iv, err := iter.IntegerValue()
    				v = NewIntegerValue(t, iv)
    			default:
    				return fmt.Errorf("unknown field type for %s: %s", string(iter.FieldKey()), p.String())
    			}
    			values[string(keyBuf)] = append(values[string(keyBuf)], v)
    		}
    	}
    }
    // first try to write to the cache
    err := e.Cache.WriteMulti(values)
	_, err = e.WAL.WriteMulti(values)
	return err
}
````