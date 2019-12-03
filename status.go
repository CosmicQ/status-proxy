package main

import (
  "fmt"
  "flag"
  "net/http"
  "os/exec"
)


func main() {
  http.HandleFunc("/", GetStatus)

  portPtr := flag.Int("port", 8181, "Port to run on, default 8181")

  listen := fmt.Sprintf("%s%d", ":", *portPtr)
  http.ListenAndServe(listen, nil)

}

func GetStatus(w http.ResponseWriter, r *http.Request) {

  check := "/usr/local/status/" + r.URL.Path[1:]
  cmd := exec.Command(check)
  out, err := cmd.Output()

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("FAIL... %s\n", out)))
  } else {
    fmt.Fprintf(w, "Pass")
  }

}