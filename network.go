/*
 * Copyright (c) 2017 AlexRuzin (stan.ruzin@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package util

import (
    "net"
    "strconv"
    "net/http"
    "errors"
    "io/ioutil"
    "encoding/json"
)

func GetLocalIP() (localIP *string, err error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }

    /* FIXME -- make the report in JSON */

    var report = "LocalIP Information: "
    var interface_c = 0
    for _, i := range interfaces {
        addresses, err := i.Addrs()
        if err != nil {
            report += "<Failed to get interface #" + strconv.Itoa(interface_c) + ">, "
            continue
        }

        for _, address := range addresses {
            var ip net.IP
            switch v := address.(type) {
            case *net.IPNet:
                ip = v.IP
                report += "[Interface " + strconv.Itoa(interface_c) + "]: [IP]" + ip.String() + ", "
            case *net.IPAddr:
                ip = v.IP
                report += "[Interface " + strconv.Itoa(interface_c) + "]: [Ethernet]" + ip.String() + ", "
            }
        }

        interface_c += 1
    }

    return &report, nil
}

const GEOIP_URI = "https://ipinfo.io/json"
type GeoIP struct {
    IPString        string `json:"ip"`
    Hostname        string `json:"hostname"`
    City            string `json:"city"`
    Region          string `json:"region"`
    CountryCode     string `json:"country"`
    Coordinates     string `json:"loc"`
    ISP             string `json:"org"`
    Postal          string `json:"postal"`
}

func GetGeoIP() (geoip *GeoIP, err error) {
    var client http.Client
    response, err := client.Get(GEOIP_URI)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, errors.New("failed to obtain GEOIP JSON structure from: " + GEOIP_URI)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var parsed = &GeoIP{}
    if err := json.Unmarshal(body, parsed); err != nil {
        return nil, err
    }

    return parsed, nil
}

