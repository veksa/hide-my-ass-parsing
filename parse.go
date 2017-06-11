package hideMyAssParsing

import (
    "fmt"
    "regexp"
    "strings"
)

type Proxy struct {
    Ip      string
    Port    string
    Country string
    Type    string
    Anon    string
}

var proxyList = []string{
    "http://proxylist.hidemyass.com/search-1292985",
    "http://proxylist.hidemyass.com/search-1292985/2",
    "http://proxylist.hidemyass.com/search-1292985/3",
    "http://proxylist.hidemyass.com/search-1292985/4",
    "http://proxylist.hidemyass.com/search-1292985/5",
    "http://proxylist.hidemyass.com/search-1292985/6",
    "http://proxylist.hidemyass.com/search-1292985/7",
    "http://proxylist.hidemyass.com/search-1292985/8",
    "http://proxylist.hidemyass.com/search-1292985/9",
    "http://proxylist.hidemyass.com/search-1292985/10",
}

func parseHideMyAssProxy(text string) []Proxy {
    var proxies []Proxy

    text = strings.Replace(text, "\n", " ", -1)

    trRegex := regexp.MustCompile("<td>\\s*<span>\\s*<style>(.*?)</span>\\s*</td>\\s*<td>(.*?)</td>\\s*<td .*?>(.*?)</td>\\s*<td>(.*?)</td>\\s*<td>(.*?)</td>\\s*<td>(.*?)</td>\\s*<td.*?>(.*?)</td>")
    nodes := trRegex.FindAllStringSubmatch(text, -1)

    if nodes != nil {
        for _, node := range nodes {
            var str = node[1]

            displayNoneRegex := regexp.MustCompile(".(\\S*){display:none}")
            displayNoneStyle := displayNoneRegex.FindAllStringSubmatch(str, -1)

            for _, displayNone := range displayNoneStyle {
                var class = strings.Replace(displayNone[1], ".", "", -1)
                displayRegex := regexp.MustCompile(fmt.Sprintf("<[^>]*class=\"%v\">[^<]*<[^>]*>", class))
                str = displayRegex.ReplaceAllString(str, "")
            }

            displayNoneBlockRegex := regexp.MustCompile("<[^>]*style=\"display:none\">[^<]*<[^>]*>")
            str = displayNoneBlockRegex.ReplaceAllString(str, "")

            styleBlockRegex := regexp.MustCompile("(.*)</style>")
            str = styleBlockRegex.ReplaceAllString(str, "")

            tagsRegex := regexp.MustCompile("<[^>]*>")
            str = tagsRegex.ReplaceAllString(str, "")

            var ip = strings.Replace(str, " ", "", -1)
            var port = strings.Replace(node[2], " ", "", -1)
            var strType = strings.ToLower(strings.Replace(node[6], " ", "", -1))
            var anon = strings.ToLower(strings.Replace(node[7], " ", "", -1))

            var proxy Proxy
            proxy.Ip = ip
            proxy.Port = port
            proxy.Type = strType
            proxy.Anon = anon

            proxies = append(proxies, proxy)
        }
    }

    return proxies
}

func proxyInSlice(proxy Proxy, list []Proxy) bool {
    for _, p := range list {
        if p.Ip == proxy.Ip {
            return true
        }
    }
    return false
}
