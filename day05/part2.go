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
// seed and map overlap partially:
//
//	seed starts first -> two sections: 1) start of seed to start of map 2) start of map to end of seed
//	map starts first -> two sections: 1) start of seed to end of map 2) end of map to end of seed
func (m *Map) MapRange(seedRanges []SeedRange) []SeedRange {
	results := []SeedRange{}
	for i := 0; i < len(seedRanges); i++ {
		sr := seedRanges[i]
		srEnd := sr.Start + sr.Length - 1
		matched := false
		for _, mr := range m.Ranges {
			mrEnd := mr.Source + mr.Length - 1
			// no overlap with this map range
			if sr.Start > mrEnd || srEnd < mr.Source {
				continue
			}
			matched = true

			if sr.Start >= mr.Source {
				if srEnd <= mrEnd {
					// seed is entirely contained by map
					results = append(results, SeedRange{Start: mr.Destination + sr.Start - mr.Source, Length: sr.Length})
					continue
				} else {
					// partial overlap: map starts first
					seedRanges = append(seedRanges, SeedRange{Start: mr.Source + mr.Length, Length: mr.Source + mr.Length - sr.Start})
					results = append(results, SeedRange{Start: mr.Destination + sr.Start - mr.Source, Length: mr.Length - (sr.Start - mr.Source)})
					continue
				}
			} else {
				if mrEnd <= srEnd {
					// map is entirely contained by seed
					// come back to the two unmapped parts later
					seedRanges = append(seedRanges, SeedRange{Start: sr.Start, Length: mr.Source - sr.Start})
					seedRanges = append(seedRanges, SeedRange{Start: mr.Source + mr.Length, Length: srEnd - mrEnd})
					results = append(results, SeedRange{Start: mr.Destination, Length: mr.Length})
					continue
				} else {
					// partial overlap: seed starts first
					seedRanges = append(seedRanges, SeedRange{Start: sr.Start, Length: mr.Source - sr.Start})
					results = append(results, SeedRange{Start: mr.Destination, Length: sr.Length - (mr.Source - sr.Start)})
					continue
				}
			}
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
	lines, err := advent.ReadInput()
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

	for _, m := range maps {
		seedRanges = m.MapRange(seedRanges)
	}
	fmt.Println(seedRanges[0].Start)
}
