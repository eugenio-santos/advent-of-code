#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int mod(int a, int m) {
    int r = a % m;
    if (r < 0) {
        return m + r;
    }
    return r;
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
    int pos = 50;
    char val[4];
    val[3] = '\0';
    int res = 0;
    int newPos = 0;
    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getline(&line, &len, file_ptr)) != -1) {
        // 'line' now holds the read line, including the newline character if present
        // printf("Read %zd characters: %s", read, line);
        strncpy(val, line+1, read-1);
        int rotation = atoi(val);
        printf("dir: %c val: %d\n", line[0], atoi(val));
        
        // printf("rot %d\n", rotation/100);
        res += rotation/100;
        switch (line[0]) {
            case 'L': 
                newPos = mod(pos - rotation, 100);
                if (pos != 0 && (newPos == 0 || (newPos > pos && newPos < 100))) res++;
                break;
            case 'R':
                newPos = mod(pos + rotation, 100);
                if (pos != 0 && (newPos >= 0 && newPos <= pos)) res++;
                break;
        }

        // if (newPos == 0) {
            //     res++;
            // }
            pos = newPos;
            printf("pos: %d \n", pos);
    }

    // 3. Close the file and free the buffer
    printf("--- Finished %d ---\n", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }
    
    fclose(file_ptr);
}

int main() {
    printf("Hello, World!\n");
    read_file("test");
    read_file("input");
    printf("%d", mod(118, 95));
    return 0;
}
