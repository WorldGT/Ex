//主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
//定义一个sort.Interface的实现用在这样的表格中。

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Tracksort struct {
	t         []*Track
	primary   string
	secondary string
	third     string
}

func (x *Tracksort) Len() int      { return len(x.t) }
func (x *Tracksort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x *Tracksort) Less(i, j int) bool {
	key := x.primary
	for k := 0; k < 3; k++ {
		switch key {
		case "Title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "Year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "Length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}

		if k == 0 {
			key = x.secondary
		} else if k == 1 {
			key = x.third
		}
	}
	return false
}

// 更新排序键优先值
func setPrimary(x *Tracksort, p string) {
	x.primary, x.secondary, x.third = p, x.primary, x.secondary
}

func SetPrimary(x sort.Interface, p string) {
	if x, ok := x.(*Tracksort); ok {
		setPrimary(x, p)
	}
}

func NewTracksort(t []*Track, p, s, th string) sort.Interface {
	return &Tracksort{
		t:         t,
		primary:   p,
		secondary: s,
		third:     th,
	}
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	printTracks(tracks)

	Tracksort := NewTracksort(tracks, "Title", "Year", "Length")
	fmt.Println("\nCustom by Title, Year, Length:")
	sort.Sort(Tracksort)
	printTracks(tracks)

	fmt.Println("\nCustom by Year, Title, Length:")
	Tracksort = NewTracksort(tracks, "Year", "Title", "Length")
	sort.Sort(Tracksort)
	printTracks(tracks)

	fmt.Println("\nCustom select Length")
	SetPrimary(Tracksort, "Length")
	sort.Sort(Tracksort)
	printTracks(tracks)

}
