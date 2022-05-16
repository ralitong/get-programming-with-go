package main

import "fmt"

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capcity %v, %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")

	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
	
	// Modify dwarfs1 if it affects dwarfs2 and dwarfs3
	dwarfs1[0] = "New Ceres"
	fmt.Println("Modifying dwarfs1[0] = \"New Ceres\"")
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
	
	// Modify dwarfs3 if it affects dwarfs1 and dwarfs2
	dwarfs3[1] = "Pluto!"
	fmt.Println("Modifying dwarfs3[1] = \"Pluto!\"")
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
	dump("dwarfs3", dwarfs3)
}