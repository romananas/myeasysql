package myeasysql

import (
	"fmt"
	"reflect"
)

// QueryRow executes an SQL query that is expected to return a single row of results.
// It maps the retrieved values to the fields of a given structure.
//
// Parameters:
// - query: The SQL query to execute (e.g., "SELECT id, username FROM users WHERE id = ?").
// - dest: A pointer to a struct where the query results will be mapped.
// - args: (optional) A variadic list of arguments for the SQL query.
//
// How it works:
// 1. Parses the query to extract the required column keys using `_ParseQuerys`.
// 2. Retrieves the pointers to the fields of the struct `dest` using `_GetPointers`.
// 3. Reads the tags (if defined) and names of the fields in the struct `dest` via `_ReadTags` and `_ReadNames`.
// 4. Reorders the pointers to match the order of columns in the query using `_SortKeys`.
// 5. Executes the SQL query and maps the retrieved values to the pointers of the struct fields.
//
// Returns:
// - `error`: Returns an error in case of failure (e.g., query execution or mapping issues), otherwise `nil`.
//
// Example Usage:
//
//	type User struct {
//	    ID       int    `db:"id"`
//	    Username string `db:"username"`
//	}
//
//	func main() {
//	    db := DB{...} // Initialize your DB instance
//	    var user User
//	    err := db.QueryRow("SELECT id, username FROM users WHERE id = ?", &user, 1)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("User: %+v\n", user)
//	}
func (d DB) QueryRow(query string, dest any, args ...any) error {
	// get row column as theres no columns method for *sql.Row
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

// Query executes an SQL query that is expected to return multiple rows of results.
// It maps the retrieved values to a slice of structs.
//
// Parameters:
// - query: The SQL query to execute (e.g., "SELECT id, username FROM users").
// - dest: A pointer to a slice of structs where the query results will be mapped.
// - args: (optional) A variadic list of arguments for the SQL query.
//
// Example Usage:
//
//	type User struct {
//	    ID       int    `db:"id"`
//	    Username string `db:"username"`
//	}
//
//	func main() {
//	    db := DB{...} // Initialize your DB instance
//	    var users []User
//	    err := db.Query("SELECT id, username FROM users", &users)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//	    fmt.Printf("Users: %+v\n", users)
//	}
func (d DB) Query(query string, dest any, args ...any) error {
	destVal := reflect.ValueOf(dest)
	if !_IsSlice(destVal) {
		return fmt.Errorf("dest must be a pointer to a slice")
	}

	// Obtenir le type de l'élément du slice
	if !_IsStruct(destVal) {
		return fmt.Errorf("slice elements must be structs")
	}
	elemType := destVal.Elem().Type().Elem()
	// Exécuter la requête
	rows, err := d.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Read all columns name of the query
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Préparer un slice temporaire pour contenir les résultats
	results := reflect.MakeSlice(destVal.Elem().Type(), 0, 0)

	for rows.Next() {
		// Créer une nouvelle instance de l'élément du slice
		elem := reflect.New(elemType).Elem()

		// Obtenir les pointeurs vers les champs correspondants
		ptrs, err := _GetPointers(elem.Addr().Interface())
		if err != nil {
			return err
		}

		// Réorganiser les pointeurs selon l'ordre des colonnes
		tags := _ReadTags(elem.Addr().Interface())
		names := _ReadNames(elem.Addr().Interface())
		order := _SortKeys(tags, names, columns)
		var sorted []any
		for _, i := range order {
			sorted = append(sorted, ptrs[i])
		}

		// Scanner les valeurs de la ligne
		if err := rows.Scan(sorted...); err != nil {
			return err
		}

		// Ajouter l'élément au slice
		results = reflect.Append(results, elem)
	}

	// Assigner les résultats au slice de destination
	destVal.Elem().Set(results)

	// Vérifier les erreurs éventuelles de parcours des lignes
	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
