package session

// The logger for the session.
import "github.com/ilyaqq1999/revel/logger"

var sessionLog logger.MultiLogger

func InitSession(coreLogger logger.MultiLogger) {
	sessionLog = coreLogger.New("section", "session")
}
