package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	raw string
	byr int    // (Birth Year)
	iyr int    // (Issue Year)
	eyr int    // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid int    // (Country ID) - optional
}

func ValidatePassportBatchWithValidationRules(batch []string) (count int) {
	for _, record := range parseBatch(batch) {
		passport := Passport{}
		passport.fillWithData(record)
		if passport.mandatoryFieldsSet() && passport.valid() {
			count++
		}
	}

	return count
}

func ValidatePassportBatch(batch []string) (count int) {
	for _, record := range parseBatch(batch) {
		passport := Passport{}
		passport.fillWithData(record)
		if passport.mandatoryFieldsSet() {
			count++
		}
	}

	return count
}

func parseBatch(batch []string) (data []string) {
	record := ""
	for _, line := range batch {
		if len(line) == 0 {
			data = append(data, strings.TrimSpace(record))
			record = ""
		} else {
			record += " " + line
		}
	}
	data = append(data, strings.TrimSpace(record))

	return data
}

func (p *Passport) fillWithData(record string) {
	entries := strings.Split(record, " ")
	m := make(map[string]string)
	for _, e := range entries {
		parts := strings.Split(e, ":")
		m[parts[0]] = parts[1]
	}

	p.raw = record
	p.byr, _ = strconv.Atoi(m["byr"])
	p.iyr, _ = strconv.Atoi(m["iyr"])
	p.eyr, _ = strconv.Atoi(m["eyr"])
	p.hgt, _ = m["hgt"]
	p.hcl, _ = m["hcl"]
	p.ecl, _ = m["ecl"]
	p.pid, _ = m["pid"]
	p.cid, _ = strconv.Atoi(m["cid"])
}

func (p *Passport) mandatoryFieldsSet() bool {
	return p.byr > 0 &&
		p.iyr > 0 &&
		p.eyr > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0
}

func (p *Passport) valid() bool {
	return isValueInRange(p.byr, 1920, 2002) &&
		isValueInRange(p.iyr, 2010, 2020) &&
		isValueInRange(p.eyr, 2020, 2030) &&
		isHeightValid(p.hgt) &&
		hclCheck.MatchString(p.hcl) &&
		eclCheck.MatchString(p.ecl) &&
		pidCheck.MatchString(p.pid)
}

func isHeightValid(height string) bool {
	heightValue, _ := strconv.Atoi(height[:len(height)-2])

	switch height[len(height)-2:] {
	case "cm":
		if !isValueInRange(heightValue, 150, 193) {
			return false
		}
	case "in":
		if !isValueInRange(heightValue, 59, 76) {
			return false
		}
	default:
		return false
	}

	return true
}

func isValueInRange(value, min, max int) bool {
	return value >= min && value <= max
}

var hclCheck = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eclCheck = regexp.MustCompile(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`)
var pidCheck = regexp.MustCompile(`^\d{9}$`)
