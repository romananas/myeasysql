package myeasysql

func (d DB) Exec(query string, args ...any) error {
	var arrArgs []any
	for arg := range args {
		tmp, err := getPointers(arg)
		if err != nil {
			return err
		}
		arrArgs = append(arrArgs, tmp...)
	}
	return nil
}
