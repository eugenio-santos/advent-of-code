#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

bool is_not_valid(long long num){
    char str[100];
    sprintf(str,"%lld",  num);

    int len = strlen(str);
    if (len % 2 != 0) {
        return false;   
    }
    char st[(len/2)+1];
    char nd[(len/2)+1];

    st[(len/2)] = '\0';
    nd[(len/2)] = '\0';

    strncpy(st, str, len/2);
    strncpy(nd, str+len/2, len/2);

    if (strcmp(st, nd) == 0) {
        return true;
    }
    return false;
}

unsigned long long solve(char *interval, ssize_t read) {
    unsigned long long sum = 0;
    int dash_index = 0;
    for (int i = 0; i < read; i++) {
        if (interval[i] == '-') {
            dash_index = i;
            break;
        }
    }

    char *start = calloc(read, sizeof(char));
    char *end = calloc(read, sizeof(char));

    strncpy(start, interval, dash_index);
    strncpy(end, interval+dash_index+1, read-dash_index);

    printf("%s-%s\n", start, end);

    unsigned long long llstrat = strtoull(start, NULL, 10); 
    unsigned long long llend = strtoull(end, NULL, 10);
    for (unsigned long long i = llstrat; i <= llend; i++){
        if(is_not_valid(i)){
            sum += i;
        }
    }

    free(start);
    free(end);
    printf("sum: %llu\n", sum);
    return sum;
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
    printf("res: %llu\n", res);

    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getdelim(&line, &len, ',',file_ptr)) != -1) {
        res +=  solve(line, read);
    }

    // 3. Close the file and free the buffer
    printf("--- Finished ---\n");
    printf("RES: %llu\n", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }
    
    fclose(file_ptr);
}

int main(int argc, char *argv[]) {
    printf("Hello, d2!\n");
    read_file(argv[1]);
    return 0;
}
