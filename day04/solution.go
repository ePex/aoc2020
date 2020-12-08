package main

import (
	"strconv"
	"strings"
)

type passport struct {
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

func ValidatePassportBatch(batch []string) (count int) {
	for _, record := range parseBatch(batch) {
		if passportValid(mapRecordToPassport(record)) {
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

func mapRecordToPassport(record string) passport {
	passport := passport{}

	entries := strings.Split(record, " ")
	m := make(map[string]string)
	for _, e := range entries {
		parts := strings.Split(e, ":")
		m[parts[0]] = parts[1]
	}

	passport.raw = record
	passport.byr, _ = strconv.Atoi(m["byr"])
	passport.iyr, _ = strconv.Atoi(m["iyr"])
	passport.eyr, _ = strconv.Atoi(m["eyr"])
	passport.hgt, _ = m["hgt"]
	passport.hcl, _ = m["hcl"]
	passport.ecl, _ = m["ecl"]
	passport.pid, _ = m["pid"]
	passport.cid, _ = strconv.Atoi(m["cid"])

	return passport
}

func passportValid(passport passport) bool {
	return passport.byr > 0 &&
		passport.iyr > 0 &&
		passport.eyr > 0 &&
		len(passport.hgt) > 0 &&
		len(passport.hcl) > 0 &&
		len(passport.ecl) > 0 &&
		len(passport.pid) > 0
}
