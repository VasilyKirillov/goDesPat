/*
Using clone method of Album struct it can be prototyped
*/
package main

import "fmt"

type Album struct {
	title       string
	releaseYear int
	authors     []string
	tracks      []string
}

func (a *Album) clone() *Album {
	authors := make([]string, len(a.authors))
	tracks := make([]string, len(a.tracks))
	copy(authors, a.authors)
	copy(tracks, a.tracks)
	return &Album{
		title:       a.title,
		authors:     authors,
		tracks:      tracks,
		releaseYear: a.releaseYear,
	}
}

func (a *Album) String() string {
	return fmt.Sprintf("Album:%v by %v released in %v\nTracks: %v", a.title, a.authors, a.releaseYear, a.tracks)
}

func ReleaseAlbum(title string, releaseYear int, authors []string, tracks []string) Album {
	return Album{
		title:       title,
		authors:     authors,
		tracks:      tracks,
		releaseYear: releaseYear,
	}
}
