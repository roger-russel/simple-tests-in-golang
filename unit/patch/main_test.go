package main

import (
	"fmt"
	"os"

	"bou.ke/monkey"
)

func Example_main() {

	defer monkey.UnpatchAll()

	monkey.Patch(os.MkdirAll, func(path string, perm os.FileMode) (err error) {
		return fmt.Errorf("permission denied")
	})

	main()
	// output:
	// permission denied

}
