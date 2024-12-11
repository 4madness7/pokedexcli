package main

import "fmt"

func exploreCallback(c *config, args []string) error {
    if len(args) > 1 {
        return fmt.Errorf("Error: too many arguments. 1 arguments expected, %d provided.", len(args))
    }
    if len(args) < 1 {
        return fmt.Errorf("Error: too few arguments. 1 arguments expected, %d provided.", len(args))
    }

    return nil
}
