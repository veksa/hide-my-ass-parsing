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

type successCallback func(Proxy)
type errorCallback func(Proxy)

func TestProxies(proxies []Proxy, successFn successCallback, errorFn errorCallback) {
    for _, proxy := range proxies {
        check, _ := testPage("http://google.ru", proxy, "<title>Google</title>")
        if !check {
            successFn(proxy)
        } else {
            errorFn(proxy)
        }
    }
}
