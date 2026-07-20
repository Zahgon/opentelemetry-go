package embedded

type LoggerProvider interface{ loggerProvider() }

type Logger interface{ logger() }
