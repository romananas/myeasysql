package myeasysql

import (
	"fmt"
	"reflect"
)

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
	// Vérifier que `dest` est un pointeur vers un slice
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a pointer to a slice")
	}

	// Obtenir le type de l'élément du slice
	elemType := destVal.Elem().Type().Elem()
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf("slice elements must be structs")
	}
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
		ptrs, err := getPointers(elem.Addr().Interface())
		if err != nil {
			return err
		}

		// Réorganiser les pointeurs selon l'ordre des colonnes
		tags := readTags(elem.Addr().Interface())
		names := readNames(elem.Addr().Interface())
		order := sortKeys(tags, names, columns)
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
