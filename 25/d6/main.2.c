#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../logger/logger.h"

int maxR; 

void printm(char **m) {
    for(int i = 0; i < maxR; i++){
        ldebug("%s", m[i]);
    }
}

int num_spaces(char *line, int c) {
    int r = 0;
    while(line[c+1] == ' ' || line[c+1] == '\n'){
        r++;
        c++;
    }

    return r;
}

int buildNum(char **m, int col, int d){
    char s[d+1];
    memset(s, '\0', sizeof(s));

    for(int i = 0; i < maxR; i++){
        char a = m[i][col+d];
        s[i] = a;
    }
    ldebug("b: %s, %d\n", s, atoi(s));
    return atoi(s);
}

unsigned long long solve(char **m, char *line){
    unsigned long long res = 0;
    int c = 0;
    while(line[c] != '\0'){
        unsigned long long sub = 0;
        if(line[c] == '*') sub=1;
        int digits = num_spaces(line, c);

        ldebug("d: %d, op: %c\n", digits, line[c]);
        for(int i = 0; i < digits; i++){
            if(line[c] == '*'){
                sub = sub * buildNum(m, c, i);
            }
            else {
                sub += buildNum(m, c, i);
            }
        }

        ldebug("sub: %d\n", sub);
        c+=digits+1;
        res += sub;
    }

    return res;
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

    unsigned long long res = 0;
    char **matrix = (char**)malloc(maxR * sizeof(char*));;
    int rown = 0;

    // getline returns -1 on failure to read a line (e.g., EOF)
    while (rown < maxR && (read = getline(&line, &len,file_ptr)) != -1 ) {
        ldebug("Read: %zd, Ints: %s", read, line);
        matrix[rown] = (char *)malloc(read * sizeof(char));
        // matrix[rown] = line;
        strcpy(matrix[rown], line);

        rown++;
    }

    printm(matrix);

    while ((read = getline(&line, &len,file_ptr)) != -1) {
        ldebug("Read: %zd, Ints: %s", read, line);
        res = solve(matrix, line);
    }

    // 3. Close the file and free the buffer
    printf("--- Finished ---\n");
    printf("RES: %llu\n ", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }

    fclose(file_ptr);
}

void up(int *i){
    *i = *i +1;
    (*i)++;
}

int main(int argc, char *argv[]) {
    printf("--- D6 ---\n");

    if(strcmp(argv[1], "test")==0){
        logger(DEBUG);
        maxR = 3;
        read_file("test");
    }else if (strcmp(argv[1], "input")==0) {
        logger(ERROR);
        maxR = 4;
        read_file("input");
    }
    return 0;
}
