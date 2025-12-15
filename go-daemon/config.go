package main

type Config struct {
    RamWarn  int
    RamStop  int
    RamKill  int

    CPUWarn  float64
    CPUStop  float64
    CPUKill  float64
}

var cfg = Config{
    RamWarn:  50,
    RamStop:  70,
    RamKill:  80,

    CPUWarn:  40.0,
    CPUStop:  60.0,
    CPUKill:  80.0,
}

// Archivo que expone el módulo del kernel
const ModulePath = "/proc/continfo_so1_202201336"
