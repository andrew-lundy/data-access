package main

import (
	"fmt"
	"testing"
)

func albumSlicesAreEqual(a, b []Album) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAlbumsByArtist(t *testing.T) {
	name := "John Coltrane"

	wantedResult := make([]Album, 0)
	wantedResult = append(wantedResult, Album{
		ID:     1,
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	})
	fmt.Printf("%v", wantedResult)

	albums, err := AlbumsByArtist(name)
	fmt.Printf("%v & %v", albums, err)

	//if err != nil {
	//	t.Fatalf("Error: %v. Received: %v", err, albums)
	//}

	//gotWantedResult := albumSlicesAreEqual(wantedResult, albums)
	//fmt.Println(gotWantedResult)

	//wantedResult = append(wantedResult,
	//	Album{
	//		ID:     1,
	//		Title:  "Blue Train",
	//		Artist: "John Coltrane",
	//		Price:  56.99,
	//	},
	//	Album{
	//		ID:     2,
	//		Title:  "Giant Steps",
	//		Artist: "John Coltrane",
	//		Price:  63.99,
	//	})
	//albums, err := albumsByArtist(name)
	//
	//gotWantedResult := albumSlicesAreEqual(wantedResult, albums)
	//
	//if !gotWantedResult || err != nil {
	//	t.Fatalf(`Expected: %v. Received: %v.`, wantedResult, albums)
	//}
}
