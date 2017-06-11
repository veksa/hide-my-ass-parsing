package hideMyAssParsing

func GetProxies() []Proxy {
    var proxies []Proxy

    for _, proxyUrl := range proxyList {
        text, _ := getPage(proxyUrl)
        proxies = append(proxies, parseHideMyAssProxy(string(text))...)
    }

    var cleanedProxies []Proxy
    for _, value := range proxies {
        if !proxyInSlice(value, cleanedProxies) {
            cleanedProxies = append(cleanedProxies, value)
        }
    }

    return cleanedProxies
}

func TestProxies(proxies []Proxy) []Proxy {
    for i, proxy := range proxies {
        check, _ := testPage("http://google.ru", proxy, "<title>Google</title>")
        if !check {
            proxies = append(proxies[:i], proxies[i+1:]...)
        }
    }

    return proxies
}
