package main

import (
    "encoding/json"
    "fmt"
    "os"
    "sort"
    "strings"
    "time"
)

type Stat struct {
    GeneratedAt int64             `json:"generated_at"`
    TotalFiles  int               `json:"total_files"`
    Protocols   map[string]int    `json:"protocols"`
    TopTalkers  map[string]int    `json:"top_talkers"`
    Notes       map[string]string `json:"notes"`
}

func fakeAnalyze(path string, st *Stat) {
    name := strings.ToLower(path)
    if strings.Contains(name, "http") {
        st.Protocols["http"] += 100
    }
    if strings.Contains(name, "dns") {
        st.Protocols["dns"] += 50
    }
    if strings.Contains(name, "tls") || strings.Contains(name, "ssl") {
        st.Protocols["tls"] += 80
    }
    st.TopTalkers["192.168.0.1"] += 10
    st.TopTalkers["10.0.0.5"] += 8
    st.Notes[path] = "t.me/Bengamin_Button t.me/XillenAdapter"
}

func main() {
    fmt.Println("t.me/Bengamin_Button t.me/XillenAdapter")
    if len(os.Args) < 2 {
        fmt.Println("usage: main <pcap1> <pcap2> ...")
        os.Exit(1)
    }
    st := &Stat{GeneratedAt: time.Now().Unix(), Protocols: map[string]int{}, TopTalkers: map[string]int{}, Notes: map[string]string{}}
    for _, p := range os.Args[1:] {
        st.TotalFiles++
        fakeAnalyze(p, st)
    }
    sorted := make([]struct{ k string; v int }, 0, len(st.Protocols))
    for k, v := range st.Protocols { sorted = append(sorted, struct{ k string; v int }{k, v}) }
    sort.Slice(sorted, func(i, j int) bool { return sorted[i].v > sorted[j].v })
    b, _ := json.MarshalIndent(st, "", "  ")
    _ = os.WriteFile("report.pcap.json", b, 0644)
    fmt.Println("report.pcap.json")
}

package main
import (
    "fmt"
)
func main(){
    fmt.Println("t.me/Bengamin_Button t.me/XillenAdapter")
}
