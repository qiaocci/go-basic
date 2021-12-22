package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// 跟python里的sorted类似
// sorted(iterable, /, *, key=None, reverse=False)
func main() {
	//sort.Interface()
	scl1 := []string{"abc", "rwp", "def", "por", "ber", "erj"}
	scl2 := []string{"Rabbit", "Fish", "Dog",
		"Parrot", "Cat", "Hamster"}

	// Sorting the slice of strings
	// Using Strings function
	sort.Strings(scl1)
	sort.Strings(scl2)

	// Displaying the result
	fmt.Println("\nSlices(After):")
	fmt.Println("Slice 1 : ", scl1)
	fmt.Println("Slice 2 : ", scl2)

	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		return x.Length < y.Length
	}})
	printTracks(tracks)

	values := []int{3, 1, 4, 1}
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))

	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
}

// Track 播放列表
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"爱不持久", "小s", "变态少女", 2001, length("4m11s")},
	{"浪流连", "茄子蛋", "我们以后要结婚", 2019, length("4m13s")},
	{"无与伦比的美丽", "苏打绿", "无与伦比的美丽", 2007, length("2m27s")},
	{"爱情的骗子我问你", "陈小云", "福建烧酒金曲", 1988, length("4m2s")},
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
	tw.Flush() // calculate column widths and print table
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}
func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
