#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../logger/logger.h"

typedef struct Interval{
    unsigned long long left;
    unsigned long long right;
} Interval;

typedef struct Intervals {
    int size;
    Interval *arr;
} Intervals;

void printInt(Intervals *intervals){
    printf("Intervals Size: %d\n", intervals->size);
    printf("Elements:\n");

    for(int i=0; i <intervals->size; i++){
        printf("{l: %llu, r: %llu}\n", intervals->arr[i].left, intervals->arr[i].right);
    }
}

void push(Intervals *intervals, Interval i){
    int newSize = intervals->size + 1;
    Interval *temp = (Interval *)reallocarray(
        intervals->arr,
        newSize, 
        sizeof(Interval)
    );

    if (temp == NULL) {
        perror("Reallocation failed. Array size is unchanged.");
        exit(-1);
    } else {
        // Reallocation succeeded, update the original pointer
        intervals->arr = temp; 
        intervals->size = newSize;
        intervals->arr[newSize-1] = i;
        ldebug("New array size: %d\n", newSize);
    }
}


unsigned long long max(unsigned long long i1, unsigned long long i2){
    if (i1 > i2) return i1;
    return i2;
}

unsigned long long min(unsigned long long i1, unsigned long long i2){
    if (i1 < i2) return i1;
    return i2;
}

int intersects(Interval *i1, Interval i2){
    if ((i1->left > i2.left && i1->left < i2.right) || (i2.left > i1->left && i2.left < i1->right)){
        i1->left = min(i1->left, i2.left);
        i1->right = max(i1->right, i2.right);
        return 1;
    }
    return 0;
}

void buildInts(char *line, ssize_t read, Intervals *intervals) {
    char *dash_ptr = strchr(line, '-');
    int dashi = dash_ptr-line;
    ldebug("di %d\n", dashi);
    char ls[dashi+1];
    ls[dashi] = '\0';

    char rs[read-dashi];
    rs[read-dashi-1] = '\0';
    strncpy(ls, line,  dashi);
    strncpy(rs, dash_ptr+1,  read-dashi);

    printf("ls %s-%llu, rs %s", ls, strtoull(ls, NULL, 10), rs);
    Interval tmpint = {
        strtoull(ls, NULL, 10),
        strtoull(rs, NULL, 10)
    };

    int intsize = intervals->size;
    for(int i = 0; i < intsize; i++){
        if (intersects(&intervals->arr[i], tmpint) == 1 ){
            // printInt(intervals);
            return;
        } 
    }

    push(intervals, tmpint);
}

int solve(char *line, ssize_t read, Intervals *intervals) {
    unsigned long long val = strtoull(line, NULL, 10);
    for(int i=0; i < intervals->size; i++){
        if (val >= intervals->arr[i].left && val <= intervals->arr[i].right) {
            ldebug("valid: %llu\n", val);
            return 1;
        }
    }
    return 0;
}

// iterate each id and check if fresh

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

    int res = 0;

    struct Intervals *intervals = (Intervals *)malloc(sizeof(Intervals));
    intervals->size = 0;
    intervals->arr = NULL;

    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getline(&line, &len,file_ptr)) != -1) {
        if (line[0] == '\n') break;
        printf("Read: %zd, Ints: %s", read, line);
        buildInts(line, read, intervals);
    }

    printf("inbetw: %s", line);
    printInt(intervals);
    int i = 0;
    while ((read = getline(&line, &len,file_ptr)) != -1) {
        // printf("Read: %zd, Ints: %s", read, line);
        res += solve(line, read, intervals);
        i++;
    }

    // 3. Close the file and free the buffer
    printf("--- Finished ---\n");
    printf("RES: %d\n it: %d\n", res, i);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }

    free(intervals);
    
    fclose(file_ptr);
}

int main(int argc, char *argv[]) {
    printf("--- D5 ---\n");
    if(strcmp(argv[1], "test")==0){
        logger(DEBUG);
        read_file("test");
    }else if (strcmp(argv[1], "input")==0) {
        logger(ERROR);
        read_file("input");
    }
    return 0;
}
