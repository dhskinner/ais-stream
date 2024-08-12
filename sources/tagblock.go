package sources

import (
	"fmt"
	"strconv"
	"strings"

	nmea "github.com/adrianmo/go-nmea"
)

// Sourced from https://github.com/adrianmo/go-nmea/blob/master/tagblock.go

// parseTagBlock adds support for tagblocks
// https://gpsd.gitlab.io/gpsd/AIVDM.html#_nmea_tag_blocks
func parseTagBlock(tags string) (nmea.TagBlock, error) {

	// check for zero length string
	if len(tags) == 0 {
		return nmea.TagBlock{}, nil
	}

	// strip any leading and trailing delimeters
	if tags[0] == '\\' {
		tags = tags[1:]
	}
	if tags[len(tags)-1] == '\\' {
		tags = tags[:len(tags)-1]
	}

	sumSepIndex := strings.Index(tags, nmea.ChecksumSep)
	if sumSepIndex == -1 {
		return nmea.TagBlock{}, fmt.Errorf("nmea: tagblock does not contain checksum separator")
	}

	var (
		fieldsRaw   = tags[0:sumSepIndex]
		checksumRaw = strings.ToUpper(tags[sumSepIndex+1:])
		checksum    = nmea.Checksum(fieldsRaw)
		tagBlock    nmea.TagBlock
		err         error
	)

	// validate the checksum
	if checksum != checksumRaw {
		return nmea.TagBlock{}, fmt.Errorf("nmea: tagblock checksum mismatch [%s != %s]", checksum, checksumRaw)
	}

	items := strings.Split(tags[:sumSepIndex], ",")
	for _, item := range items {
		parts := strings.SplitN(item, ":", 2)
		if len(parts) != 2 {
			return nmea.TagBlock{},
				fmt.Errorf("nmea: tagblock field is malformed (should be <key>:<value>) [%s]", item)
		}
		key, value := parts[0], parts[1]
		switch key {
		case "c": // UNIX timestamp
			tagBlock.Time, err = parseInt64(value)
			if err != nil {
				return nmea.TagBlock{}, err
			}
		case "d": // Destination ID
			tagBlock.Destination = value
		case "g": // Grouping
			tagBlock.Grouping = value
		case "n": // Line count
			tagBlock.LineCount, err = parseInt64(value)
			if err != nil {
				return nmea.TagBlock{}, err
			}
		case "r": // Relative time
			tagBlock.RelativeTime, err = parseInt64(value)
			if err != nil {
				return nmea.TagBlock{}, err
			}
		case "s": // Source ID
			tagBlock.Source = value
		case "t": // Text string
			tagBlock.Text = value
		}
	}
	return tagBlock, nil
}

func parseInt64(raw string) (int64, error) {
	i, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("nmea: tagblock unable to parse uint64 [%s]", raw)
	}
	return i, nil
}

func tagBlockAsString(tagBlock *nmea.TagBlock) string {

	// walnut \c:1713228743,s:rPiAIS001*11\!AIVDM,1,1,,B,17PCT2?P00:u5a3hDQ`ESgvd0L1C,0*61
	// \s:Cape Cleveland,c:1678711508*77\!AIVDM,1,1,2,B,B7P@UJ0002`33iM?ko<@`=T1VDb:,0*1E
	var result string = ""
	if tagBlock.Time != 0 {
		result += fmt.Sprintf("c:%d,", tagBlock.Time)
	}

	// source is limited to 15 characters
	if len(tagBlock.Source) > 15 {
		result += fmt.Sprintf("s:%s,", tagBlock.Source[:15])
	} else if len(tagBlock.Source) > 0 {
		result += fmt.Sprintf("s:%s,", tagBlock.Source)
	}

	// destination is limited to 15 characters
	if len(tagBlock.Destination) > 15 {
		result += fmt.Sprintf("d:%s,", tagBlock.Destination[:15])
	} else if len(tagBlock.Destination) > 0 {
		result += fmt.Sprintf("s:%s,", tagBlock.Destination)
	}

	if tagBlock.RelativeTime > 0 {
		result += fmt.Sprintf("r:%d,", tagBlock.RelativeTime)
	}
	if tagBlock.LineCount > 0 {
		result += fmt.Sprintf("n:%d,", tagBlock.LineCount)
	}
	if len(tagBlock.Grouping) > 0 {
		result += fmt.Sprintf("g:%s,", tagBlock.Grouping)
	}
	if len(tagBlock.Text) > 0 {
		result += fmt.Sprintf("t:%s,", tagBlock.Text)
	}

	// if the tagblock is empty, return ""
	if len(result) == 0 {
		return result
	}

	// remove the last comma
	result = result[:len(result)-1]

	// add a checksum and delimeters
	result = fmt.Sprintf("\\%s*%s\\", result, nmea.Checksum(result))

	return result
}
