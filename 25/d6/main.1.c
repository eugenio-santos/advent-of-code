#include "../logger/logger.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int maxR;

void printm(int **m) {
    for (int i = 0; i < maxR; i++) {
        for (int j = 0; j < 4; j++) {
            ldebug("%d ", m[i][j]);
        }
        ldebug("\n");
    }
}

void buildMatrix(int **m, char *line, int *rown) {
    char *token;
    int coln = 0;
    while ((token = strsep(&line, " "))) {
        if (strcmp(token, " ") != 0 && strcmp(token, "\n") != 0 &&
            strcmp(token, "") != 0) {
            ldebug("##%s##\n", token);
            coln++;
            m[*rown] = (int *)realloc(m[*rown], coln * sizeof(int));
            m[*rown][coln - 1] = atoi(token);
        }
    }

    (*rown)++;
}

unsigned long long solve(int **m, char *line) {
    unsigned long long res = 0;
    char *token;
    int coln = 0;
    while ((token = strsep(&line, " "))) {
        if (strcmp(token, " ") != 0 && strcmp(token, "\n") != 0 &&
            strcmp(token, "") != 0) {
            ldebug("##%s##\n", token);
            unsigned long long sub = 0;
            if (strcmp(token, "*") == 0) {
                sub = 1;
                for (int i = 0; i < maxR; i++) {
                    sub = sub * m[i][coln];
                }
            } else {
                for (int i = 0; i < maxR; i++) {
                    sub += m[i][coln];
                }
            }
            ldebug("%llu\n", sub);
            res += sub;
            coln++;
        }
    }

    return res;
}

void read_file(const char *filename) {
    FILE *file_ptr;
    char *line = NULL; // Buffer for the line
    size_t len = 0;    // Size of the allocated buffer
    ssize_t read;      // Number of characters read

    file_ptr = fopen(filename, "r");
    if (file_ptr == NULL) {
        perror("Error opening file");
        return;
    }

    // 2. Read lines one by one using getline
    printf("--- Reading file: %s ---\n", filename);

    unsigned long long res = 0;
    int **matrix = (int **)malloc(maxR * sizeof(int *));
    ;
    int rown = 0;

    // getline returns -1 on failure to read a line (e.g., EOF)
    while (rown < maxR && (read = getline(&line, &len, file_ptr)) != -1) {
        ldebug("Read: %zd, Ints: %s", read, line);
        buildMatrix(matrix, line, &rown);
    }

    printm(matrix);

    while ((read = getline(&line, &len, file_ptr)) != -1) {
        ldebug("Read: %zd, Ints: %s\n", read, line);
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

void up(int *i) {
    *i = *i + 1;
    (*i)++;
}

int main(int argc, char *argv[]) {
    printf("--- D5 ---\n");

    if (strcmp(argv[1], "test") == 0) {
        logger(DEBUG);
        maxR = 3;
        read_file("test");
    } else if (strcmp(argv[1], "input") == 0) {
        logger(ERROR);
        maxR = 4;
        read_file("input");
    }
    return 0;
}
