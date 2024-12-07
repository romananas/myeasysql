package myeasysql

func (d DB) QueryRow(query string, dest any, args ...any) error {
	keys := _ParseQuerys(query)

	ptrs, err := _GetPointers(dest)
	tags := _ReadTags(dest)
	names := _ReadNames(dest)
	if err != nil {
		return err
	}
	order := _SortKeys(tags, names, keys)
	var sorted []any
	for _, i := range order {
		sorted = append(sorted, ptrs[i])
	}
	ptrs = sorted
	return d.db.QueryRow(query, args...).Scan(ptrs...)
}
