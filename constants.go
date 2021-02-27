package golog

var (
	// DEBUG is lowest level in logging
	DEBUG Level = Level{
		Code:  uint8(10),
		Name:  "Debug",
		Short: "DBG",
	}

	// TIME is logging for timing measure only
	TIME Level = Level{
		Code:  uint8(9),
		Name:  "Time",
		Short: "TIM",
	}

	// INFO is common message
	INFO Level = Level{
		Code:  uint8(8),
		Name:  "Info",
		Short: "INF",
	}

	// WARN is a calm error message
	WARN Level = Level{
		Code:  uint8(7),
		Name:  "Warn",
		Short: "WRN",
	}

	// ERROR is a critical error message
	ERROR Level = Level{
		Code:  uint8(6),
		Name:  "Error",
		Short: "ERR",
	}

	// SILENT is a special level for mute all logging
	SILENT Level = Level{
		Code:  uint8(0),
		Name:  "Silent",
		Short: "",
	}
)

// Levels is a list of possible level
var Levels = []Level{
	DEBUG,
	TIME,
	INFO,
	WARN,
	ERROR,
	SILENT,
}
