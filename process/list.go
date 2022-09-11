package process

var processes []*Process

// save *Process into memory
func SaveProcess(p *Process) {
	processes = append(processes, p)
}

func GetAll() []*Process {
	return processes
}
