package process

import (
	"database/sql"
	"regexp"

	"github.com/jasontconnell/sqlhelp"
)

const query string = `
	select
	lower(Replace(Replace(Replace(cast(TemplateID as char(36)), '{', ''), '}', ''), '-', '')) as Value
	from Items
	union
	select
	lower(Replace(Replace(Replace(cast(MasterID as char(36)), '{', ''), '}', ''), '-', '')) as Value
	from Items
	union
	select
	lower(Replace(Replace(Replace(Value, '{', ''), '}', ''), '-', '')) as Value
	from
	VersionedFields
	where ((Value like '[0-9A-F]%' or value like '{%') and len(Value)> 31) or lower(value) like '%<a %' or lower(value) like '%<img %' or lower(value) like '<image %'
	union
	select
	lower(Replace(Replace(Replace(Value, '{', ''), '}', ''), '-', '')) as Value
	from
	UnversionedFields
	where ((Value like '[0-9A-F]%' or value like '{%') and len(Value)> 31) or lower(value) like '%<a %' or lower(value) like '%<img %' or lower(value) like '<image %'
	union
	select
	lower(Replace(Replace(Replace(Value, '{', ''), '}', ''), '-', '')) as Value
	from
	SharedFields
	where ((Value like '[0-9A-F]%' or value like '{%') and len(Value)> 31) or lower(value) like '%<a %' or lower(value) like '%<img %' or lower(value) like '<image %'
`

func GetValues(connstr string) (map[string]bool, error) {
	conn, cerr := sql.Open("mssql", connstr)
	if cerr != nil {
		return nil, cerr
	}
	defer conn.Close()

	records, rerr := sqlhelp.GetResultSet(conn, query)
	if rerr != nil {
		return nil, rerr
	}

	m := make(map[string]bool)
	for _, row := range records {
		value := row["Value"].(string)
		ext := extract(value)
		for _, s := range ext {
			m[s] = true
		}
	}
	return m, nil
}

var reg *regexp.Regexp = regexp.MustCompile("[a-f0-9]{32}")

func extract(val string) []string {
	list := []string{}
	m := reg.FindAllStringSubmatch(val, -1)
	for _, s := range m {
		for _, ss := range s {
			list = append(list, ss)
		}
	}
	return list
}
