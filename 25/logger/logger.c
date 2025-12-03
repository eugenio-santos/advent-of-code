#include "logger.h"
#include <stdio.h>
#include <stdarg.h>

LEVEL lvl = WARN;

void logger(LEVEL l) {
    lvl = l;
}

void ldebug(const char* format, ...) {
    if (lvl <= DEBUG){
        va_list args;
        va_start(args, format);
        vprintf(format, args);
        va_end(args);
    }
}