#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../logger/logger.h"

typedef struct Node {
    int data;           // The value stored in the node
    struct Node* next;  // Pointer to the next node in the list
} Node;

typedef struct Stack {
    Node* top;          
} Stack;

Stack* createStack() {
    Stack* stack = (Stack*)malloc(sizeof(Stack));
    if (stack == NULL) {
        perror("Failed to allocate memory for Stack");
        exit(EXIT_FAILURE);
    }
    stack->top = NULL; 
    return stack;
}

int isEmpty(Stack* stack) {
    return stack->top == NULL;
}

void push(Stack* stack, int data) {
    // 1. Create a new node
    Node* newNode = (Node*)malloc(sizeof(Node));
    if (newNode == NULL) {
        perror("Failed to allocate memory for Node");
        return; // or exit(EXIT_FAILURE)
    }
    
    newNode->data = data;
    newNode->next = stack->top;
    stack->top = newNode;
}

int pop(Stack* stack) {
    if (isEmpty(stack)) {
        printf("Stack Underflow: Cannot pop from an empty stack.\n");
        return -1; 
    }
    
    Node* temp = stack->top;
    int poppedData = temp->data;
    stack->top = stack->top->next;
    free(temp);
    
    return poppedData;
}


void read_file(const char *filename) {
    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    char *buffer = NULL;
    long size = 0;

    fp = fopen(filename, "r");
    if (fp == NULL) {
        perror("Error opening file");
        return;
    }

    printf("--- Reading file: %s ---\n", filename);
    fseek(fp, 0, SEEK_END);
    size = ftell(fp);

    buffer = (char *)malloc(size + 1);
    rewind(fp);
    fread(buffer, 1, size, fp);
    buffer[size] = '\0';

    Stack *stack = createStack();

    char *new = strchr(buffer, '\n');
    int ll = new-buffer;
    printf("1st newline %d\n", ll);
    int res = 0;
    int bl = strlen(buffer);
    do {
        while (!isEmpty(stack)) {
            buffer[pop(stack)] = '.';
        }

        for (int i=0; i < bl; i++) {
            int sum = 0;
            if(buffer[i] == '.' || buffer[i] == '\n') continue;
            
            if(i-ll-2>=0 && buffer[i-ll-2] == '@') sum++;
            if(i-ll-1>=0 && buffer[i-ll-1] == '@') sum++;
            if(i-ll>=0 && buffer[i-ll] == '@') sum++;
            
            if(i-1>=0 && buffer[i-1] == '@') sum++;
            if(i+1>=0 && buffer[i+1] == '@') sum++;
            
            if(sum >= 4) continue;

            if(i+ll<bl && buffer[i+ll] == '@') sum++;
            if(i+ll+1<bl && buffer[i+ll+1] == '@') sum++;
            if(i+ll+2<bl && buffer[i+ll+2] == '@') sum++;

            if(sum < 4) {
                push(stack, i);
                res++;
                ldebug("it i %d res: %d\n", i, res);
            }
        }
    }while (!isEmpty(stack));

    printf("--- Finished ---\n");
    printf("RES: %d\n", res);
    
    // getline allocates memory for 'line', so it must be freed
    if (line) {
        free(line);
    }
    
    fclose(fp);
}

int main(int argc, char *argv[]) {
    printf("--- D4 ---\n");
    if(strcmp(argv[1], "test")==0){
        logger(DEBUG);
        read_file("test");
    }else if (strcmp(argv[1], "input")==0) {
        logger(ERROR);
        read_file("input");
    }
    return 0;
}
