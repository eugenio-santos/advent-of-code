typedef enum{
    DEBUG,
    INFO,
    WARN,
    ERROR
} LEVEL;

void logger(LEVEL l);
void ldebug(const char* format, ...);