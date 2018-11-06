package model

// Config is used to pass data for the construction of a model instance
type Config struct {
	NZ        int32   // number of vertical levels
	NX        int32   // zonal grid resolution
	NY        int32   // meridional grid resolution
	L         float32 // domain length in meters
	W         float32 // domain width in meters
	DT        int64   // numerical timestep
	Twrite    int64   // interval for cfl and ke writeout (in timesteps)
	Tmax      int64   // total time of integration
	Tavestart int64   // start time for averaging
	Taveint   int64   // time interval used for summation in longterm average in seconds
	UseAB2    bool    // use second order Adams Bashforth timestepping instead of 3rd
	Rek       float64 // linear drag in lower layer
	Rbg       float64 // linear drag in each layer
	Filterfac float32 // the factor for use in the exponential filter
	F         float64 // coriolis parameter
	G         float32 // gravitational acceleration
	DiagList  string  // which diagnostics to output
	LogLevel  int32   // logger level: from 0 for quiet (no log) to 4 for verbose
	Logfile   string  // path to logfile, none prints to stdout
}
