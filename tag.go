package main

// The tag tree is used to look up music clips
// using fuzzy matching (so that the name does not
// have to be typed 100% correct on the command line)
// A tag may represent, e.g., an artist, album, clip...
// Some tags, like albums, have child Tags while
// others, like clips, are leaf nodes.
// The leaf nodes point to 