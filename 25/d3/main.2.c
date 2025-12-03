#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "../logger/logger.h"

long long solve(char *line, ssize_t read) {
    char res[13] = {'0','0','0','0','0','0','0','0','0','0','0','0','\0'};
    int cursor = 0;
    for (int i=0; i < 12; i++){
        for (int j=cursor; j<read-1-(11-i); j++){
            if (line[j]=='9'){
                res[i]='9';
                cursor = j+1;
                break;
            }
            if (line[j] > res[i]) {
                res[i] = line[j];
                cursor = j+1;
            }
        }
    }

    // printf("s: %s :: %lld \n", res, strtoll(res, NULL, 10));
    return strtoll(res, NULL, 10);
}

void read_file(const char *filename) {
    FILE *file_ptr;
    char *line = NULL; // Buffer for the line
    size_t len = 0;   // Size of the allocated buffer
    ssize_t read;     // Number of characters read

    file_ptr = fopen(filename, "r");
    if (file_ptr == NULL) {
        perror("Error opening file");
        return;
    }

    // 2. Read lines one by one using getline
    printf("--- Reading file: %s ---\n", filename);

    long long res = 0;

    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getline(&line, &len,file_ptr)) != -1) {
        ldebug("Read: %zd, Ints: %s", read, line);
        res += solve(line, read);
    }

    // 3. Close the file and free the buffer
    printf("--- Finished ---\n");
    printf("RES: %lld\n", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }
    
    fclose(file_ptr);
}

int main(int argc, char *argv[]) {
    printf("Hello, d3!\n");
    if(strcmp(argv[1], "test")==0){
        logger(DEBUG);
        read_file("test");
    }else if (strcmp(argv[1], "input")==0) {
        logger(ERROR);
        read_file("input");
    }
    return 0;
}

