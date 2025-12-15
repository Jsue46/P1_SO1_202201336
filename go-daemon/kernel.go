package main

import (
    "encoding/json"
    "os"
)

type KernelProcess struct {
    PID        int    `json:"pid"`
    Nombre     string `json:"nombre"`
    UID        int    `json:"uid"`
    Cmdline    string `json:"cmdline"`
    VSZ        int    `json:"vsz_kb"`
    RSS        int    `json:"rss_kb"`
    CPU        int    `json:"cpu"`
    RAMPercent int    `json:"ram_percent"`
}

type KernelData struct {
    Procesos      []KernelProcess `json:"procesos"`
    TotalProcesos int             `json:"total_procesos"`
}

func ReadKernelProcesses() (KernelData, error) {
    const path = "/proc/continfo_so1_202201336"

    data, err := os.ReadFile(path)
    if err != nil {
        return KernelData{}, err
    }

    var kd KernelData
    err = json.Unmarshal(data, &kd)
    return kd, err
}
