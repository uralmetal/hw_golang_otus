package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/valyala/fastjson" //nolint:depguard
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	statDomains := make(DomainStat, 300)
	var err error

	var p fastjson.Parser
	var v *fastjson.Value

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		v, err = p.Parse(line)
		if err != nil {
			continue
		}
		email := string(v.GetStringBytes("Email"))
		statDomains = countDomain(email, domain, statDomains)
	}
	err = scanner.Err()
	return statDomains, err
}

func countDomain(email, domain string, stat DomainStat) DomainStat {
	domainSuffix := "." + domain
	email = strings.ToLower(email)
	if strings.HasSuffix(email, domainSuffix) {
		domainEmail := strings.SplitN(email, "@", 2)[1]
		stat[domainEmail]++
	}
	return stat
}
