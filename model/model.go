package model

import "math"

// Model represents the information needed for a model run
type model struct {
	nz        int32   // number of vertical levels
	nx        int32   // zonal grid resolution
	ny        int32   // meridional grid resolution
	l         float32 // domain length in meters
	w         float32 // domain width in meters
	dt        int64   // numerical timestep
	twrite    int64   // interval for cfl and ke writeout (in timesteps)
	tmax      int64   // total time of integration
	tavestart int64   // start time for averaging
	taveint   int64   // time interval used for summation in longterm average in seconds
	useAB2    bool    // use second order Adams Bashforth timestepping instead of 3rd
	rek       float64 // linear drag in lower layer
	rbg       float64 // linear drag in each layer
	filterfac float32 // the factor for use in the exponential filter
	f         float64 // coriolis parameter
	f2        float64 // coriolis parameter squared
	g         float32 // gravitational acceleration
	diagList  string  // which diagnostics to output
	logLevel  int32   // logger level: from 0 for quiet (no log) to 4 for verbose
	logfile   string  // path to logfile, none prints to stdout
}

// NewModel constructs a new instance of the private model type
func NewModel(config Config) *model {
	if config.NZ == 0 {
		config.NZ = 1
	}
	if config.NX == 0 {
		config.NX = 64
	}
	if config.NY == 0 {
		config.NY = config.NX
	}
	if config.L == 0 {
		config.L = 1e6
	}
	if config.W == 0 {
		config.W = config.L
	}
	if config.DT == 0 {
		config.DT = 7200
	}
	if config.Twrite == 0 {
		config.Twrite = 1000
	}
	if config.Tavestart == 0 {
		config.Tavestart = 315360000
	}
	if config.Taveint == 0 {
		config.Taveint = 86400
	}
	if config.Rek == 0 {
		config.Rek = 5.787e-7
	}
	if config.Rbg == 0 {
		config.Rbg = config.Rek
	}
	if config.Filterfac == 0 {
		config.Filterfac = 23.6
	}
	if config.F == 0 {
		config.F = 1.028e-4
	}
	if config.G == 0 {
		config.G = 9.81
	}
	if config.DiagList == "" {
		config.DiagList = "all"
	}

	return &model{
		nz:        config.NZ,
		nx:        config.NX,
		ny:        config.NY,
		l:         config.L,
		w:         config.W,
		dt:        config.DT,
		twrite:    config.Twrite,
		tavestart: config.Tavestart,
		taveint:   config.Taveint,
		useAB2:    config.UseAB2,
		rek:       config.Rek,
		rbg:       config.Rbg,
		filterfac: config.Filterfac,
		f:         config.F,
		f2:        math.Pow(config.F, 2),
		g:         config.G,
		diagList:  config.DiagList,
		logLevel:  config.LogLevel,
		logfile:   config.Logfile,
	}
}
