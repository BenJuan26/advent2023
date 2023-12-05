package main

import (
	"fmt"
	"sort"
	"strings"

	advent "github.com/BenJuan26/advent2023"
)

type SeedRange struct {
	Start  int
	Length int
}

// possibilities:
// seed and map have no overlap -> add seed range verbatim
// seed is entirely contained by map -> map of seed start to map of seed end
// map is entirely contained by seed -> three sections: 1) start of seed to start of overlap 2) mapped overlap section 3) end of map to end of seed
// seed and map overlap partially
//   seed starts first -> two sections: 1) start of seed to start of map 2) start of map to end of seed
//   map starts first -> two sections: 1) start of seed to end of map 2) end of map to end of seed

func (m *Map) MapRange(seedRanges []SeedRange) []SeedRange {
	results := []SeedRange{}
	for i := 0; i < len(seedRanges); i++ {
		sr := seedRanges[i]
		matched := false
		for _, mr := range m.Ranges {
			// no overlap with this map range
			if sr.Start > (mr.Source+mr.Length-1) || sr.Start+sr.Length-1 < mr.Source {
				continue
			}
			matched = true
			// seed is entirely contained by map
			if sr.Start >= mr.Source && sr.Start+sr.Length-1 <= mr.Source+mr.Length-1 {
				results = append(results, SeedRange{Start: mr.Destination + sr.Start - mr.Source, Length: sr.Length})
				continue
			}
			// map is entirely contained by seed
			if mr.Source >= sr.Start && mr.Source+mr.Length-1 <= sr.Start+sr.Length-1 {
				// come back to the two unmapped parts later
				seedRanges = append(seedRanges, SeedRange{Start: sr.Start, Length: mr.Source - sr.Start})
				seedRanges = append(seedRanges, SeedRange{Start: mr.Source + mr.Length, Length: sr.Start + sr.Length - (mr.Source + mr.Length)})
				results = append(results, SeedRange{Start: mr.Destination, Length: mr.Length})
				continue
			}
			// partial overlap: seed starts first
			if sr.Start < mr.Source && sr.Start+sr.Length < mr.Source+mr.Length {
				seedRanges = append(seedRanges, SeedRange{Start: sr.Start, Length: mr.Source - sr.Start})
				results = append(results, SeedRange{Start: mr.Destination, Length: sr.Length - (mr.Source - sr.Start)})
				continue
			}
			// partial overlap: map starts first
			if mr.Source < sr.Start && mr.Source+mr.Length < sr.Start+sr.Length {
				seedRanges = append(seedRanges, SeedRange{Start: mr.Source + mr.Length, Length: mr.Source + mr.Length - sr.Start})
				results = append(results, SeedRange{Start: mr.Destination + sr.Start - mr.Source, Length: mr.Length - (sr.Start - mr.Source)})
				continue
			}
			panic("shouldn't reach here")
		}
		if !matched {
			results = append(results, sr)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Start < results[j].Start
	})

	return results
}

func Part2() {
	lines, err := advent.ReadTestInput()
	if err != nil {
		panic(err)
	}

	seedRanges := []SeedRange{}
	seedsString := strings.Split(lines[0], ": ")[1]
	seedsStrings := strings.Fields(seedsString)
	for i := 0; i < len(seedsStrings); i += 2 {
		seedRanges = append(seedRanges, SeedRange{Start: advent.MustAtoi(seedsStrings[i]), Length: advent.MustAtoi(seedsStrings[i+1])})
	}

	maps := PopulateMaps(lines)

	for _, sr := range seedRanges {
		fmt.Printf("%d -> %d, ", sr.Start, sr.Start+sr.Length-1)
	}
	fmt.Println("")
	for _, m := range maps {
		seedRanges = m.MapRange(seedRanges)
		for _, sr := range seedRanges {
			fmt.Printf("%d -> %d, ", sr.Start, sr.Start+sr.Length-1)
		}
		fmt.Println("")
	}
	// sort.Slice(seedRanges, func(i, j int) bool {
	// 	return seedRanges[i].Start < seedRanges[j].Start
	// })
	fmt.Println(seedRanges[0].Start)
}
