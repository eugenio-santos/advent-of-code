#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../logger/logger.h"

typedef struct Beam{
    int row;
    int col;
    unsigned long long merged;
    struct Beam* next;
    struct Beam* prev;
} Beam;

typedef struct Stack {
    Beam* top;
    Beam* bot;
} Stack;

Stack* createStack() {
    Stack* stack = (Stack*)malloc(sizeof(Stack));
    if (stack == NULL) {
        perror("Failed to allocate memory for Stack");
        exit(EXIT_FAILURE);
    }
    stack->top = NULL; 
    stack->bot = NULL; 
    return stack;
}

int isEmpty(Stack* stack) {
    return stack->top == NULL;
}

void push(Stack* stack, int row, int col, unsigned long long merged) {
    // 1. Create a new Beam
    Beam* newBeam = (Beam*)malloc(sizeof(Beam));
    if (newBeam == NULL) {
        perror("Failed to allocate memory for Beam");
        return; // or exit(EXIT_FAILURE)
    }
    
    newBeam->row = row;
    newBeam->col = col;
    newBeam->merged = merged;
    newBeam->next = NULL;
    newBeam->prev = stack->top;
    
    if (stack->top != NULL) stack->top->next = newBeam;
    stack->top = newBeam;

    if(stack->bot == NULL) stack->bot = newBeam;
}

Beam* pop(Stack* stack) {
    if (isEmpty(stack)) {
        printf("Stack Underflow: Cannot pop from an empty stack.\n");
        return NULL; 
    }
    
    Beam* temp = stack->bot;
    if (stack->bot->next != NULL) {
        stack->bot->next->prev = NULL;
        stack->bot = stack->bot->next;
    }else{
        stack->bot = NULL;
        stack->top = NULL;
    }

    return temp;
}

void printm(char **m, int rown) {
    for(int i = 0; i < rown; i++){
        ldebug("%d %s",i, m[i]);
    }
    ldebug("m printed\n");
}

void find_and_merge(Stack* stack, int row, int col, unsigned long long merged, Beam *cursor) {
    if (cursor == NULL) cursor = stack->top;

    while(cursor->row !=row && cursor->col != col) {

        cursor = cursor->prev;
    }

    cursor->merged = cursor->merged + merged;
}

unsigned long long solve(char **m, int rown){
    unsigned long long res = 0;
    Stack *stack = createStack();
    char * S = strchr(m[0], 'S');
    push(stack, 0, (int)(S-m[0]), 1);
    ldebug("bug\n");
    Beam *b = pop(stack);

    for (int i = 0; i < rown-1; i++) {
        // ldebug("i: %d\n", i);
        while (b->row == i) {
            if (m[i+1][b->col] == '|') {
                find_and_merge(stack, i+1, b->col-1, b->merged, NULL);
            }

            if (m[i+1][b->col] == '^') {
                if (m[i+1][b->col-1] != '|') {
                    m[i+1][b->col-1] = '|';
                    push(stack, i+1, b->col-1, b->merged);
                } else { //merge
                    find_and_merge(stack, i+1, b->col-1, b->merged, NULL);
                }
                if (m[i+1][b->col+1] != '|') {
                    m[i+1][b->col+1] = '|';
                    push(stack, i+1, b->col+1, b->merged);
                } else { //merge
                    find_and_merge(stack, i+1, b->col-1, b->merged, NULL);
                }
                
            }
            
            if (m[i+1][b->col] == '.') {
                m[i+1][b->col] = '|';
                push(stack, i+1, b->col, b->merged);
            }
            
            free(b);
            b = pop(stack);
            if (b == NULL) {
                break;
            }
        }
        
    }
    
    printm(m, rown);
    while(b!=NULL){
        ldebug("m: %llu\n", b->merged);
        res += b->merged;
        free(b);
        b = pop(stack);
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
    int rown = 0;
    char **matrix = NULL;
    int i = 0;
    // getline returns -1 on failure to read a line (e.g., EOF)
    while ((read = getline(&line, &len,file_ptr)) != -1 ) {
        // ldebug("Read: %zd, Ints: %s", read, line);
        if (i % 2 == 0) {
            matrix = (char**)reallocarray(matrix, rown+1, sizeof(char*));
            matrix[rown] = (char *)malloc(read * sizeof(char));
            strcpy(matrix[rown], line);
            rown++;
        }
        i++;
    }

    // printm(matrix, rown);


    res = solve(matrix, rown);
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
    printf("--- D7 ---\n");

    if(strcmp(argv[1], "test")==0){
        logger(DEBUG);
        read_file("test");
    }else if (strcmp(argv[1], "input")==0) {
        logger(ERROR);
        read_file("input");
    }
    return 0;
}
