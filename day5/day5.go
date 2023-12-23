package day5

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 5: If You Give A Seed A Fertilizer ---
You take the boat and find the gardener right where you were told he would be: managing a giant "garden" that looks
more to you like a farm.

"A water source? Island Island is the water source!" You point out that Snow Island isn't receiving any water.

"Oh, we had to stop the water because we ran out of sand to filter it with! Can't make snow with dirty water.
Don't worry, I'm sure we'll get more sand soon; we only turned off the water a few days... weeks... oh no."
His face sinks into a look of horrified realization.

"I've been so busy making sure everyone here has food that I completely forgot to check why we stopped getting more
sand! There's a ferry leaving soon that is headed over in that direction - it's much faster than your boat.
Could you please go check it out?"

You barely have time to agree to this request when he brings up another. "While you wait for the ferry,
maybe you can help us with our food production problem. The latest Island Island Almanac just arrived and
we're having trouble making sense of it."

The almanac (your puzzle input) lists all of the seeds that need to be planted. It also lists what type of soil
to use with each kind of seed, what type of fertilizer to use with each kind of soil, what type of water to use
with each kind of fertilizer, and so on. Every type of seed, soil, fertilizer and so on is identified with a number,
but numbers are reused by each category - that is, soil 123 and fertilizer 123 aren't necessarily related to each other.

For example:

seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
The almanac starts by listing which seeds need to be planted: seeds 79, 14, 55, and 13.

The rest of the almanac contains a list of maps which describe how to convert numbers from a source category into
numbers in a destination category. That is, the section that starts with seed-to-soil map: describes how to convert
a seed number (the source) to a soil number (the destination). This lets the gardener and his team know which soil
to use with which seeds, which water to use with which fertilizer, and so on.

Rather than list every source number and its corresponding destination number one by one, the maps describe entire
ranges of numbers that can be converted. Each line within a map contains three numbers: the destination range start,
the source range start, and the range length.

Consider again the example seed-to-soil map:

50 98 2
52 50 48
The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line
means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length,
but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds
to soil number 50 and that seed number 99 corresponds to soil number 51.

The second line means that the source range starts at 50 and contains 48 values: 50, 51, ..., 96, 97. This corresponds
to a destination range starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed number 53
corresponds to soil number 55.

Any source numbers that aren't mapped correspond to the same destination number. So, seed number 10 corresponds
to soil number 10.

So, the entire list of seed numbers and their corresponding soil numbers looks like this:

seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51
With this map, you can look up the soil number required for each initial seed number:

Seed number 79 corresponds to soil number 81.
Seed number 14 corresponds to soil number 14.
Seed number 55 corresponds to soil number 57.
Seed number 13 corresponds to soil number 13.
The gardener and his team want to get started as soon as possible, so they'd like to know the closest location that
needs a seed. Using these maps, find the lowest location number that corresponds to any of the initial seeds.
To do this, you'll need to convert each seed number through other categories until you can find its corresponding
location number. In this example, the corresponding types are:

Seed 79, soil 81, fertilizer 81, water 81, light 74, temperature 78, humidity 78, location 82.
Seed 14, soil 14, fertilizer 53, water 49, light 42, temperature 42, humidity 43, location 43.
Seed 55, soil 57, fertilizer 57, water 53, light 46, temperature 82, humidity 82, location 86.
Seed 13, soil 13, fertilizer 52, water 41, light 34, temperature 34, humidity 35, location 35.
So, the lowest location number in this example is 35.

What is the lowest location number that corresponds to any of the initial seed numbers?
*/

type MappingTable struct {
	Name     string
	Mappings []Mapping
}

type Mapping struct {
	SourceRange      Range
	DestinationRange Range
}

// Range represents an interval of the form [Start, End). The beginning is inclusive,
// and the ending is exclusive.
type Range struct {
	Start int
	End   int
}

func Part1(lines []string) (int, error) {
	seeds, maps := parseInput(lines)
	return findLowestLocation(seeds, maps), nil
}

func parseInput(lines []string) ([]int, []MappingTable) {
	sections := splitIntoSections(lines)

	seeds := parseSeeds(sections[0])
	maps := parseMappingTables(sections[1:])
	return seeds, maps
}

func splitIntoSections(lines []string) [][]string {
	sections := [][]string{}

	temp := []string{}
	for _, line := range lines {
		if line == "" {
			sections = append(sections, temp)
			temp = []string{}
		} else {
			temp = append(temp, line)
		}
	}
	// Make sure to add the last section.
	sections = append(sections, temp)
	return sections
}

func parseSeeds(lines []string) []int {
	// There is only 1 line for seeds.
	line := lines[0]

	split1 := strings.Split(line, ":")
	seedsStr := strings.Fields(split1[1])

	var seeds []int
	for _, seedStr := range seedsStr {
		num, err := strconv.Atoi(seedStr)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, num)
	}
	return seeds
}

func parseMappingTables(sections [][]string) []MappingTable {
	var maps []MappingTable
	for _, section := range sections {
		mapping := parseMappingTable(section)
		maps = append(maps, mapping)
	}
	return maps
}

// parseMappingTable converts the text input into a mappings. Each of those mappings is sort in ascending
// order to more easily find the ranges where there are or aren't intersections.
func parseMappingTable(section []string) MappingTable {
	mappingTable := MappingTable{}

	split1 := strings.Fields(section[0])
	mappingTable.Name = split1[0]

	for _, line := range section[1:] {
		split2 := strings.Fields(line)
		destinationStart, err := strconv.Atoi(split2[0])
		if err != nil {
			panic(err)
		}

		sourceStart, err := strconv.Atoi(split2[1])
		if err != nil {
			panic(err)
		}

		length, err := strconv.Atoi(split2[2])
		if err != nil {
			panic(err)
		}

		sourceRange := Range{
			Start: sourceStart,
			End:   sourceStart + length,
		}
		destinationRange := Range{
			Start: destinationStart,
			End:   sourceStart + length,
		}
		mappingTable.Mappings = append(mappingTable.Mappings, Mapping{
			SourceRange:      sourceRange,
			DestinationRange: destinationRange,
		})
	}

	// Sort the mappings in the table in ascending order by the start of the ranges in each mapping.
	sort.Slice(mappingTable.Mappings, func(i, j int) bool {
		return mappingTable.Mappings[i].SourceRange.Start < mappingTable.Mappings[j].SourceRange.Start
	})
	return mappingTable
}

func findLowestLocation(seeds []int, maps []MappingTable) int {
	lowest := math.MaxInt
	for _, seed := range seeds {
		location := findLocationForSeed(seed, maps)
		if location < lowest {
			lowest = location
		}
	}
	return lowest
}

// findLocationForSeed takes the initial seed, and translates through all the mappings
// to find the final location.
func findLocationForSeed(seed int, maps []MappingTable) int {
	source := seed
	for _, mapping := range maps {
		// The next source is the destination of the previous mapping.
		source = getDestination(source, mapping)
	}
	return source
}

func getDestination(source int, mappingTable MappingTable) int {
	for _, mapping := range mappingTable.Mappings {
		// If the source is within the range, return the destination.
		if source >= mapping.SourceRange.Start && source < mapping.SourceRange.End {
			// The destination is the same distance from the start of the range as the source is from the start of the range.
			return mapping.DestinationRange.Start + (source - mapping.SourceRange.Start)
		}
	}
	// If the source is not within any of the ranges, return the source.
	return source
}

/*
--- Part Two ---
Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds:
line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and
the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13
This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79
and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values:
55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84,
fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest
location number that corresponds to any of the initial seed numbers?
*/

// Brute force doesn't quite work for this. We instead need to find a smarter way to do this.
func Part2(lines []string) (int, error) {
	seeds, maps := parseInput(lines)
	seedRanges := convertSeedsToSeedRanges(seeds)
	return findLowestLocationForSeedRanges(seedRanges, maps), nil
}

func convertSeedsToSeedRanges(seeds []int) []Range {
	var seedRanges []Range
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		seedRanges = append(seedRanges, Range{
			Start: start,
			End:   start + length,
		})
	}
	return seedRanges
}

func findLowestLocationForSeedRanges(seedRanges []Range, maps []MappingTable) int {
	locationRanges := findLocationRangesForSeedRanges(seedRanges, maps)
	lowest := math.MaxInt
	for _, r := range locationRanges {
		if r.Start < lowest {
			lowest = r.Start
		}
	}
	return lowest
}

// findLocationRangesForSeedRanges takes the initial seed, and translates through all the mappings
// to find the final location.
func findLocationRangesForSeedRanges(seedRanges []Range, maps []MappingTable) []Range {
	sources := seedRanges
	for _, mapping := range maps {
		var nextSources []Range
		for _, s := range sources {
			mappedDestinations := getDestinationsForRange(s, mapping)
			nextSources = append(nextSources, mappedDestinations...)
		}
		sources = nextSources
	}
	// We want the lowest
	return sources
}

// getDestinationsForRange will go through the range and try to translate each number in its range based on how it
// appears in the mapping table. However, it is too slow to actually iterate through each number in this range. So
// instead we will check for overlaps of these ranges and translate each overlapping segment. If a section of the
// sourceRange doesn't overlap, then we will simply return the same numbers of that range that doesn't overlap,
// as per the rules.
func getDestinationsForRange(sourceRange Range, mappingTable MappingTable) []Range {
	var destinationRanges []Range

	r := sourceRange
	mappings := mappingTable.Mappings

	for r.Start != r.End && len(mappings) > 0 {
		m := mappings[0]

		// If these ranges don't overlap, then we can discard the mapping we're comparing to
		// and move on to the next one.
		if !doOverlap(r, m.SourceRange) {
			mappings = mappings[1:]
			continue
		}

		// If our start is before the mapping we're comparing with, then we directly use the source values
		// as the destination values. We don't want to actually convert to the destination values since this
		// range represents values that aren't mapped.
		if r.Start < m.SourceRange.Start {
			destinationRange := Range{
				Start: r.Start,
				End:   m.SourceRange.Start,
			}
			destinationRanges = append(destinationRanges, destinationRange)
			// Chop off that range of values so we don't try to look them up in our mapping again.
			r.Start = destinationRange.End
		} else {
			rangeToConvert := Range{
				Start: r.Start,
				End:   min(r.End, m.SourceRange.End),
			}
			destinationRange := convertToDestinationRange(rangeToConvert, m)
			destinationRanges = append(destinationRanges, destinationRange)
			// Chop off that range of values so we don't try to look them up in our mapping again.
			r.Start = rangeToConvert.End
		}
	}
	// At the end, if we still have a non-empty sourceRange, then add that to the destination.
	if r.Start != r.End {
		destinationRanges = append(destinationRanges, r)
	}
	return destinationRanges
}

func doOverlap(r, s Range) bool {
	return (r.Start >= s.Start && r.Start < s.End) ||
		(s.Start >= r.Start && s.Start < r.End)
}

func convertToDestinationRange(rangeToConvert Range, mapping Mapping) Range {
	startDist := rangeToConvert.Start - mapping.SourceRange.Start
	endDist := rangeToConvert.End - mapping.SourceRange.Start
	return Range{
		Start: mapping.DestinationRange.Start + startDist,
		End:   mapping.DestinationRange.Start + endDist,
	}
}
