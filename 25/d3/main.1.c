#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <sys/select.h>


int solve(char *line, ssize_t read) {

    char max_l = '0';
    int left_index = 0;
    char max_r = '0';
    for (int i = 0; i < read-2; i++){
        if (line[i] == '9') {
            max_l = line[i];
            left_index = i;
            break;
        }

        if (line[i] > max_l){
            max_l = line[i];
            left_index = i;
        }
    }

    for(int i=left_index+1; i < read-1; i++){
        if (line[i] == '9') {
            max_r = line[i];
            break;
        }

        if (line[i] > max_r){
            max_r = line[i];
        }
    }

    char res[3];
    res[0] = max_l;
    res[1] = max_r;
    res[2] = '\0';
    printf("s: %s l:%d\n", res, left_index);
    return atoi(res);
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

    int res = 0;

    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getline(&line, &len,file_ptr)) != -1) {
        printf("Read: %zd, Ints: %s", read, line);
        res += solve(line, read);
    }

    // 3. Close the file and free the buffer
    printf("--- Finished ---\n");
    printf("RES: %d\n", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }
    
    fclose(file_ptr);
}

int main(int argc, char *argv[]) {
    printf("Hello, d3!\n");
    read_file(argv[1]);
    return 0;
}
