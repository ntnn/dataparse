# dataparse

Too often I have to work with these annyances:

1. CSV files that require to rewrite parsing for varying types from
   string
1. Excel files that report erroneous types
1. Unreliable APIs reporting e.g. integer as floats
1. Unreliable, (almost) undocumented APIs that return an object or
   a list depending on the number of results

etc.pp.

To solve these annoyances `dataparse` was born.

A onestop shop that makes it easy to retrieve information from varying
sources and handles the transformation between types.

## General use

### APIs

If an API does not offer an OpenAPI spec it is left to the consumer to
implement a client. Usually it is enough to have a look at the results
with `curl`, define structs with tags accordingly and then
`json.Unmarshal` into those.

Sometimes these APIs (especially SGML-to-JSON-wrapped and to a lesser
extend Java-backed APIs) report values in wild inconsistency, e.g.
reporting integers as floats or numbers as strings.

In those cases `dataparse.Map` can help:

```go
// Execute the request to the API
resp, err := http.Get("https://outdated-but-important.api/path/to/endpoint")
if err != nil {
    return err
}

// Read the returned JSON data into a dataparse.Map
m, err := dataparse.FromJsonSingle(resp.Body)
if err != nil {
    return err
}

i, err := m.Int("integer_value")
if err != nil {
    return err
}

log.Printf("integer value: %d")
```

In this case the API can return the integer as integer, string or float
and dataparse will transform it into the desired integer.

### Unmarshalling into structs

Another useful utility is unmarshalling data into structs, e.g. when
reading CSVs:

Assuming a CSV file with the headers `hostname,ip,logsize`:

```go
type myData struct {
    Hostname string  `dataparse:"hostname"`
    IPAddress net.IP `dataparse:"ip"`
    Logsize int      `dataparse:"logsize"`
}

// If the CSV file has no headers they can also be passed like this:
// dataparse.From("...", dataparse.WithHeaders("hostname", "ip", "logsize"))
mapCh, errCh, err := dataparse.From("/path/to/data.csv")
if err != nil {
    return err
}

for mapCh != nil || errCh != nil {
    select {
    case m, ok := <- mapCh:
        if !ok {
            mapCh = nil
            continue
        }
        // Read the CSV data into a struct to utilize the discrete types.
        d := myData{}
        if err := m.To(&d); err != nil {
            log.Errorf("error reading data: %v | %#v", err, m)
            continue
        }
        // handle d further
    case err, ok := <- errCh:
        if !ok {
            errCh = nil
            continue
        }
        log.Errorf("error from dataparse: %v", err)
    }
}
```
