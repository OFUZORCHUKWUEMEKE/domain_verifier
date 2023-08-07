package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error : could not read from input :%v\n", err)
	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}
	txt_records, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, record := range txt_records {
		if strings.HasPrefix(record, "v=spfi") {
			hasSPF = true
			sprRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	for _, records := range dmarcRecords {
		if strings.HasPrefix(records, "V=DMARC1") {
			hasDMARC = true
			dmarcRecord = records
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}
